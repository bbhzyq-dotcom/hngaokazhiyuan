package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gaokao-advisor/backend/internal/config"
	"github.com/gaokao-advisor/backend/internal/model"
	"github.com/gaokao-advisor/backend/internal/repository"
	"github.com/gaokao-advisor/backend/pkg/utils"
)

type QAService struct {
	ragRepo  *repository.ESRepository
	userRepo *repository.UserRepository
	convRepo *repository.ConversationRepository
	aiConfig *config.AIConfig
}

func NewQAService(
	ragRepo *repository.ESRepository,
	userRepo *repository.UserRepository,
	convRepo *repository.ConversationRepository,
	aiConfig *config.AIConfig,
) *QAService {
	return &QAService{
		ragRepo:  ragRepo,
		userRepo: userRepo,
		convRepo: convRepo,
		aiConfig: aiConfig,
	}
}

type ChatRequest struct {
	Message        string          `json:"message"`
	ConversationID string          `json:"conversation_id"`
	UserContext    *UserContextDTO `json:"user_context"`
}

type UserContextDTO struct {
	Score             int      `json:"score"`
	Rank              int      `json:"rank"`
	PreferredSubjects []string `json:"preferred_subjects"`
	PreferredRegions  []string `json:"preferred_regions"`
	InterestTags      []string `json:"interest_tags"`
}

type ChatResponse struct {
	ConversationID  string              `json:"conversation_id"`
	MessageID       string              `json:"message_id"`
	Answer          string              `json:"answer"`
	Sources         []SourceDTO         `json:"sources"`
	Recommendations *RecommendationsDTO `json:"recommendations"`
}

type SourceDTO struct {
	Type      string  `json:"type"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Relevance float64 `json:"relevance"`
}

type RecommendationsDTO struct {
	Colleges []CollegeRecommendation `json:"colleges"`
	Majors   []MajorRecommendation   `json:"majors"`
}

type CollegeRecommendation struct {
	Name        string `json:"name"`
	Probability int    `json:"probability"`
	Reason      string `json:"reason"`
}

type MajorRecommendation struct {
	Name       string `json:"name"`
	MatchScore int    `json:"match_score"`
}

func (s *QAService) Chat(ctx context.Context, userID uint64, req *ChatRequest) (*ChatResponse, error) {
	var conv *model.Conversation
	var err error

	if req.ConversationID != "" {
		conv, err = s.convRepo.FindByID(parseUint64(req.ConversationID))
		if err != nil {
			conv = nil
		}
	}

	if conv == nil {
		conv = &model.Conversation{
			UserID: userID,
			Title:  truncateString(req.Message, 50),
		}
		s.convRepo.Create(conv)
	}

	userMsg := &model.Message{
		ConversationID: conv.ID,
		Role:           "user",
		Content:        req.Message,
	}
	if req.UserContext != nil {
		ctxJSON, _ := json.Marshal(req.UserContext)
		userMsg.UserContext = string(ctxJSON)
	}
	s.convRepo.CreateMessage(userMsg)

	prompt := s.buildPrompt(req.Message, req.UserContext)

	answer, err := s.callAI(ctx, prompt)
	if err != nil {
		answer = "抱歉，我现在无法回答您的问题，请稍后再试。"
	}

	assistantMsg := &model.Message{
		ConversationID: conv.ID,
		Role:           "assistant",
		Content:        answer,
	}
	s.convRepo.CreateMessage(assistantMsg)

	resp := &ChatResponse{
		ConversationID:  fmt.Sprintf("%d", conv.ID),
		MessageID:       fmt.Sprintf("%d", assistantMsg.ID),
		Answer:          answer,
		Recommendations: s.buildRecommendations(req.UserContext),
	}

	return resp, nil
}

func (s *QAService) buildPrompt(message string, ctx *UserContextDTO) string {
	var buf bytes.Buffer
	buf.WriteString("你是高考志愿填报指导助手，专门帮助河南高考考生提供报考建议。\n\n")

	if ctx != nil {
		buf.WriteString(fmt.Sprintf("考生信息：分数%d分，省内位次%d名。\n", ctx.Score, ctx.Rank))
		if len(ctx.PreferredSubjects) > 0 {
			buf.WriteString(fmt.Sprintf("偏好专业：%s\n", joinStrings(ctx.PreferredSubjects)))
		}
		if len(ctx.PreferredRegions) > 0 {
			buf.WriteString(fmt.Sprintf("偏好地区：%s\n", joinStrings(ctx.PreferredRegions)))
		}
		buf.WriteString("\n")
	}

	buf.WriteString(fmt.Sprintf("考生问题：%s\n\n", message))
	buf.WriteString("请根据考生信息和知识库中的数据，给出专业的建议。回答要具体、有针对性。")

	return buf.String()
}

func (s *QAService) callAI(ctx context.Context, prompt string) (string, error) {
	if s.aiConfig.Provider == "baidu" {
		return s.callBaiduAI(ctx, prompt)
	}
	return s.callMockAI(ctx, prompt)
}

type BaiduRequest struct {
	Messages []BaiduMessage `json:"messages"`
}

type BaiduMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type BaiduResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Result  string `json:"result"`
}

func (s *QAService) callBaiduAI(ctx context.Context, prompt string) (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions?access_token=%s", s.aiConfig.APIKey)

	reqBody := BaiduRequest{
		Messages: []BaiduMessage{
			{Role: "user", Content: prompt},
		},
	}

	data, _ := json.Marshal(reqBody)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result BaiduResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Result, nil
}

func (s *QAService) callMockAI(ctx context.Context, prompt string) (string, error) {
	time.Sleep(100 * time.Millisecond)

	var buf bytes.Buffer
	buf.WriteString("根据您的情况，我给出以下建议：\n\n")

	if s.aiConfig.APIKey == "" || s.aiConfig.APIKey == "your-api-key" {
		buf.WriteString("【说明】当前为演示模式，请配置AI模型API密钥以获得完整回答。\n\n")
	}

	buf.WriteString("1. 关于院校选择：建议您关注往年录取分数线与您分数相近的高校，可以考虑'冲稳保'策略。\n")
	buf.WriteString("2. 关于专业选择：结合您的兴趣和就业前景，建议选择计算机、电子信息等相关专业。\n")
	buf.WriteString("3. 关于志愿填报：建议将志愿拉开梯度，既有冲刺院校，也有保底院校。\n")

	return buf.String(), nil
}

func (s *QAService) buildRecommendations(ctx *UserContextDTO) *RecommendationsDTO {
	if ctx == nil {
		return nil
	}

	var recs RecommendationsDTO

	if ctx.Score > 600 {
		recs.Colleges = []CollegeRecommendation{
			{Name: "郑州大学", Probability: 90, Reason: "河南省内顶尖高校，往年录取分数线与您分数匹配"},
			{Name: "河南大学", Probability: 95, Reason: "河南省重点大学，录取概率较高"},
			{Name: "北京邮电大学", Probability: 65, Reason: "计算机专业强校，适合您冲刺"},
		}
	} else if ctx.Score > 500 {
		recs.Colleges = []CollegeRecommendation{
			{Name: "河南大学", Probability: 85, Reason: "综合实力强，往年录取分数线与您分数匹配"},
			{Name: "河南师范大学", Probability: 90, Reason: "师范类院校，录取概率较高"},
			{Name: "郑州轻工业大学", Probability: 92, Reason: "河南省重点高校，适合作为稳妥选择"},
		}
	} else {
		recs.Colleges = []CollegeRecommendation{
			{Name: "河南理工大学", Probability: 88, Reason: "河南省重点高校，录取概率较高"},
			{Name: "河南工业大学", Probability: 90, Reason: "综合实力不错，适合您的情况"},
			{Name: "河南农业大学", Probability: 95, Reason: "河南省属高校，录取概率高"},
		}
	}

	if len(ctx.PreferredSubjects) > 0 {
		for _, subj := range ctx.PreferredSubjects {
			recs.Majors = append(recs.Majors, MajorRecommendation{
				Name:       subj,
				MatchScore: 85 + (len(recs.Majors) % 10),
			})
		}
	} else {
		recs.Majors = []MajorRecommendation{
			{Name: "计算机科学与技术", MatchScore: 90},
			{Name: "电子信息工程", MatchScore: 85},
			{Name: "机械设计制造", MatchScore: 80},
		}
	}

	return &recs
}

type RecommendRequest struct {
	Score             int      `json:"score"`
	Rank              int      `json:"rank"`
	PreferredSubjects []string `json:"preferred_subjects"`
	PreferredRegions  []string `json:"preferred_regions"`
	Batch             string   `json:"batch"`
	PlanType          string   `json:"plan_type"`
}

type RecommendResponse struct {
	Strategy StrategyDTO `json:"strategy"`
	Analysis string      `json:"analysis"`
}

type StrategyDTO struct {
	Chong []CollegeMajorDTO `json:"chong"`
	Wen   []CollegeMajorDTO `json:"wen"`
	Bao   []CollegeMajorDTO `json:"bao"`
}

type CollegeMajorDTO struct {
	CollegeName string `json:"college_name"`
	MajorName   string `json:"major_name"`
	Probability int    `json:"probability"`
	ScoreDiff   int    `json:"score_diff"`
}

func (s *QAService) Recommend(ctx context.Context, req *RecommendRequest) (*RecommendResponse, error) {
	var strategy StrategyDTO

	strategy.Chong = []CollegeMajorDTO{
		{CollegeName: "北京航空航天大学", MajorName: "计算机科学与技术", Probability: 55, ScoreDiff: -15},
		{CollegeName: "西安电子科技大学", MajorName: "电子信息工程", Probability: 60, ScoreDiff: -10},
	}

	strategy.Wen = []CollegeMajorDTO{
		{CollegeName: "郑州大学", MajorName: "计算机科学与技术", Probability: 78, ScoreDiff: 5},
		{CollegeName: "北京邮电大学", MajorName: "软件工程", Probability: 75, ScoreDiff: 3},
	}

	strategy.Bao = []CollegeMajorDTO{
		{CollegeName: "河南大学", MajorName: "计算机科学与技术", Probability: 95, ScoreDiff: 30},
		{CollegeName: "河南理工大学", MajorName: "软件工程", Probability: 92, ScoreDiff: 25},
	}

	analysis := fmt.Sprintf("根据您的分数%d分和位次%d名分析：", req.Score, req.Rank)
	analysis += "冲刺志愿建议选择往年录取分数线略高于您分数的高校；"
	analysis += "稳妥志愿建议选择与您分数相当的高校；"
	analysis += "保底志愿建议选择往年录取分数线低于您分数20分以上的高校。"

	return &RecommendResponse{
		Strategy: strategy,
		Analysis: analysis,
	}, nil
}

func parseUint64(s string) uint64 {
	var v uint64
	fmt.Sscanf(s, "%d", &v)
	return v
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func joinStrings(arr []string) string {
	result := ""
	for i, s := range arr {
		if i > 0 {
			result += "、"
		}
		result += s
	}
	return result
}

func (s *QAService) GetHistory(userID uint64, page, pageSize int) ([]model.Conversation, int64, error) {
	return s.convRepo.FindByUserID(userID, page, pageSize)
}

func (s *QAService) GetRecommendColleges(userID uint64) ([]CollegeRecommendation, error) {
	profile, _ := s.userRepo.FindProfile(userID)

	recs := []CollegeRecommendation{
		{Name: "郑州大学", Probability: 90, Reason: "河南省内顶尖高校"},
		{Name: "河南大学", Probability: 95, Reason: "河南省重点大学"},
	}

	if profile != nil && profile.Score > 0 {
		if profile.Score > 600 {
			recs = append([]CollegeRecommendation{
				{Name: "北京大学", Probability: 30, Reason: "顶尖名校，冲刺目标"},
				{Name: "清华大学", Probability: 35, Reason: "顶尖名校，冲刺目标"},
			}, recs...)
		}
	}

	return recs, nil
}

func (s *QAService) GetRecommendMajors(userID uint64) ([]MajorRecommendation, error) {
	profile, _ := s.userRepo.FindProfile(userID)

	recs := []MajorRecommendation{
		{Name: "计算机科学与技术", MatchScore: 90},
		{Name: "电子信息工程", MatchScore: 85},
		{Name: "机械设计制造", MatchScore: 80},
	}

	if profile != nil && profile.PreferredSubjects != "" && profile.PreferredSubjects != "[]" {
		recs = []MajorRecommendation{}
		subjects := utils.ParseJSONArray(profile.PreferredSubjects)
		for _, subj := range subjects {
			recs = append(recs, MajorRecommendation{
				Name:       subj,
				MatchScore: 95,
			})
		}
	}

	return recs, nil
}
