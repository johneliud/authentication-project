package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/johneliud/authentication_project/backend/models"
	"github.com/johneliud/authentication_project/backend/utils"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var db *sql.DB

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

		if err := r.ParseForm(); err != nil {
			log.Printf("Failed to parse form: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		user.FirstName = strings.TrimSpace(r.FormValue("first_name"))
		user.LastName = strings.TrimSpace(r.FormValue("last_name"))
		user.Email = strings.TrimSpace(r.FormValue("email"))
		user.Password = strings.TrimSpace(r.FormValue("password"))
		user.ConfirmedPassword = strings.TrimSpace(r.FormValue("confirmed_password"))

		if err := utils.ValidateUserFields(user.FirstName, user.LastName, user.Email, user.Password, user.ConfirmedPassword); err != nil {
			log.Printf("Failed to validate user fields: %v\n", err)
			response := Response{Success: false, Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		hashedPassword, err := utils.HashPassword([]byte(user.Password), 12)
		if err != nil {
			log.Printf("Failed to hash password: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = utils.InsertUser(db, "users", []string{"first_name", "last_name", "email", "user_password"}, user.FirstName, user.FirstName, user.Email, string(hashedPassword))
		if err != nil {
			log.Printf("Error adding user: %v\n", err)
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		response := Response{Success: true, Message: "Success"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		http.Redirect(w, r, "/verify", http.StatusSeeOther)
		return

	default:
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
