package models

type User struct {
	ID                int    `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
	Verified          bool   `json:"verified"`
}
