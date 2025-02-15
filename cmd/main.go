package main

import (
	"TanAgah/internal/config"
	"TanAgah/internal/controller"
	"TanAgah/internal/logger"
	"TanAgah/internal/repository"
	"TanAgah/internal/service"
	"TanAgah/pkg/middleware"
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

	// Initialize jwt dependencies
	jwtRepo := repository.NewJWTRepo(db)

	// Initialize file dependencies
	messageRepo := repository.NewMessageRepo(db)
	messageService := service.NewMessageService(messageRepo)
	messageController := controller.NewMessageController(*messageService)

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

	// Auth API group
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", userController.RegisterUser)
		authGroup.POST("/login", userController.LoginUser)
		authGroup.GET("/ws/chat", messageController.ChatWebSocket)
		authGroup.GET("/chat/history/:sender_id/:receiver_id", messageController.GetChatHistory)
	}

	// App API group
	appGroup := router.Group("/api/v1/app")
	appGroup.Use(middleware.JWTMiddleware(jwtRepo)) // Using the JWT middleware
	{
		appGroup.POST("/users/:id/upload", fileController.HandleFileUpload)
		appGroup.POST("/users/:id/delete", userController.DeleteUser)
	}

	// Start server
	err := router.Run(cfg.AppPort)
	if err != nil {
		return
	}
}
