package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/handler"
	"github.com/gaokao-advisor/backend/internal/middleware"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{userHandler: userHandler}
}

func (r *UserRouter) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")
	{
		users.POST("/register", r.userHandler.Register)
		users.POST("/login", r.userHandler.Login)
		users.GET("/profile", middleware.JWTAuth(), r.userHandler.GetProfile)
		users.PUT("/profile", middleware.JWTAuth(), r.userHandler.UpdateProfile)
		users.POST("/send-code", r.userHandler.SendCode)
	}
}
