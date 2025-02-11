package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;uniqueIndex;not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
