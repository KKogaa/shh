package main

import (
	"github.com/KKogaa/shh/controller/rest"
	"github.com/KKogaa/shh/repository"
	"github.com/KKogaa/shh/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo) {
	// messageWS := ws.NewMessageWS()
	// e.Static("/", "./index.html")
	// e.GET("/ws", messageWS.GetMessages)

	//WIRING
	messageRepo := repository.NewMessageRepo()
	messageService := service.NewMessageService(messageRepo)
	messages := rest.NewMessagesREST(messageService)
	e.GET("/messages", messages.ListMessages)
	e.POST("/messages", messages.CreateMessage)

}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Routes(e)
	e.Logger.Fatal(e.Start(":1323"))

	// TODO: add gracefull stop

}
