package port

import "github.com/KKogaa/shh/model"

type MessagePort interface {
	GetMessages() ([]model.Message, error)
	CreateMessage() (model.Message, error)
}
