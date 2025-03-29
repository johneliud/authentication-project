package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// InitDB initializes the database and returns a pointer to the database.
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/authentication.db")
	if err != nil {
		log.Printf("Could not open database: %v\n", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Could not ping database: %v\n", err)
		return nil
	}

	err = executeSchema(db)
	if err != nil {
		log.Printf("Could not execute schema: %v\n", err)
		return nil
	}
	log.Println("Database initialized successfully")
	return db
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
