package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type MessageWS struct {
	clients *ConnectedClients
}

func NewMessageWS() MessageWS {

	clients := ConnectedClients{
		clients: make(map[*websocket.Conn]bool),
		mutex:   sync.Mutex{},
	}
	return MessageWS{
		clients: &clients,
	}
}

func (m MessageWS) SendMessage(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}
	return nil
}

func (m MessageWS) GetMessages(c echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer conn.Close()

	m.clients.add(conn)
	defer m.clients.remove(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		log.Println(string(message))

		var decodedMsg Message
		err = json.Unmarshal(message, &decodedMsg)
		if err != nil {
			return err
		}

		//TODO: for every message that is read here persist in database
		//TODO: map request message to persistance message

		jsonMsg, err := json.Marshal(decodedMsg)
		if err != nil {
			return err
		}

		m.clients.broadcast(jsonMsg)
	}
}
