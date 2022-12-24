package router

import (
	"Web-chat/pkg/handler"
	"Web-chat/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(authHandler *handler.AuthHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/sign-up", authHandler.SignUp)
	e.POST("/sign-in", authHandler.SignIn)

	a := e.Group("/a")
	config := middleware.JWTConfig{
		Claims:     &service.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}

	a.Use(middleware.JWTWithConfig(config))
	a.POST("/hello", handler.Hello)
	return e
}
