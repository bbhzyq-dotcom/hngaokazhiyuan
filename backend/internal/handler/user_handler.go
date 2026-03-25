package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/middleware"
	"github.com/gaokao-advisor/backend/internal/service"
	"github.com/gaokao-advisor/backend/pkg/response"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := h.userService.Register(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, response.CodeParamError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := h.userService.Login(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, response.CodeParamError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	resp, err := h.userService.GetProfile(userID)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	var req service.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := h.userService.UpdateProfile(userID, &req); err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "profile updated"})
}

func (h *UserHandler) SendCode(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		response.MissingParam(c, "phone is required")
		return
	}

	if err := h.userService.SendCode(c.Request.Context(), phone); err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "code sent"})
}
