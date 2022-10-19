package models

import (
	"chat-app/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null,uniqueIndex"`
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
