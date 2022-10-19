package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	db *sql.DB
}

var databaseInstance *DatabaseManager

func InitDatabase() {
	if databaseInstance == nil {
		databaseUrl := "host=localhost user=postgres password=1234567890 dbname=little-chat-app port=5555 sslmode=disable"

		db, err := sql.Open("postgres", databaseUrl)

		if err != nil {
			panic(err)
		}

		err = db.Ping()

		if err != nil {
			panic(err)
		}

		databaseInstance = &DatabaseManager{
			db: db,
		}
	}

	fmt.Println("Successfully connected!")
}

func GetConnection() *sql.DB {
	if databaseInstance == nil {
		InitDatabase()
	}

	return databaseInstance.db
}
