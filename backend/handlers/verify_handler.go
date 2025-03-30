package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/johneliud/authentication_project/backend/database"
	"github.com/johneliud/authentication_project/backend/middleware"
)

// VerifyHandler handles the verification page.
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/verify" {
		log.Printf("Path not found: %s\n", r.URL.Path)
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("frontend/views/verify.html")
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
		session, err := middleware.Store.Get(r, "session")
		if err != nil {
			log.Printf("Failed to get session: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		userEmail, ok := session.Values["email"].(string)
		if !ok || userEmail == "" {
			log.Println("User email not found in session")
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: "User email not found in session"}
			json.NewEncoder(w).Encode(response)
			return
		}

		var verificationData struct {
			VerificationCode string `json:"verification_code"`
		}

		if err := json.NewDecoder(r.Body).Decode(&verificationData); err != nil {
			log.Printf("Failed to decode request body: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		defer r.Body.Close()

		var dbCode string

		err = database.DB.QueryRow("SELECT verification_code FROM users WHERE email = ?", userEmail).Scan(&dbCode)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("Email %q found: %v\n", userEmail, err)
				w.Header().Set("Content-Type", "application/json")
				response := Response{Success: false, Message: err.Error()}
				json.NewEncoder(w).Encode(response)
				return
			} else {
				log.Printf("Failed to query database: %v\n", err)
				w.Header().Set("Content-Type", "application/json")
				response := Response{Success: false, Message: err.Error()}
				json.NewEncoder(w).Encode(response)
				return
			}
		}

		if dbCode != verificationData.VerificationCode {
			log.Printf("Verification code %q does not match: %v\n", verificationData.VerificationCode, err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: "Verification code does not match"}
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = database.DB.Exec("UPDATE users SET verified = 1 WHERE email = ?", userEmail)
		if err != nil {
			log.Printf("Failed to update user: %v\n", err)
			w.Header().Set("Content-Type", "application/json")
			response := Response{Success: false, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		session.Values["authenticated"] = true
		if err := session.Save(r, w); err != nil {
			log.Printf("Failed to save session: %v\n", err)
			response := Response{Success: false, Message: err.Error()}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := Response{Success: true, Message: "Email successfully verified"}
		json.NewEncoder(w).Encode(response)

	default:
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
