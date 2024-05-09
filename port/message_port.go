package port

import "github.com/KKogaa/shh/model"

type MessagePort interface {
	CreateMessageInChatroom(chatroomId int64, message model.Message) (model.Message, error)
	GetMessagesByChatroom(chatroomId int64) ([]model.Message, error)
}
