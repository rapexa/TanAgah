package main

import (
	"TanAgah/internal/config"
	"TanAgah/internal/controller"
	"TanAgah/internal/logger"
	"TanAgah/internal/repository"
	"TanAgah/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger()

	// Create uploads directory if not exists
	if err := os.MkdirAll(config.UPLOAD_DIR, 0755); err != nil {
		log.Fatal(err)
	}

	// Initialize database
	db := config.InitDB(cfg)

	// Initialize user dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Initialize file dependencies
	fileRepo := repository.NewFileRepository(cfg.UploadDir)
	fileService := service.NewFileService(fileRepo, cfg)
	fileController := controller.NewFileController(*fileService, userService)

	// Setup router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:  cfg.AllowedOrigins,
		AllowMethods:  []string{"POST", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	{
		api.POST("/users", userController.CreateUser)
		api.GET("/users/:id", userController.GetUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
		api.POST("/users/:id/upload", fileController.HandleFileUpload)
	}

	// Start server
	err := router.Run(cfg.AppPort)
	if err != nil {
		return
	}
}
