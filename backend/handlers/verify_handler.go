package handlers

import (
	"log"
	"net/http"
	"text/template"
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
		}

	default:
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
