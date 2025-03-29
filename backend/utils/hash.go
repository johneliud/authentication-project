package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword uses bcrypt to hash a password with a cost of 12 for secure storage in the database.
func HashPassword(password []byte, cost int) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	return string(hashedPassword), nil
}
