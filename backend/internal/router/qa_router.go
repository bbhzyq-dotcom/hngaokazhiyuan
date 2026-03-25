package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/handler"
	"github.com/gaokao-advisor/backend/internal/middleware"
)

type QARouter struct {
	qaHandler *handler.QAHandler
}

func NewQARouter(qaHandler *handler.QAHandler) *QARouter {
	return &QARouter{qaHandler: qaHandler}
}

func (r *QARouter) RegisterRoutes(rg *gin.RouterGroup) {
	qa := rg.Group("/qa")
	qa.Use(middleware.JWTAuth())
	{
		qa.POST("/chat", r.qaHandler.Chat)
		qa.GET("/history", r.qaHandler.GetHistory)
		qa.POST("/recommend", r.qaHandler.Recommend)
		qa.GET("/colleges", r.qaHandler.GetRecommendColleges)
		qa.GET("/majors", r.qaHandler.GetRecommendMajors)
	}
}
