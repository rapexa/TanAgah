package repository

import (
	"TanAgah/internal/model"
	"gorm.io/gorm"
)

type MessageRepository interface {
	GetMessages(senderID, receiverID uint) ([]model.Message, error)
	SaveMessage(message *model.Message) error
	DeleteMessage(messageID uint) error
	EditMessage(messageID uint, newContent string) error
	GetMessageByID(id uint) (*model.Message, error)
}

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) SaveMessage(message *model.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepo) GetMessages(senderID, receiverID uint) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", senderID, receiverID, receiverID, senderID).Find(&messages).Error
	return messages, err
}

func (r *MessageRepo) GetMessageByID(id uint) (*model.Message, error) {
	var Message model.Message
	err := r.db.First(&Message, id).Error
	return &Message, err
}

func (r *MessageRepo) DeleteMessage(messageID uint) error {
	return r.db.Delete(&model.Message{}, messageID).Error
}

func (r *MessageRepo) EditMessage(messageID uint, newContent string) error {
	return r.db.Model(&model.Message{}).Where("id = ?", messageID).Update("content", newContent).Error
}
