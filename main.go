package main

import (
	apiMessages "chat-app/api/messages"
	apiUser "chat-app/api/users"
	"chat-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.POST("/create-user", apiUser.CreateUser)
	router.POST("/create-message", apiMessages.CreateMessage)
	router.GET("/get-messages", apiMessages.GetMessages)

	router.Run()
}
