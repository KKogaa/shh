package rest

import "github.com/labstack/echo/v4"

type MessagesREST struct {
}

func NewMessagesREST() MessagesREST {
	return MessagesREST{}
}

func (m MessagesREST) ListMessages(c echo.Context) error {
	return nil
}

func (m MessagesREST) CreateMessage(c echo.Context) error {
	return nil
}
