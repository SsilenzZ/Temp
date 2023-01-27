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

var Clients = make(map[int]*ClientConn)

type ClientConn struct {
	Conn     *websocket.Conn
	Requests int
}

type ChatMessage struct {
	SenderID   int    `json:"senderID"`
	ReceiverID int    `json:"receiverID"`
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

func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func HandleMessages() {
	log.Print("Handler works")
	for {
		msg := <-MessageChannel
		blocked := social.SocialRepository.CheckBlocked(msg.ReceiverID, msg.SenderID)
		if blocked {
			Clients[msg.SenderID].Conn.WriteJSON(map[string]interface{}{"code": 0})
		} else {
			chatID := chat.ChatRepository.GetChatID(msg.ReceiverID, msg.SenderID)
			sentMessage := chat.MessageRepository.SendMessage(chatID, msg.SenderID, msg.Text)
			if Clients[msg.SenderID] != nil {
				sendMessageWithChatPreview(Clients[msg.SenderID].Conn, sentMessage.Text, chatID,
					sentMessage.SenderID, msg.ReceiverID, sentMessage.ID, sentMessage.CreatedAt)
			}
			if Clients[msg.ReceiverID] != nil {
				sendMessageWithChatPreview(Clients[msg.ReceiverID].Conn, sentMessage.Text, chatID,
					sentMessage.SenderID, sentMessage.SenderID, sentMessage.ID, sentMessage.CreatedAt)
			}
		}
	}
}

func sendMessageWithChatPreview(client *websocket.Conn, text string, chatID, senderID, receiverID,
	messageID int, time time.Time) {
	err := client.WriteJSON(map[string]interface{}{"chatID": chatID, "messageID": messageID, "userLogin": senderID,
		"userID": senderID, "text": text, "time": time, "senderID": receiverID})
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(Clients, senderID)
	}
}

func CloseConn(id int) {
	if Clients[id].Requests == 0 {
		Clients[id].Conn.Close()
		delete(Clients, id)
	} else {
		Clients[id].Requests--
	}
}
