package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KKogaa/shh/service"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type MessageWS struct {
	chatroomService service.ChatroomService
	messageService  service.MessageService
	chatrooms       *Chatrooms
}

func NewMessageWS(chatroomService service.ChatroomService,
	messageService service.MessageService) MessageWS {
	return MessageWS{
		chatrooms:       NewChatrooms(),
		chatroomService: chatroomService,
		messageService:  messageService,
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

	m.chatrooms.AddNewConnection(conn)

	defer m.chatrooms.RemoveConnection(conn)

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

		//TODO: refactor this
		if !m.chatrooms.Exists(decodedMsg.Chatroom) {
			m.chatrooms.InitializeChatroom(decodedMsg.Chatroom)
			m.chatroomService.CreateChatroom(decodedMsg.Chatroom)
		}

		m.chatrooms.AddNewConnectionToChatroom(decodedMsg.Chatroom, conn)

		chatroom, err := m.chatroomService.GetChatroomByName(decodedMsg.Chatroom)
		if err != nil {
			return err
		}
		_, err = m.messageService.CreateMessage(chatroom.ID,
			decodedMsg.Msg, decodedMsg.User)

		if err != nil {
			return err
		}

		err = m.chatrooms.Broadcast(decodedMsg.Chatroom, decodedMsg)
		if err != nil {
			return err
		}
	}
}
