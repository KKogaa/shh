package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

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
