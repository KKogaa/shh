package model

type Chatroom struct {
	ID       int    `json:"id" db:"id"`
	Chatroom string `json:"chatroom" db:"chatroom"`
}
