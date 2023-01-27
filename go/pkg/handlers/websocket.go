package handlers

import (
	"Web-chat/pkg/services"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Websocket(c echo.Context) error {
	log.Print("It works!!!!!")
	req, err := http.NewRequest("GET", "http://localhost"+os.Getenv("SERVER_PORT")+"/jwt/id", nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	token := c.Param("Authorization")
	req.Header.Set("Authorization", "Bearer "+token[1:])
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	resp.Body.Close()

	id, err := strconv.Atoi(string(respBody))
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	ws, err := services.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if services.Clients[id] != nil {
		services.Clients[id].Requests++
		services.Clients[id].Conn.Close()
		services.Clients[id].Conn = ws
	} else {
		services.Clients[id] = &services.ClientConn{Conn: ws, Requests: 0}
	}
	defer services.CloseConn(id)

	for {
		var msg services.ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
			break
		}
		if msg.Text == "ping" || msg.ReceiverID == 0 {
			err := ws.WriteJSON(map[string]interface{}{"code": 1, "text": "pong"})
			if err != nil {
				c.Logger().Error(err)
				break
			}
			continue
		}
		msg.SenderID = id
		services.MessageChannel <- msg
	}
	return err
}
