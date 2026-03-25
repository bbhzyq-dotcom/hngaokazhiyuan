package model

import (
	"time"
)

type User struct {
	ID           uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string    `json:"username" gorm:"uniqueIndex;size:50;not null"`
	PasswordHash string    `json:"-" gorm:"size:255;not null"`
	Phone        string    `json:"phone" gorm:"uniqueIndex;size:20;not null"`
	UserType     string    `json:"user_type" gorm:"size:20;default:student"` // student/parent
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

type StudentProfile struct {
	UserID            uint64    `json:"user_id" gorm:"primaryKey"`
	Province          string    `json:"province" gorm:"size:50;default:河南省"`
	GaokaoYear        int       `json:"gaokao_year" gorm:"default:2025"`
	Score             int       `json:"score" gorm:"default:0"`
	Rank              int       `json:"rank" gorm:"default:0"`
	PreferredSubjects string    `json:"preferred_subjects" gorm:"type:text"` // JSON array
	PreferredRegions  string    `json:"preferred_regions" gorm:"type:text"`  // JSON array
	InterestTags      string    `json:"interest_tags" gorm:"type:text"`      // JSON array
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User              *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (StudentProfile) TableName() string {
	return "student_profiles"
}

type College struct {
	ID              uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string    `json:"name" gorm:"size:100;not null;index"`
	Province        string    `json:"province" gorm:"size:50;index"`
	City            string    `json:"city" gorm:"size:50"`
	District        string    `json:"district" gorm:"size:50"`
	Type            string    `json:"type" gorm:"size:50"`  // 综合/理工/师范...
	Level           string    `json:"level" gorm:"size:20"` // 本科/专科
	EstablishedYear int       `json:"established_year" gorm:"default:0"`
	Departments     int       `json:"departments" gorm:"default:0"`
	Faculties       int       `json:"faculties" gorm:"default:0"`
	Website         string    `json:"website" gorm:"size:255"`
	Description     string    `json:"description" gorm:"type:text"`
	Rankings        string    `json:"rankings" gorm:"type:json"`    // JSON object
	Disciplines     string    `json:"disciplines" gorm:"type:text"` // JSON array
	Statistics      string    `json:"statistics" gorm:"type:json"`  // JSON object
	LogoURL         string    `json:"logo_url" gorm:"size:255"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (College) TableName() string {
	return "colleges"
}

type Major struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:100;not null;index"`
	Category    string    `json:"category" gorm:"size:50;index"` // 学科门类
	Code        string    `json:"code" gorm:"size:20"`           // 专业代码
	Degree      string    `json:"degree" gorm:"size:50"`         // 学位类型
	Duration    int       `json:"duration" gorm:"default:4"`     // 学制年
	Description string    `json:"description" gorm:"type:text"`
	CoreCourses string    `json:"core_courses" gorm:"type:text"` // JSON array
	Employment  string    `json:"employment" gorm:"type:json"`   // JSON object
	CollegeIDs  string    `json:"college_ids" gorm:"type:text"`  // JSON array
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Major) TableName() string {
	return "majors"
}

type AdmissionScore struct {
	ID             uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	CollegeID      uint64    `json:"college_id" gorm:"index;not null"`
	MajorID        uint64    `json:"major_id" gorm:"index;not null"`
	Province       string    `json:"province" gorm:"size:50;default:河南省;index"`
	Year           int       `json:"year" gorm:"index;default:2024"`
	Batch          string    `json:"batch" gorm:"size:50;index"`    // 本科一批/本科二批...
	Category       string    `json:"category" gorm:"size:20;index"` // 理科/文科/综合
	ScienceScore   int       `json:"science_score" gorm:"default:0"`
	ArtsScore      int       `json:"arts_score" gorm:"default:0"`
	ScienceRankMin int       `json:"science_rank_min" gorm:"default:0"`
	ScienceRankMax int       `json:"science_rank_max" gorm:"default:0"`
	ArtsRankMin    int       `json:"arts_rank_min" gorm:"default:0"`
	ArtsRankMax    int       `json:"arts_rank_max" gorm:"default:0"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	College        *College  `json:"college,omitempty" gorm:"foreignKey:CollegeID"`
	Major          *Major    `json:"major,omitempty" gorm:"foreignKey:MajorID"`
}

func (AdmissionScore) TableName() string {
	return "admission_scores"
}

type Conversation struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `json:"user_id" gorm:"index;not null"`
	Title     string    `json:"title" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Messages  []Message `json:"messages,omitempty" gorm:"foreignKey:ConversationID"`
}

func (Conversation) TableName() string {
	return "conversations"
}

type Message struct {
	ID             uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ConversationID uint64    `json:"conversation_id" gorm:"index;not null"`
	Role           string    `json:"role" gorm:"size:20;not null"` // user/assistant/system
	Content        string    `json:"content" gorm:"type:text;not null"`
	UserContext    string    `json:"user_context" gorm:"type:json"` // JSON object
	Sources        string    `json:"sources" gorm:"type:text"`      // JSON array
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Message) TableName() string {
	return "messages"
}
