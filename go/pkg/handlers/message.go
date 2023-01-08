package handlers

import (
	"Web-chat/pkg/db/init/chat"
	"Web-chat/pkg/db/init/user"
	"Web-chat/pkg/services"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type MessageHandler struct {
}

func (h MessageHandler) GetChat(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)
	id := claims.ID

	id2, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if id2 == id {
		return err
	}

	_, login, err := user.UserRepository.GetById(id2)
	if err != nil {
		return err
	}

	chatID := chat.ChatRepository.GetChatID(id, id2)

	messages := chat.MessageRepository.GetAllChatMessages(chatID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"user": map[string]interface{}{"id": id2, "login": login},
		"chatID": chatID, "chats": messages})
}

func (h MessageHandler) GetChatsPreview(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)
	id := claims.ID

	chats := chat.ChatRepository.GetChatsID(id)

	chatPreviews := make([]map[string]interface{}, len(chats))

	for i := 0; i < len(chats); i++ {
		msg, err := chat.MessageRepository.GetLastChatMessage(chats[i].ID)
		text := msg.Text
		time := map[int]interface{}{0: msg.CreatedAt, 1: ""}
		s := 0
		if err != nil {
			s = 1
		}
		messageId := msg.ID

		senderId := 0

		if chats[i].User_ID1 != id {
			senderId = chats[i].User_ID1
		} else {
			senderId = chats[i].User_ID2
		}

		senderId, senderLogin, err := user.UserRepository.GetById(senderId)

		if err != nil {
			return err
		}

		chatPreviews[i] = map[string]interface{}{"chatId": chats[i].ID, "messageId": messageId,
			"senderLogin": senderLogin, "senderId": senderId, "text": text, "time": time[s]}
	}

	return c.JSON(http.StatusOK, chatPreviews)
}
