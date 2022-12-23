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
	var err error
	if os.Getenv("SERVER_PORT") == "" {
		err = os.Setenv("SERVER_PORT", ":3000")
		if err != nil {
			panic(err)
		}
	}
	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
