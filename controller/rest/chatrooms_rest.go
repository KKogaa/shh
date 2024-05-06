package rest

import (
	"log"
	"net/http"

	"github.com/KKogaa/shh/model"
	"github.com/KKogaa/shh/service"
	"github.com/labstack/echo/v4"
)

type ChatroomsREST struct {
	chatroomService service.ChatroomService
}

func NewChatroomsREST(chatroomService service.ChatroomService) ChatroomsREST {
	return ChatroomsREST{
		chatroomService: chatroomService,
	}
}

type ChatroomDTO struct {
	Chatroom string `json:"chatroom"`
}

func chatroomToDTO(chatroom model.Chatroom) ChatroomDTO {
	return ChatroomDTO{
		Chatroom: chatroom.Chatroom,
	}
}

func (c ChatroomsREST) CreateChatroom(cont echo.Context) error {
	chatroomName := cont.Param("chatroomName")
	chatroom, err := c.chatroomService.CreateChatroom(chatroomName)
	if err != nil {
		log.Println(err)
		return cont.JSON(http.StatusInternalServerError, "Error creating chatroom")
	}
	return cont.JSON(http.StatusOK, chatroomToDTO(chatroom))

}
