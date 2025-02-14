package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
	RoleUser   = "user"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	AppPort        string
	UploadDir      string
	MaxUploadSize  int64
	AllowedOrigins []string
	AllowedTypes   []string
	JwtSecret      string
}

func LoadConfig() *Config {
	godotenv.Load()
	return &Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		AppPort:        os.Getenv("APP_PORT"),
		UploadDir:      os.Getenv("UPLOAD_DIR"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		MaxUploadSize:  10 << 20, // 10MB default
		AllowedOrigins: []string{os.Getenv("ALLOWED_ORIGINS")},
		AllowedTypes:   []string{"image/jpeg", "image/png", "application/pdf"},
	}
}

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
