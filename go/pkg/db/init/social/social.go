package social

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/db/repositories/social"
)

var SocialRepository social.SocialRepository

func init() {
	SocialRepository = social.ProvideSocialRepository(connect.DB)
}
