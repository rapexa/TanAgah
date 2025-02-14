package repository

import (
	"TanAgah/internal/entity"
	"log"

	"gorm.io/gorm"
)

// JWTRepository interface for token methods
type JWTRepository interface {
	IsTokenValid(token string) bool
}

// JWTRepo implements JWTRepository
type JWTRepo struct {
	db *gorm.DB
}

// NewJWTRepo creates a new repository instance
func NewJWTRepo(db *gorm.DB) *JWTRepo {
	return &JWTRepo{db: db}
}

// IsTokenValid checks if the token is valid in the database
func (r *JWTRepo) IsTokenValid(token string) bool {
	var exists bool
	err := r.db.Model(&entity.User{}).Select("count(*) > 0").Where("jwt_token = ?", token).Scan(&exists).Error
	if err != nil {
		log.Println("Error checking token:", err)
		return false
	}
	return exists
}
