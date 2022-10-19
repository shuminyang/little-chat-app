package services

import (
	"chat-app/database"
	"chat-app/database/models"
	"fmt"
)

func CreateUser(name *string, email *string) (id string, err error) {
	db := database.GetConnection()
	user := models.User{
		Name:  *name,
		Email: *email,
	}

	fmt.Println("this is the user", user)

	result := db.Create(&user)

	if result.Error != nil {
		return "", result.Error

	}

	return user.ID.String(), nil
}
