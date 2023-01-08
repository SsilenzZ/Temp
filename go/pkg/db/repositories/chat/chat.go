package chat

import (
	"Web-chat/pkg/models"
	"gorm.io/gorm"
)

type ChatRepositoryI interface {
}

type ChatRepository struct {
	DB *gorm.DB
}

func ProvideChatRepository(DB *gorm.DB) ChatRepository {
	return ChatRepository{DB: DB}
}

func (c ChatRepository) GetChatID(id1, id2 int) int {
	var chat models.Chat
	if c.DB.Where("user_id1 = ? AND user_id2 = ? OR user_id1 = ? AND user_id2 = ?", id1, id2, id2, id1).
		Take(&chat).Error != nil {
		chat.User_ID1 = id1
		chat.User_ID2 = id2
		c.DB.Create(&chat)
		return chat.ID
	}
	return chat.ID
}

func (c ChatRepository) GetChatsID(id int) []models.Chat {
	var chats []models.Chat
	c.DB.Where("user_id1 = ? OR user_id2 = ?", id, id).Find(&chats)
	return chats
}
