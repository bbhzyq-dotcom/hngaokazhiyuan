package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/gaokao-advisor/backend/internal/model"
)

type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{db: db, redis: redis}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) CreateProfile(profile *model.StudentProfile) error {
	return r.db.Create(profile).Error
}

func (r *UserRepository) FindProfile(userID uint64) (*model.StudentProfile, error) {
	var profile model.StudentProfile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *UserRepository) UpdateProfile(profile *model.StudentProfile) error {
	return r.db.Save(profile).Error
}

func (r *UserRepository) SaveCode(ctx context.Context, phone, code string) error {
	key := fmt.Sprintf("sms:code:%s", phone)
	return r.redis.Set(ctx, key, code, 5*time.Minute).Err()
}

func (r *UserRepository) VerifyCode(ctx context.Context, phone, code string) (bool, error) {
	key := fmt.Sprintf("sms:code:%s", phone)
	storedCode, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return storedCode == code, nil
}

type CollegeRepository struct {
	db *gorm.DB
}

func NewCollegeRepository(db *gorm.DB) *CollegeRepository {
	return &CollegeRepository{db: db}
}

func (r *CollegeRepository) Create(college *model.College) error {
	return r.db.Create(college).Error
}

func (r *CollegeRepository) FindByID(id uint64) (*model.College, error) {
	var college model.College
	err := r.db.First(&college, id).Error
	if err != nil {
		return nil, err
	}
	return &college, nil
}

func (r *CollegeRepository) FindAll(page, pageSize int) ([]model.College, int64, error) {
	var colleges []model.College
	var total int64

	r.db.Model(&model.College{}).Count(&total)
	err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&colleges).Error
	return colleges, total, err
}

func (r *CollegeRepository) Search(keyword string, province, collegeType, level string, rankMin, rankMax, minScore int, page, pageSize int) ([]model.College, int64, error) {
	var colleges []model.College
	var total int64

	query := r.db.Model(&model.College{})

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if province != "" {
		query = query.Where("province = ?", province)
	}
	if collegeType != "" {
		query = query.Where("type = ?", collegeType)
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}

	query.Count(&total)
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&colleges).Error
	return colleges, total, err
}

type MajorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) *MajorRepository {
	return &MajorRepository{db: db}
}

func (r *MajorRepository) Create(major *model.Major) error {
	return r.db.Create(major).Error
}

func (r *MajorRepository) FindByID(id uint64) (*model.Major, error) {
	var major model.Major
	err := r.db.First(&major, id).Error
	if err != nil {
		return nil, err
	}
	return &major, nil
}

func (r *MajorRepository) FindAll(page, pageSize int) ([]model.Major, int64, error) {
	var majors []model.Major
	var total int64

	r.db.Model(&model.Major{}).Count(&total)
	err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&majors).Error
	return majors, total, err
}

func (r *MajorRepository) Search(keyword, category, degree string, page, pageSize int) ([]model.Major, int64, error) {
	var majors []model.Major
	var total int64

	query := r.db.Model(&model.Major{})

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if degree != "" {
		query = query.Where("degree = ?", degree)
	}

	query.Count(&total)
	err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&majors).Error
	return majors, total, err
}

type ScoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (r *ScoreRepository) Create(score *model.AdmissionScore) error {
	return r.db.Create(score).Error
}

func (r *ScoreRepository) FindByCollege(collegeID uint64, year int, province, category string) ([]model.AdmissionScore, error) {
	var scores []model.AdmissionScore
	query := r.db.Where("college_id = ?", collegeID)

	if year > 0 {
		query = query.Where("year = ?", year)
	}
	if province != "" {
		query = query.Where("province = ?", province)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Find(&scores).Error
	return scores, err
}

func (r *ScoreRepository) FindByMajor(majorID uint64, year int, province, category string) ([]model.AdmissionScore, error) {
	var scores []model.AdmissionScore
	query := r.db.Where("major_id = ?", majorID)

	if year > 0 {
		query = query.Where("year = ?", year)
	}
	if province != "" {
		query = query.Where("province = ?", province)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Find(&scores).Error
	return scores, err
}

func (r *ScoreRepository) FindProbability(collegeID uint64, majorID uint64, score, rank int, category string) (*model.AdmissionScore, error) {
	var s model.AdmissionScore
	query := r.db.Where("college_id = ?", collegeID)

	if majorID > 0 {
		query = query.Where("major_id = ?", majorID)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Order("year DESC").First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) Create(conv *model.Conversation) error {
	return r.db.Create(conv).Error
}

func (r *ConversationRepository) FindByID(id uint64) (*model.Conversation, error) {
	var conv model.Conversation
	err := r.db.Preload("Messages").First(&conv, id).Error
	if err != nil {
		return nil, err
	}
	return &conv, nil
}

func (r *ConversationRepository) FindByUserID(userID uint64, page, pageSize int) ([]model.Conversation, int64, error) {
	var convs []model.Conversation
	var total int64

	query := r.db.Where("user_id = ?", userID)
	query.Count(&total)
	err := query.Order("updated_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&convs).Error
	return convs, total, err
}

func (r *ConversationRepository) Update(conv *model.Conversation) error {
	return r.db.Save(conv).Error
}

func (r *ConversationRepository) CreateMessage(msg *model.Message) error {
	return r.db.Create(msg).Error
}

func (r *ConversationRepository) FindMessages(convID uint64) ([]model.Message, error) {
	var msgs []model.Message
	err := r.db.Where("conversation_id = ?", convID).Order("created_at ASC").Find(&msgs).Error
	return msgs, err
}

type ESRepository struct {
}

func NewESRepository() *ESRepository {
	return &ESRepository{}
}

type CollegeDoc struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Type        string    `json:"type"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	Majors      []string  `json:"majors"`
	Embedding   []float32 `json:"embedding,omitempty"`
}

func (r *ESRepository) BuildCollegeDoc(college *model.College) (*CollegeDoc, error) {
	var majors []string
	if college.Disciplines != "" {
		json.Unmarshal([]byte(college.Disciplines), &majors)
	}

	return &CollegeDoc{
		ID:          fmt.Sprintf("%d", college.ID),
		Name:        college.Name,
		Province:    college.Province,
		City:        college.City,
		Type:        college.Type,
		Level:       college.Level,
		Description: college.Description,
		Majors:      majors,
	}, nil
}

func (r *ESRepository) BuildMajorText(major *model.Major) string {
	var parts []string
	parts = append(parts, major.Name)
	parts = append(parts, major.Category)
	parts = append(parts, major.Description)

	if major.CoreCourses != "" {
		var courses []string
		json.Unmarshal([]byte(major.CoreCourses), &courses)
		parts = append(parts, strings.Join(courses, ","))
	}

	return strings.Join(parts, " ")
}
