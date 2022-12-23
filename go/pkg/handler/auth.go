package handler

import (
	"Web-chat/pkg/db/repository/user"
	"Web-chat/pkg/requests"
	"Web-chat/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	Hasher service.HasherI
	Jwt    service.JwtI
}

func NewAuthHandler(hasher service.HasherI, jwt service.JwtI) AuthHandler {
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
