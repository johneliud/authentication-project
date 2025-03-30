package handlers

import (
	"log"
	"net/http"
	"html/template"
)

// HomeHandler handles the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("Path not found: %s\n", r.URL.Path)
		http.Error(w, "Path not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("frontend/views/home.html")
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
}
