package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDatabase() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNECTION"))
	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}
	DB = db
}

func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("Could not close database connection: %v", err)
	}
}
