package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chatID"`
	SenderID  int       `json:"senderID"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
