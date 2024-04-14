package rest

import (
	"net/http"

	"github.com/KKogaa/shh/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type MessagesREST struct {
	messageService service.MessageService
}

func NewMessagesREST(messageService service.MessageService) MessagesREST {
	return MessagesREST{
		messageService: messageService,
	}
}

// dto
// list of
// id user
// message
// timestamp
// TODO: in the end this response is encrypted
func (m MessagesREST) ListMessages(c echo.Context) error {
	messages, err := m.messageService.ListMessages()
	if err != nil {
		log.Errorf("error: ", err)
	}
	return c.JSON(http.StatusOK, messages)
}

// dto
// id user
// message
// timestamp from server

func (m MessagesREST) CreateMessage(c echo.Context) error {
	return nil
}
