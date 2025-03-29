package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

// InsertUser inserts a new user into the database.
func InsertUser(db *sql.DB, table string, fields []string, values ...interface{}) (int64, error) {
	columnString := strings.Join(fields, ", ")
	valueString := strings.Repeat("?, ", len(fields))
	valueString = strings.TrimSuffix(valueString, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columnString, valueString)

	result, err := db.Exec(query, values...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last inserted ID for %v: %w", table, err)
	}
	return id, nil
}
