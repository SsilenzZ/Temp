package router

import (
	"Web-chat/pkg/handlers"
	"Web-chat/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, socialHandler *handlers.SocialHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/sign-up", authHandler.SignUp)
	e.POST("/sign-in", authHandler.SignIn)

	j := e.Group("/jwt")
	config := middleware.JWTConfig{
		Claims:     &services.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}

	j.Use(middleware.JWTWithConfig(config))
	j.POST("/find", userHandler.Find)

	j.POST("/friend", socialHandler.Friend)
	j.POST("/unfriend", socialHandler.Unfriend)
	j.GET("/friends", socialHandler.GetFriends)
	j.POST("/block", socialHandler.Block)
	j.POST("/unblock", socialHandler.Unblock)
	j.GET("/blacklist", socialHandler.GetBlacklist)
	return e
}
