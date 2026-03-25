package errors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrInvalidToken         = errors.New("invalid token")
	ErrTokenExpired         = errors.New("token expired")
	ErrInvalidParam         = errors.New("invalid parameter")
	ErrMissingParam         = errors.New("missing required parameter")
	ErrCollegeNotFound      = errors.New("college not found")
	ErrMajorNotFound        = errors.New("major not found")
	ErrScoreNotFound        = errors.New("score not found")
	ErrConversationNotFound = errors.New("conversation not found")
	ErrInternalServer       = errors.New("internal server error")
	ErrServiceUnavailable   = errors.New("service unavailable")
)
