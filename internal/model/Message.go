package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	SenderID   uint           `gorm:"not null" json:"sender_id"`
	ReceiverID uint           `gorm:"not null" json:"receiver_id"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	Edited     bool           `gorm:"default:false;not null" json:"edited"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
