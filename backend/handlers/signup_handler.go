package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/johneliud/authentication_project/backend/database"
	"github.com/johneliud/authentication_project/backend/models"
	"github.com/johneliud/authentication_project/backend/utils"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-up" {
		log.Printf("Path not found: %s\n", r.URL.Path)
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("frontend/views/signup.html")
		if err != nil {
			log.Printf("Failed to parse template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to execute template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Printf("Failed to decode request body: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		defer r.Body.Close()

		if err := utils.ValidateUserFields(user.FirstName, user.LastName, user.Email, user.Password, user.ConfirmedPassword); err != nil {
			log.Printf("Failed to validate user fields: %v\n", err)
			response := Response{Success: false, Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		hashedPassword, err := utils.HashPassword([]byte(user.Password), 12)
		if err != nil {
			log.Printf("Failed to hash password: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		verificationCode := utils.GenerateVerificationCode()

		_, err = utils.InsertUser(database.DB, "users", []string{"first_name", "last_name", "email", "password_hash, verification_code"}, user.FirstName, user.LastName, user.Email, string(hashedPassword), verificationCode)
		if err != nil {
			log.Printf("Error adding user: %v\n", err)
			response := Response{Success: false, Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		if err := utils.SendVerificationEmail(user.Email, verificationCode); err != nil {
			log.Printf("Failed to send verification email: %v\n", err)
			response := Response{Success: false, Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		
		response := Response{Success: true, Message: "Registration successful"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return

	default:
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
