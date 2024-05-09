package service

import (
	"fmt"

	"github.com/KKogaa/shh/model"
	"github.com/KKogaa/shh/port"
)

type ChatroomService struct {
	ChatroomRepo port.ChatroomPort
}

func NewChatroomService(chatroomRepo port.ChatroomPort) ChatroomService {
	return ChatroomService{
		ChatroomRepo: chatroomRepo,
	}
}

func (c ChatroomService) GetChatroomByName(name string) (model.Chatroom, error) {
	existingChatroom, err := c.ChatroomRepo.GetChatroomByName(name)
	if err != nil {
		return model.Chatroom{}, err
	}

	return existingChatroom, nil
}

func (c ChatroomService) CreateChatroom(name string) (model.Chatroom, error) {
	existingChatroom, err := c.ChatroomRepo.GetChatroomByName(name)
	if err != nil {
		return model.Chatroom{}, err
	}

	if existingChatroom.ID != -1 {
		return existingChatroom, fmt.Errorf("error chatroom already exists")
	}

	chatroom, err := c.ChatroomRepo.CreateChatroom(name)
	if err != nil {
		return model.Chatroom{}, err
	}

	return chatroom, nil
}
