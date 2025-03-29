package utils

import (
	"errors"
	"strings"
)

// ValidateUserFields validates the user fields.
func ValidateUserFields(firstName, lastName, email, password, confirmedPassword string) error {
	if strings.TrimSpace(firstName) == "" {
		return errors.New("first name cannot be empty")
	}

	if strings.TrimSpace(lastName) == "" {
		return errors.New("last name cannot be empty")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("password cannot be empty")
	}

	if strings.TrimSpace(confirmedPassword) == "" {
		return errors.New("confirmed password cannot be empty")
	}

	if strings.TrimSpace(password) != strings.TrimSpace(confirmedPassword) {
		return errors.New("passwords do not match")
	}
	return nil
}
