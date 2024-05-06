package model

import "time"

type Message struct {
	ID         int       `json:"id" db:"id"`
	ChatroomId int       `json:"chatroom_id" db:"chatroom_id"`
	Username   string    `json:"username" db:"username"`
	Payload    string    `json:"payload" db:"payload"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
