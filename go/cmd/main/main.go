package main

import (
	_ "Web-chat/config"
	"Web-chat/pkg/handler"
	"Web-chat/pkg/router"
	"Web-chat/pkg/service"
	"os"
)

func main() {
	authHandler := handler.NewAuthHandler(service.BcryptHasher{Cost: 5}, service.Jwt{})
	e := router.New(&authHandler)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
