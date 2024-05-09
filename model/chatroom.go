package model

type Chatroom struct {
	ID       int64  `json:"id" db:"id"`
	Chatroom string `json:"chatroom" db:"chatroom"`
}
