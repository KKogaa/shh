package ws

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type Chatrooms struct {
	chatrooms      map[string]*ConnectedClients
	mutex          sync.Mutex
	connectionPool *ConnectedClients
}

func NewChatrooms() *Chatrooms {
	return &Chatrooms{
		chatrooms: make(map[string]*ConnectedClients),
		mutex:     sync.Mutex{},
		connectionPool: &ConnectedClients{
			clients: make(map[*websocket.Conn]bool),
			mutex:   sync.Mutex{},
		},
	}
}

func (c *Chatrooms) InitializeChatroom(name string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.chatrooms[name]; exists {
		return errors.New("Chatroom already exists")
	}

	clients := ConnectedClients{
		clients: make(map[*websocket.Conn]bool),
		mutex:   sync.Mutex{},
	}
	c.chatrooms[name] = &clients
	return nil
}

func (c *Chatrooms) AddNewConnectionToChatroom(chatroomName string,
	conn *websocket.Conn) {
	_, exists := c.chatrooms[chatroomName].clients[conn]
	if !exists {
		c.chatrooms[chatroomName].add(conn)
	}
}

func (c *Chatrooms) Exists(name string) bool {
	_, exists := c.chatrooms[name]
	return exists
}

func (c *Chatrooms) Broadcast(chatroomName string, message Message) error {
	clients, exists := c.chatrooms[chatroomName]
	if !exists {
		return errors.New("chatroom doesn't exist")
	}

	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	clients.broadcast(jsonMsg)
	return nil
}

func (c *Chatrooms) AddNewConnection(conn *websocket.Conn) {
	c.connectionPool.add(conn)
}

func (c *Chatrooms) RemoveConnection(conn *websocket.Conn) {
	c.connectionPool.add(conn)
	//TODO: also search this connection recursively on all chatrooms
}
