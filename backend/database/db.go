package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database and returns a pointer to the database.
func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./data/authentication.db")
	if err != nil {
		log.Printf("Could not open database: %v", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Printf("Could not ping database: %v", err)
		return
	}

	err = executeSchema(DB)
	if err != nil {
		log.Printf("Could not execute schema: %v", err)
		return
	}
	log.Println("Database initialized successfully")
}

// executeSchema reads the schema file and executes it on the database.
func executeSchema(db *sql.DB) error {
	content, err := os.ReadFile("backend/database/schema.sql")
	if err != nil {
		return fmt.Errorf("could not read schema file: %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("could not execute schema: %w", err)
	}
	return nil
}
