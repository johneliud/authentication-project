package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/johneliud/authentication_project/backend/config"
)

// LogoutHandler handles the logout page.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, "session")
	if err != nil {
		log.Printf("Failed to get session: %v\n", err)
		response := Response{Success: false, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Clear session
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	if err := session.Save(r, w); err != nil {
		log.Printf("Failed to save session: %v\n", err)
		response := Response{Success: false, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{Success: false, Message: "Successfully logged out"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
}
