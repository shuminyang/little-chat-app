package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
}

func main() {

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", homeHandler)
	router.POST("/createUser", createUser)

	router.Run()
}
