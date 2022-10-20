package user

import (
	"chat-app/database/models/users"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserDomain struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPass), err

}

func CreateUser(createUserDomain *CreateUserDomain) (id string, err error) {

	hashedPassword, err := hashPassword(createUserDomain.Password)

	if err != nil {
		return "", err
	}

	user := users.UserModel{
		FirstName: createUserDomain.FirstName,
		LastName:  createUserDomain.LastName,
		Email:     createUserDomain.Email,
		Password:  hashedPassword,
	}

	result, err := users.CreateUser(&user)

	if err != nil {
		return "", err

	}

	return result.String(), nil
}
