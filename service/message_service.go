package service

import (
	"github.com/KKogaa/shh/model"
	"github.com/KKogaa/shh/port"
)

type MessageService struct {
	MessageRepo  port.MessagePort
	ChatroomRepo port.ChatroomPort
}

func NewMessageService(messageRepo port.MessagePort,
	chatroomRepo port.ChatroomPort) MessageService {
	return MessageService{
		MessageRepo:  messageRepo,
		ChatroomRepo: chatroomRepo,
	}
}

func (m MessageService) ListMessages(chatroomName string) ([]model.Message, error) {

	chatroom, err := m.ChatroomRepo.GetChatroomByName(chatroomName)
	if err != nil {
		return nil, err
	}

	messages, err := m.MessageRepo.GetMessagesByChatroom(chatroom.ID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
