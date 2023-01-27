package main

import (
	_ "Web-chat/config"
	"Web-chat/pkg/handlers"
	"Web-chat/pkg/router"
	"Web-chat/pkg/services"
	"os"
)

func main() {
	authHandler := handlers.NewAuthHandler(services.BcryptHasher{Cost: 5}, services.Jwt{})
	userHandler := handlers.UserHandler{}
	socialHandler := handlers.SocialHandler{}
	msgHandler := handlers.MessageHandler{}
	go services.HandleMessages()
	e := router.New(&authHandler, &userHandler, &socialHandler, &msgHandler)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
