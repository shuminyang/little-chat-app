package database

import (
	"fmt"

	"chat-app/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	db *gorm.DB
}

var databaseInstance *DatabaseManager

func InitDatabase() {
	if databaseInstance == nil {
		databaseUrl := "host=localhost user=postgres password=1234567890 dbname=little-chat-app port=5555 sslmode=disable"

		db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		databaseInstance = &DatabaseManager{
			db: db,
		}
	}

	fmt.Println("Successfully connected!")
}

func GetConnection() *gorm.DB {
	if databaseInstance == nil {
		InitDatabase()
	}

	return databaseInstance.db
}

func RunAutoMigration() {
	db := GetConnection()

	db.AutoMigrate(&models.User{})

	fmt.Println("Schema is up to date.")
}
