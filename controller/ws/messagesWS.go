package ws

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type MessageWS struct {
}

func NewMessageWS() MessageWS {
	return MessageWS{}
}

func (m MessageWS) GetMessages(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			log.Println("HERE")
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
