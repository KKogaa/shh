package repository

import (
	"fmt"

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

func (m MessageRepo) GetMessagesByChatroom(chatroomId int64) ([]model.Message, error) {
	sql := `
    select * 
    from messages where chatroom_id = :chatroomId
    order by created_at
  `
	messages := []model.Message{}
	err := m.db.Select(&messages, sql, chatroomId)

	if err != nil {
		return messages, fmt.Errorf("error executing sql select messages %s", err)
	}

	return messages, nil
}

func (m MessageRepo) CreateMessageInChatroom(chatroomId int64,
	message model.Message) (model.Message, error) {

	sql := ` insert into messages (username, chatroom_id, payload, created_at) 
  values ($1, $2, $3, $4) `

	result, err := m.db.Exec(sql, message.Username, chatroomId, message.Payload,
		message.CreatedAt)

	if err != nil {
		return message, fmt.Errorf("error executing sql insert messages %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return message, fmt.Errorf("error getting id from messages %s", err)
	}

	message.ID = id

	return message, nil
}
