package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/johneliud/authentication_project/backend/config"
	"github.com/johneliud/authentication_project/backend/database"
	"github.com/johneliud/authentication_project/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// SigninHandler handles the sign-in page.
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-in" {
		log.Printf("Path not found: %s\n", r.URL.Path)
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("frontend/views/signin.html")
		if err != nil {
			log.Printf("Failed to parse template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to execute template: %v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			log.Printf("Failed to decode request body: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		defer r.Body.Close()

		var user models.User

		err := database.DB.QueryRow(
			"SELECT id, first_name, last_name, password_hash, verified FROM users WHERE email = ?",
			credentials.Email,
		).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Verified)

		if err != nil {
			log.Printf("User not found: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		if !user.Verified {
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: "Please verify your email before signing in"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		session, err := config.Store.Get(r, "session")
		if err != nil {
			log.Printf("Failed to get session: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		session.Values["authenticated"] = true
		session.Values["user_id"] = user.ID
		session.Values["email"] = credentials.Email
		session.Values["first_name"] = user.FirstName
		session.Values["last_name"] = user.LastName

		if err := session.Save(r, w); err != nil {
			log.Printf("Failed to save session: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := Response{Success: true, Message: "Sign in successful"}
		json.NewEncoder(w).Encode(response)
		return
	default:
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
