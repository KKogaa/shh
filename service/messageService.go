package service

import (
	"fmt"

	"github.com/KKogaa/shh/model"
	"github.com/KKogaa/shh/port"
)

type MessageService struct {
	MessageRepo port.MessagePort
}

func NewMessageService(messageRepo port.MessagePort) MessageService {
	return MessageService{
		MessageRepo: messageRepo,
	}
}

func (m MessageService) ListMessages() ([]model.Message, error) {

	messages, err := m.MessageRepo.GetMessages()
	if err != nil {
		return nil, fmt.Errorf("error obtaining messages: %s", err)
	}

	return messages, nil
}
