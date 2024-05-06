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

type ConnectedClients struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

func (m *ConnectedClients) add(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clients[conn] = true
}

func (m *ConnectedClients) remove(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.clients, conn)
}

func (m *ConnectedClients) broadcast(message []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	log.Printf("connected clients: %d", len(m.clients))
	for conn := range m.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error sending message to client:", err)
		}
	}
}

func (m MessageWS) SendMessage(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}
	return nil
}

type Message struct {
	Msg  string `json:"msg"`
	User string `json:"user"`
}

func (m MessageWS) GetMessages(c echo.Context) error {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow any origin for WebSocket connection
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

		jsonMsg, err := json.Marshal(decodedMsg)
		if err != nil {
			return err
		}

		m.clients.broadcast(jsonMsg)
	}
}
