package utils

import "errors"

// ValidateUserFields validates the user fields.
func ValidateUserFields(firstName, lastName, email, password, confirmedPassword string) error {
	if firstName == "" {
		return errors.New("first name cannot be empty")
	}

	if lastName == "" {
		return errors.New("last name cannot be empty")
	}

	if email == "" {
		return errors.New("email cannot be empty")
	}

	if password == "" {
		return errors.New("password cannot be empty")
	}

	if confirmedPassword == "" {
		return errors.New("confirmed password cannot be empty")
	}

	if password != confirmedPassword {
		return errors.New("passwords do not match")
	}
	return nil
}
