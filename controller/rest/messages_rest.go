package rest

import (
	"log"
	"net/http"

	"github.com/KKogaa/shh/service"
	"github.com/labstack/echo/v4"
)

type MessagesREST struct {
	messageService service.MessageService
}

func NewMessagesREST(messageService service.MessageService) MessagesREST {
	return MessagesREST{
		messageService: messageService,
	}
}

func (m MessagesREST) ListMessages(c echo.Context) error {
	chatroomName := c.Param("chatroomName")
	messages, err := m.messageService.ListMessages(chatroomName)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error listing messages")
	}
	return c.JSON(http.StatusOK, messages)
}

func (m MessagesREST) CreateMessage(c echo.Context) error {
	return nil
}
