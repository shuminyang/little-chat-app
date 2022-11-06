package main

import (
	apiAuth "chat-app/api/auth"
	apiMessages "chat-app/api/messages"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	router.SetTrustedProxies([]string{"localhost"})

	router.POST("/register", apiAuth.Register)
	router.POST("/sign-in", apiAuth.SignIn)

	router.POST("/create-message", apiMessages.CreateMessage)
	router.GET("/get-messages", apiMessages.GetMessages)

	router.Run()
}
