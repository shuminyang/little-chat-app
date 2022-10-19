package users

import (
	"net/http"

	"chat-app/services"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func CreateUser(c *gin.Context) {
	var user UserApi

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateUser(&user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}
