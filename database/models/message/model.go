package messages

import (
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	Message   string
	CreatedAt uuid.Time
	UpdatedAt time.Time
}
