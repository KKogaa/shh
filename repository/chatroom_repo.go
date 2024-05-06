package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/KKogaa/shh/model"
	"github.com/jmoiron/sqlx"
)

type ChatroomRepo struct {
	db *sqlx.DB
}

func NewChatroomRepo(db *sqlx.DB) ChatroomRepo {
	return ChatroomRepo{
		db: db,
	}
}

func (c ChatroomRepo) GetChatroomByName(name string) (model.Chatroom, error) {
	chatroom := model.Chatroom{}
	err := c.db.Get(&chatroom, "select * from chatrooms where chatroom = :name", name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Chatroom{ID: -1, Chatroom: "none"}, nil
		}
		return chatroom, err
	}

	return chatroom, nil

}

func (c ChatroomRepo) CreateChatroom(name string) (model.Chatroom, error) {
	chatroom := model.Chatroom{
		Chatroom: name,
	}

	sql := `
    insert into chatrooms (chatroom) values (:chatroom) returning id
  `

	_, err := c.db.Exec(sql, chatroom.Chatroom)
	if err != nil {
		return chatroom, fmt.Errorf("error executing sql insert chatrooms %s", err)
	}

	return chatroom, nil
}
