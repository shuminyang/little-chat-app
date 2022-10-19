package message

import (
	"chat-app/database"
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateMessage(userId *string, content *string) (uuid.UUID, error) {
	db := database.GetConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO messages (user_id, content) values ($1, $2) RETURNING id;`

	var id uuid.UUID
	err := db.QueryRow(sqlStatement, userId, content).Scan(&id)

	return id, err
}

func GetMessages() (messages []MessageModel, err error) {
	db := database.GetConnection()
	defer db.Close()

	sqlStatement := `select id, user_id, content, created_at, updated_at from messages;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var message MessageModel
		err = rows.Scan(&message.ID, &message.UserId, &message.Content, &message.CreatedAt, &message.UpdatedAt)

		if err != nil {
			panic(err)
		}

		messages = append(messages, message)
	}

	if err != nil {
		return nil, err
	}

	return messages, nil

}
