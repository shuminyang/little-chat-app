package services

import (
	"chat-app/database/models"
)

func CreateUser(firstName *string, lastName *string, email *string) (id string, err error) {

	user := models.User{
		FirstName: *firstName,
		LastName:  *lastName,
		Email:     *email,
	}

	result, err := models.CreateUser(&user)

	if err != nil {
		return "", err

	}

	return result.String(), nil
}
