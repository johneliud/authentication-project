package routes

import (
	"net/http"

	"github.com/johneliud/authentication_project/backend/handlers"
)

// InitRoutes initializes the routes for the server.
func InitRoutes() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend", fs))

	http.HandleFunc("/sign-up", handlers.SignupHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)
	http.HandleFunc("/sign-in", handlers.SigninHandler)
	http.HandleFunc("/", handlers.HomeHandler)
}
