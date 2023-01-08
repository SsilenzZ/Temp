package chat

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/db/repositories/chat"
)

var ChatRepository chat.ChatRepository

func init() {
	ChatRepository = chat.ProvideChatRepository(connect.DB)
}
