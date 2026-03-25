package service

import (
	"github.com/gaokao-advisor/backend/internal/model"
	"github.com/gaokao-advisor/backend/internal/repository"
	apperrors "github.com/gaokao-advisor/backend/pkg/errors"
)

type DataService struct {
	collegeRepo *repository.CollegeRepository
	majorRepo   *repository.MajorRepository
	scoreRepo   *repository.ScoreRepository
}

func NewDataService(
	collegeRepo *repository.CollegeRepository,
	majorRepo *repository.MajorRepository,
	scoreRepo *repository.ScoreRepository,
) *DataService {
	return &DataService{
		collegeRepo: collegeRepo,
		majorRepo:   majorRepo,
		scoreRepo:   scoreRepo,
	}
}

type CollegeListRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
	Province string `json:"province"`
	Type     string `json:"type"`
	Level    string `json:"level"`
	RankMin  int    `json:"rank_min"`
	RankMax  int    `json:"rank_max"`
	MinScore int    `json:"min_score"`
}

type CollegeListResponse struct {
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     []CollegeItem `json:"list"`
}

type CollegeItem struct {
	ID             uint64       `json:"id"`
	Name           string       `json:"name"`
	Province       string       `json:"province"`
	City           string       `json:"city"`
	Type           string       `json:"type"`
	Level          string       `json:"level"`
	Ranking        *RankingInfo `json:"ranking,omitempty"`
	LogoURL        string       `json:"logo_url,omitempty"`
	AdmissionScore *ScoreInfo   `json:"admission_score,omitempty"`
}

type RankingInfo struct {
	Comprehensive int `json:"comprehensive"`
	Province      int `json:"province"`
}

type ScoreInfo struct {
	Science int `json:"science"`
	Arts    int `json:"arts"`
}

func (s *DataService) GetColleges(req *CollegeListRequest) (*CollegeListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	colleges, total, err := s.collegeRepo.Search(
		req.Keyword, req.Province, req.Type, req.Level,
		req.RankMin, req.RankMax, req.MinScore,
		req.Page, req.PageSize,
	)
	if err != nil {
		return nil, err
	}

	list := make([]CollegeItem, 0, len(colleges))
	for _, c := range colleges {
		item := CollegeItem{
			ID:       c.ID,
			Name:     c.Name,
			Province: c.Province,
			City:     c.City,
			Type:     c.Type,
			Level:    c.Level,
			LogoURL:  c.LogoURL,
		}
		list = append(list, item)
	}

	return &CollegeListResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}

func (s *DataService) GetCollegeByID(id uint64) (*model.College, error) {
	return s.collegeRepo.FindByID(id)
}

func (s *DataService) GetCollegeMajors(collegeID uint64, page, pageSize int, category string) ([]model.Major, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	var majors []model.Major
	var total int64
	var err error

	majors, total, err = s.majorRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return majors, total, nil
}

type MajorListRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
	Category string `json:"category"`
	Degree   string `json:"degree"`
}

type MajorListResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	List     []MajorItem `json:"list"`
}

type MajorItem struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Code     string `json:"code"`
	Degree   string `json:"degree"`
	Duration int    `json:"duration"`
}

func (s *DataService) GetMajors(req *MajorListRequest) (*MajorListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	majors, total, err := s.majorRepo.Search(req.Keyword, req.Category, req.Degree, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]MajorItem, 0, len(majors))
	for _, m := range majors {
		item := MajorItem{
			ID:       m.ID,
			Name:     m.Name,
			Category: m.Category,
			Code:     m.Code,
			Degree:   m.Degree,
			Duration: m.Duration,
		}
		list = append(list, item)
	}

	return &MajorListResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}

func (s *DataService) GetMajorByID(id uint64) (*model.Major, error) {
	return s.majorRepo.FindByID(id)
}

type ScoreQueryRequest struct {
	CollegeID uint64 `json:"college_id"`
	MajorID   uint64 `json:"major_id"`
	Province  string `json:"province"`
	Year      int    `json:"year"`
	Batch     string `json:"batch"`
	Category  string `json:"category"`
}

type ScoreResponse struct {
	CollegeID   uint64      `json:"college_id"`
	CollegeName string      `json:"college_name"`
	MajorID     uint64      `json:"major_id"`
	MajorName   string      `json:"major_name"`
	Year        int         `json:"year"`
	Batch       string      `json:"batch"`
	Category    string      `json:"category"`
	Score       *ScoreRange `json:"score"`
	Rank        *RankRange  `json:"rank"`
}

type ScoreRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
	Avg int `json:"avg"`
}

type RankRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (s *DataService) GetScores(req *ScoreQueryRequest) ([]ScoreResponse, error) {
	if req.Province == "" {
		req.Province = "河南省"
	}
	if req.Year == 0 {
		req.Year = 2024
	}

	var scores []model.AdmissionScore
	var err error

	if req.CollegeID > 0 {
		scores, err = s.scoreRepo.FindByCollege(req.CollegeID, req.Year, req.Province, req.Category)
	} else if req.MajorID > 0 {
		scores, err = s.scoreRepo.FindByMajor(req.MajorID, req.Year, req.Province, req.Category)
	} else {
		return nil, apperrors.ErrMissingParam
	}

	if err != nil {
		return nil, err
	}

	result := make([]ScoreResponse, 0, len(scores))
	for _, sc := range scores {
		item := ScoreResponse{
			CollegeID: sc.CollegeID,
			MajorID:   sc.MajorID,
			Year:      sc.Year,
			Batch:     sc.Batch,
			Category:  sc.Category,
			Score: &ScoreRange{
				Min: sc.ScienceScore,
				Max: sc.ArtsScore,
			},
			Rank: &RankRange{
				Min: sc.ScienceRankMin,
				Max: sc.ScienceRankMax,
			},
		}
		result = append(result, item)
	}

	return result, nil
}

type ProbabilityRequest struct {
	CollegeID uint64 `json:"college_id"`
	MajorID   uint64 `json:"major_id"`
	Score     int    `json:"score"`
	Rank      int    `json:"rank"`
	Category  string `json:"category"`
}

type ProbabilityResponse struct {
	CollegeName string  `json:"college_name"`
	MajorName   string  `json:"major_name"`
	Probability float64 `json:"probability"`
	Level       string  `json:"level"`
	Analysis    string  `json:"analysis"`
}

func (s *DataService) GetProbability(req *ProbabilityRequest) (*ProbabilityResponse, error) {
	if req.Category == "" {
		req.Category = "理科"
	}

	score, err := s.scoreRepo.FindProbability(req.CollegeID, req.MajorID, req.Score, req.Rank, req.Category)
	if err != nil {
		return nil, err
	}

	college, _ := s.collegeRepo.FindByID(req.CollegeID)

	var probability float64 = 50
	var level string

	if score.ScienceScore > 0 {
		diff := req.Score - score.ScienceScore
		if diff > 10 {
			probability = 95
			level = "稳"
		} else if diff > 0 {
			probability = 80
			level = "稳"
		} else if diff > -10 {
			probability = 60
			level = "冲"
		} else {
			probability = 40
			level = "冲"
		}
	}

	resp := &ProbabilityResponse{
		Probability: probability,
		Level:       level,
		Analysis:    "根据您的高考分数和位次分析",
	}

	if college != nil {
		resp.CollegeName = college.Name
	}

	return resp, nil
}
