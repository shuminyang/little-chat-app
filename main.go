package main

import (
	"chat-app/api"
	"chat-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
}

func main() {
	database.InitDatabase()

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", homeHandler)
	router.POST("/create-user", api.CreateUser)

	router.Run()
}
