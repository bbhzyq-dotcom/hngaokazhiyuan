package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/middleware"
	"github.com/gaokao-advisor/backend/internal/service"
	"github.com/gaokao-advisor/backend/pkg/response"
)

type QAHandler struct {
	qaService *service.QAService
}

func NewQAHandler(qaService *service.QAService) *QAHandler {
	return &QAHandler{qaService: qaService}
}

func (h *QAHandler) Chat(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req service.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := h.qaService.Chat(c.Request.Context(), userID, &req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *QAHandler) GetHistory(c *gin.Context) {
	userID := middleware.GetUserID(c)
	page := 1
	pageSize := 20

	convs, total, err := h.qaService.GetHistory(userID, page, pageSize)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"total": total,
		"list":  convs,
	})
}

func (h *QAHandler) Recommend(c *gin.Context) {
	var req service.RecommendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := h.qaService.Recommend(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *QAHandler) GetRecommendColleges(c *gin.Context) {
	userID := middleware.GetUserID(c)

	colleges, err := h.qaService.GetRecommendColleges(userID)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"list": colleges})
}

func (h *QAHandler) GetRecommendMajors(c *gin.Context) {
	userID := middleware.GetUserID(c)

	majors, err := h.qaService.GetRecommendMajors(userID)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"list": majors})
}
