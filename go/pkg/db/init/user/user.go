package user

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/db/repositories/user"
)

var UserRepository user.UserRepository

func init() {
	UserRepository = user.ProvideUserRepository(connect.DB)
}
