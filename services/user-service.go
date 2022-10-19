package services

import (
	"chat-app/database/models/users"
)

func CreateUser(firstName *string, lastName *string, email *string) (id string, err error) {

	user := users.User{
		FirstName: *firstName,
		LastName:  *lastName,
		Email:     *email,
	}

	result, err := users.CreateUser(&user)

	if err != nil {
		return "", err

	}

	return result.String(), nil
}
