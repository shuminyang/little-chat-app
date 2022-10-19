package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
}

type CreateMessageInterface struct {
	UserId  string `json:"userId" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func CreateMessage(c *gin.Context) {
	var msgInterface CreateMessageInterface

	if err := c.ShouldBindJSON(&msgInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
