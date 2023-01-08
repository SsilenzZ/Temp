package services

import (
	"Web-chat/pkg/db/init/chat"
	"Web-chat/pkg/db/init/social"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"time"
)

type WebSocketI interface {
	CloseConn(id int)
}

type WebSocket struct {
}

func (w WebSocket) init() {
	go w.handleMessages()
}

var Clients = make(map[int]*ClientConn)

type ClientConn struct {
	Conn     *websocket.Conn
	Requests int
}

type ChatMessage struct {
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
	Text       string `json:"text"`
}

var MessageChannel = make(chan ChatMessage, 1024)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (w WebSocket) unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func (w WebSocket) handleMessages() {
	for {
		msg := <-MessageChannel
		blocked := social.SocialRepository.CheckBlocked(msg.ReceiverId, msg.SenderId)
		if blocked {
			Clients[msg.SenderId].Conn.WriteJSON(map[string]interface{}{"code": 0})
		} else {
			chatId := chat.ChatRepository.GetChatID(msg.ReceiverId, msg.SenderId)
			sentMessage := chat.MessageRepository.SendMessage(chatId, msg.SenderId, msg.Text)
			if Clients[msg.SenderId] != nil {
				w.sendMessageWithChatPreview(Clients[msg.SenderId].Conn, sentMessage.Text, chatId,
					sentMessage.SenderID, msg.ReceiverId, sentMessage.ID, sentMessage.CreatedAt)
			}
			if Clients[msg.ReceiverId] != nil {
				w.sendMessageWithChatPreview(Clients[msg.ReceiverId].Conn, sentMessage.Text, chatId,
					sentMessage.SenderID, sentMessage.SenderID, sentMessage.ID, sentMessage.CreatedAt)
			}
		}
	}
}

func (w WebSocket) sendMessageWithChatPreview(client *websocket.Conn, text string, chatId, senderId, receiverId,
	messageId int, time time.Time) {
	err := client.WriteJSON(map[string]interface{}{"chatId": chatId, "messageId": messageId, "userLogin": senderId,
		"userId": senderId, "text": text, "time": time, "senderId": receiverId})
	if err != nil && w.unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(Clients, senderId)
	}
}

func (w WebSocket) CloseConn(id int) {
	if Clients[id].Requests == 0 {
		Clients[id].Conn.Close()
		delete(Clients, id)
	} else {
		Clients[id].Requests--
	}
}
