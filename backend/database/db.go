package database

import (
	"database/sql"
	"fmt"
	"os"
)

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
