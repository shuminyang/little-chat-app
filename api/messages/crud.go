package messages

import (
	"net/http"

	"chat-app/database/models/message"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	messages, err := message.GetMessages()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

type CreateMessageInterface struct {
	UserId  string `json:"userId" binding:"required,uuid"`
	Content string `json:"content" binding:"required"`
}

func CreateMessage(c *gin.Context) {
	var msgInterface CreateMessageInterface

	if err := c.ShouldBindJSON(&msgInterface); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := message.CreateMessage(&msgInterface.UserId, &msgInterface.Content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result.String())
}
