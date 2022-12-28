package user

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/models/users"
)

var Repository users.UserRepository

func init() {
	Repository = users.ProvideUserRepository(connect.DB)
}
