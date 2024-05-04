package model

import "time"

type Message struct {
	ID        int       `json:"id"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
