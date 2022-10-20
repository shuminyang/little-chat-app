package users

import (
	"chat-app/database"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user *UserModel) (uuid.UUID, error) {
	db := database.GetConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO users (first_name, last_name, email, password) values ($1, $2, $3, $4) RETURNING id;`

	var id uuid.UUID
	err := db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)

	return id, err

}

func Test() {}
