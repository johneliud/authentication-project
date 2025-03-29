package database

import (
	"database/sql"
	"fmt"
	"os"
)

// InitDB initializes the database and returns a pointer to the database.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/authentication.db")
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	err = executeSchema(db)
	if err != nil {
		return nil, fmt.Errorf("could not execute schema: %v", err)
	}
	return db, nil
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
