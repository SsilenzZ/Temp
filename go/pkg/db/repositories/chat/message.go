package chat

import (
	"Web-chat/pkg/models"
	"gorm.io/gorm"
)

type MessageRepositoryI interface{}

type MessageRepository struct {
	DB *gorm.DB
}

func ProvideMessageRepository(DB *gorm.DB) MessageRepository {
	return MessageRepository{DB: DB}
}

func (m MessageRepository) GetAllChatMessages(id int) []models.Message {
	var msgs []models.Message
	m.DB.Where("chat_id = ?", id).Find(&msgs)
	return msgs
}

func (m MessageRepository) GetLastChatMessage(id int) (models.Message, error) {
	var msg models.Message
	err := m.DB.Where("chat_id = ?", id).Last(&msg).Error
	return msg, err
}

func (m *MessageRepository) SendMessage(chatId, userId int, text string) models.Message {
	msg := models.Message{ChatID: chatId, SenderID: userId, Text: text}
	m.DB.Create(&msg)
	return msg
}
