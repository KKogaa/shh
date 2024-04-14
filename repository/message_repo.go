package repository

import (
	"time"

	"github.com/KKogaa/shh/model"
)

type MessageRepo struct{}

func NewMessageRepo() MessageRepo {
	return MessageRepo{}
}

func (m MessageRepo) GetMessages() ([]model.Message, error) {
	//TODO: get from database encrypted
	message1 := model.Message{
		ID:        1,
		Message:   "hello",
		UserId:    1,
		Timestamp: time.Now(),
	}

	message2 := model.Message{
		ID:        2,
		Message:   "hello2",
		UserId:    2,
		Timestamp: time.Now(),
	}

	messages := []model.Message{message1, message2}

	return messages, nil
}

func (m MessageRepo) CreateMessage() (model.Message, error) {
	message1 := model.Message{
		ID:        1,
		Message:   "hello",
		UserId:    1,
		Timestamp: time.Now(),
	}

	return message1, nil
}
