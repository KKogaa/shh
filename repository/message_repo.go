package repository

import (
	"github.com/KKogaa/shh/model"
	"github.com/jmoiron/sqlx"
)

type MessageRepo struct {
	db *sqlx.DB
}

func NewMessageRepo(db *sqlx.DB) MessageRepo {
	return MessageRepo{
		db: db,
	}
}

func (m MessageRepo) GetMessagesByChatroom(chatroomId int) ([]model.Message, error) {
	messages := []model.Message{}
	err := m.db.Select(&messages, "select * from messages where chatroom_id = :chatroomId",
		chatroomId)

	if err != nil {
		return messages, err
	}

	return messages, nil
}

func (m MessageRepo) CreateMessageInChatroom(chatroomId int) (model.Message, error) {
	return model.Message{}, nil
}
