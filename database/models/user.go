package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null,uniqueIndex"`
	CreatedAt int64     `gorm:"autoCreateTime"`
	UpdatedAt int64     `gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
