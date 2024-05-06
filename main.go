package main

import (
	"github.com/KKogaa/shh/controller/rest"
	"github.com/KKogaa/shh/controller/ws"
	"github.com/KKogaa/shh/db"
	"github.com/KKogaa/shh/repository"
	"github.com/KKogaa/shh/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {
	db := db.CreateDB()

	messageWS := ws.NewMessageWS()
	e.GET("/ws", messageWS.GetMessages)

	messageRepo := repository.NewMessageRepo(db)
	chatroomRepo := repository.NewChatroomRepo(db)
	messageService := service.NewMessageService(messageRepo, chatroomRepo)
	messages := rest.NewMessagesREST(messageService)

	chatroomService := service.NewChatroomService(chatroomRepo)
	chatrooms := rest.NewChatroomsREST(chatroomService)

	e.GET("/messages/chatrooms/:chatroomName", messages.ListMessages)
	e.POST("/messages", messages.CreateMessage)
	e.POST("/chatrooms/:chatroomName", chatrooms.CreateChatroom)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
