package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gaokao-advisor/backend/internal/config"
	"github.com/gaokao-advisor/backend/internal/handler"
	"github.com/gaokao-advisor/backend/internal/model"
	"github.com/gaokao-advisor/backend/internal/repository"
	"github.com/gaokao-advisor/backend/internal/router"
	"github.com/gaokao-advisor/backend/internal/service"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DB.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.StudentProfile{},
		&model.College{},
		&model.Major{},
		&model.AdmissionScore{},
		&model.Conversation{},
		&model.Message{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr(),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	userRepo := repository.NewUserRepository(db, rdb)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRouter := router.NewUserRouter(userHandler)

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userRouter.RegisterRoutes(v1)
	}

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
