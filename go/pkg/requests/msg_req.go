package requests

import "time"

type Message struct {
	ID          int       `json:"id"`
	ChatID      int       `json:"chatID"`
	SenderLogin string    `json:"senderLogin"`
	SenderID    int       `json:"senderID"`
	Text        string    `json:"text"`
	Time        time.Time `json:"time"`
}
