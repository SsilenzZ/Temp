package requests

type Message struct {
	ChatID      int                 `json:"chatId"`
	MessageID   int                 `json:"messageId"`
	SenderLogin string              `json:"senderLogin"`
	SenderID    int                 `json:"senderId"`
	Text        string              `json:"text"`
	Time        map[int]interface{} `json:"time"`
}
