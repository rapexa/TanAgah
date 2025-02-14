package main

import (
	"TanAgah/internal/config"
	"TanAgah/internal/controller"
	"TanAgah/internal/logger"
	"TanAgah/internal/repository"
	"TanAgah/internal/service"
	"TanAgah/pkg/middleware" // adjusted import path if needed
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	// Auth API group
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", userController.RegisterUser)
		authGroup.POST("/login", userController.LoginUser)
		authGroup.GET("/users/:id", userController.GetUser)
		authGroup.PUT("/users/:id", userController.UpdateUser)
		authGroup.DELETE("/users/:id", userController.DeleteUser)
		authGroup.POST("/users/:id/upload", fileController.HandleFileUpload)
	}

	// App API group
	appGroup := router.Group("/api/v1/app")
	appGroup.Use(middlewares.JWTMiddleware) // Using the JWT middleware
	{
		appGroup.GET("/data", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World")
		})
	}

	// Start server
	err := router.Run(cfg.AppPort)
	if err != nil {
		return
	}
}
