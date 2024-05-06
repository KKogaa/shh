package port

import "github.com/KKogaa/shh/model"

type MessagePort interface {
	CreateMessageInChatroom(chatroomId int) (model.Message, error)
	GetMessagesByChatroom(chatroomId int) ([]model.Message, error)
}
