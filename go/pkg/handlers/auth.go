package handlers

import (
	"Web-chat/pkg/db/repositories/user"
	"Web-chat/pkg/requests"
	"Web-chat/pkg/services"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type AuthHandler struct {
	Hasher services.HasherI
	Jwt    services.JwtI
}

func NewAuthHandler(hasher services.HasherI, jwt services.JwtI) AuthHandler {
	return AuthHandler{Hasher: hasher, Jwt: jwt}
}

func (h AuthHandler) SignUp(c echo.Context) error {
	var user_ requests.SignUp
	err := c.Bind(&user_)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	password, err := h.Hasher.HashPassword(user_.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	success := user.Repository.CreateUser(user_.Email, password, user_.Name)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (h AuthHandler) SignIn(c echo.Context) error {
	var user_ requests.SignIn
	err := c.Bind(&user_)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	id, err := user.Repository.GetHashedPassword(user_.Email, user_.Password)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	t := h.Jwt.GenerateAccessToken(id)
	return c.JSON(http.StatusOK, map[string]string{
		"jwt": t,
	})
}
