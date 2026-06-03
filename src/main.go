package src

import (
	"github.com/labstack/echo/v4"
)



func Run() {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	
	server.Static("/", "./fe/dist")
	server.GET("/api/models", handleGetModels)
	server.POST("/api/chat", handleChat)

	server.Start(":3000")
}