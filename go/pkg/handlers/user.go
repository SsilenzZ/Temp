package handlers

import (
	"Web-chat/pkg/db/repositories/user"
	"Web-chat/pkg/requests"
	"Web-chat/pkg/services"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
}

func (h UserHandler) Find(c echo.Context) error {
	var user_ requests.Find
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)
	err := c.Bind(&user_)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	foundUsers, err := user.Repository.Find(user_.Name, claims.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	users := make([]map[string]interface{}, len(foundUsers))

	for i := 0; i < len(foundUsers); i++ {
		users[i] = map[string]interface{}{"id": foundUsers[i].ID, "name": foundUsers[i].Name}
	}
	return c.JSON(http.StatusOK, users)
}
