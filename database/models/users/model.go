package users

import (
	"chat-app/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user *User) (uuid.UUID, error) {
	db := database.GetConnection()
	sqlStatement := `INSERT INTO users (first_name, last_name, email) values ($1, $2, $3) RETURNING id;`

	var id uuid.UUID
	err := db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email).Scan(&id)

	return id, err

}

func Test() {}
