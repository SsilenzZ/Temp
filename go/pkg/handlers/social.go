package handlers

import (
	"Web-chat/pkg/db/repositories/social"
	"Web-chat/pkg/db/repositories/user"
	"Web-chat/pkg/requests"
	"Web-chat/pkg/services"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SocialHandler struct {
}

func (s SocialHandler) Block(c echo.Context) error {
	var user_ requests.Social
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	err := json.NewDecoder(c.Request().Body).Decode(&user_)
	if err != nil {
		return err
	}

	success := social.Repository.Block(claims.ID, user_.ID)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (s SocialHandler) Friend(c echo.Context) error {
	var user_ requests.Social
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	err := json.NewDecoder(c.Request().Body).Decode(&user_)
	if err != nil {
		return err
	}

	success := social.Repository.Friend(claims.ID, user_.ID)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (s SocialHandler) Unfriend(c echo.Context) error {
	var user_ requests.Social
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	err := json.NewDecoder(c.Request().Body).Decode(&user_)
	if err != nil {
		return err
	}

	social.Repository.Unfriend(claims.ID, user_.ID)
	return c.JSON(http.StatusOK, nil)
}

func (s SocialHandler) Unblock(c echo.Context) error {
	var user_ requests.Social
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	err := json.NewDecoder(c.Request().Body).Decode(&user_)
	if err != nil {
		return err
	}

	social.Repository.Unblock(claims.ID, user_.ID)
	return c.JSON(http.StatusOK, nil)
}

func (s SocialHandler) GetFriends(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	var err error

	friends := social.Repository.GetFriends(claims.ID)

	res := make([]map[string]interface{}, len(friends))

	for i := 0; i < len(friends); i++ {
		res[i] = make(map[string]interface{})
		res[i]["id"], res[i]["name"], err = user.Repository.GetById(friends[i].User_ID2)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, res)
}

func (s SocialHandler) GetBlacklist(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*services.JWTCustomClaims)

	var err error

	friends := social.Repository.GetBlocked(claims.ID)

	res := make([]map[string]interface{}, len(friends))

	for i := 0; i < len(friends); i++ {
		res[i] = make(map[string]interface{})
		res[i]["id"], res[i]["name"], err = user.Repository.GetById(friends[i].User_ID2)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, res)
}
