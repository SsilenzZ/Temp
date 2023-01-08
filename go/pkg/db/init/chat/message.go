package chat

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/db/repositories/chat"
)

var MessageRepository chat.MessageRepository

func init() {
	MessageRepository = chat.ProvideMessageRepository(connect.DB)
}
