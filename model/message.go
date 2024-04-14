package model

import "time"

type Message struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
