package service

import (
	"context"
	"errors"

	"github.com/gaokao-advisor/backend/internal/middleware"
	"github.com/gaokao-advisor/backend/internal/model"
	"github.com/gaokao-advisor/backend/internal/repository"
	apperrors "github.com/gaokao-advisor/backend/pkg/errors"
	"github.com/gaokao-advisor/backend/pkg/utils"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Code     string `json:"code" binding:"required,len=6"`
	UserType string `json:"user_type"`
}

type RegisterResponse struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

func (s *UserService) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	if req.UserType == "" {
		req.UserType = "student"
	}

	exists, _ := s.userRepo.FindByUsername(req.Username)
	if exists != nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	exists, _ = s.userRepo.FindByPhone(req.Phone)
	if exists != nil {
		return nil, apperrors.ErrUserAlreadyExists
	}

	valid, err := s.userRepo.VerifyCode(ctx, req.Phone, req.Code)
	if err != nil || !valid {
		return nil, apperrors.ErrInvalidParam
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: hashedPassword,
		Phone:        req.Phone,
		UserType:     req.UserType,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, apperrors.ErrInternalServer
	}

	profile := &model.StudentProfile{
		UserID:   user.ID,
		Province: "河南省",
	}
	s.userRepo.CreateProfile(profile)

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	return &RegisterResponse{
		UserID: user.ID,
		Token:  token,
	}, nil
}

type LoginRequest struct {
	LoginType string `json:"login_type" binding:"required"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Code      string `json:"code"`
}

type LoginResponse struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	UserType  string `json:"user_type"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func (s *UserService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var user *model.User
	var err error

	if req.LoginType == "password" {
		if req.Phone == "" || req.Password == "" {
			return nil, apperrors.ErrMissingParam
		}
		user, err = s.userRepo.FindByPhone(req.Phone)
		if err != nil {
			return nil, apperrors.ErrUserNotFound
		}
		if !utils.CheckPassword(req.Password, user.PasswordHash) {
			return nil, apperrors.ErrInvalidPassword
		}
	} else if req.LoginType == "phone" {
		if req.Phone == "" || req.Code == "" {
			return nil, apperrors.ErrMissingParam
		}
		valid, err := s.userRepo.VerifyCode(ctx, req.Phone, req.Code)
		if err != nil || !valid {
			return nil, apperrors.ErrInvalidParam
		}
		user, err = s.userRepo.FindByPhone(req.Phone)
		if err != nil {
			return nil, apperrors.ErrUserNotFound
		}
	} else {
		return nil, apperrors.ErrInvalidParam
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, apperrors.ErrInternalServer
	}

	return &LoginResponse{
		UserID:    user.ID,
		Username:  user.Username,
		Phone:     user.Phone,
		UserType:  user.UserType,
		Token:     token,
		ExpiresAt: "2025-06-26T12:00:00Z",
	}, nil
}

type ProfileResponse struct {
	UserID    uint64                `json:"user_id"`
	Username  string                `json:"username"`
	Phone     string                `json:"phone"`
	UserType  string                `json:"user_type"`
	CreatedAt string                `json:"created_at"`
	Profile   *model.StudentProfile `json:"profile,omitempty"`
}

func (s *UserService) GetProfile(userID uint64) (*ProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}

	profile, _ := s.userRepo.FindProfile(userID)

	return &ProfileResponse{
		UserID:    user.ID,
		Username:  user.Username,
		Phone:     user.Phone,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		Profile:   profile,
	}, nil
}

type UpdateProfileRequest struct {
	Province          string   `json:"province"`
	GaokaoYear        int      `json:"gaokao_year"`
	Score             int      `json:"score"`
	Rank              int      `json:"rank"`
	PreferredSubjects []string `json:"preferred_subjects"`
	PreferredRegions  []string `json:"preferred_regions"`
	InterestTags      []string `json:"interest_tags"`
}

func (s *UserService) UpdateProfile(userID uint64, req *UpdateProfileRequest) error {
	profile, err := s.userRepo.FindProfile(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			profile = &model.StudentProfile{UserID: userID}
		} else {
			return err
		}
	}

	if req.Province != "" {
		profile.Province = req.Province
	}
	if req.GaokaoYear > 0 {
		profile.GaokaoYear = req.GaokaoYear
	}
	if req.Score > 0 {
		profile.Score = req.Score
	}
	if req.Rank > 0 {
		profile.Rank = req.Rank
	}
	if len(req.PreferredSubjects) > 0 {
		profile.PreferredSubjects = utils.ToJSONArray(req.PreferredSubjects)
	}
	if len(req.PreferredRegions) > 0 {
		profile.PreferredRegions = utils.ToJSONArray(req.PreferredRegions)
	}
	if len(req.InterestTags) > 0 {
		profile.InterestTags = utils.ToJSONArray(req.InterestTags)
	}

	if profile.UserID == 0 {
		profile.UserID = userID
		return s.userRepo.CreateProfile(profile)
	}
	return s.userRepo.UpdateProfile(profile)
}

func (s *UserService) SendCode(ctx context.Context, phone string) error {
	code, err := utils.GenerateCode(6)
	if err != nil {
		return apperrors.ErrInternalServer
	}
	return s.userRepo.SaveCode(ctx, phone, code)
}
