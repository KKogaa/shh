package service

import (
	"time"

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

func (m MessageService) CreateMessage(chatroomId int64,
	payload string, username string) (model.Message, error) {

	message := model.Message{
		Payload:    payload,
		Username:   username,
		CreatedAt:  time.Now(),
		ChatroomId: chatroomId,
	}

	message, err := m.MessageRepo.CreateMessageInChatroom(message.ChatroomId, message)
	if err != nil {
		return message, err
	}

	return message, nil
}
