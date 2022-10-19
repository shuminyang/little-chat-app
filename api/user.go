package api

import (
	"net/http"
	"net/mail"

	"chat-app/services"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func CreateUser(c *gin.Context) {
	var user UserApi

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := services.CreateUser(&user.Name, &user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)

}
