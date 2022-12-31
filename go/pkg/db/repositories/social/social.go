package social

import (
	"Web-chat/pkg/db/connect"
	"Web-chat/pkg/models/socials"
)

var Repository socials.SocialRepository

func init() {
	Repository = socials.ProvideSocialRepository(connect.DB)
}
