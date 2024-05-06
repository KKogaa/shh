package port

import "github.com/KKogaa/shh/model"

type ChatroomPort interface {
	GetChatroomByName(chatroom string) (model.Chatroom, error)
	CreateChatroom(chatroom string) (model.Chatroom, error)
}
