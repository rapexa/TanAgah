package service

import (
	"TanAgah/internal/model"
	"TanAgah/internal/repository"
)

type MessageService struct {
	Repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) *MessageService {
	return &MessageService{
		Repo: repo,
	}
}

func (s *MessageService) SaveMessageDb(message *model.Message) error {
	return s.Repo.SaveMessage(message)
}

func (s *MessageService) EditMessageDb(messageID uint, newContent string, senderID uint) error {

	message, err := s.Repo.GetMessageByID(messageID)

	if err != nil {
		return err
	}

	println(message.SenderID, senderID)

	if message.SenderID != senderID {
		return nil
	}

	return s.Repo.EditMessage(messageID, newContent)

}

func (s *MessageService) DeleteMessageDb(messageID uint, senderID uint) error {
	return s.Repo.DeleteMessage(messageID)
}

func (s *MessageService) GetChatHistory(senderID, receiverID uint) ([]model.Message, error) {
	return s.Repo.GetMessages(senderID, receiverID)
}
