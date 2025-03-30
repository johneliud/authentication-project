package routes

import (
	"net/http"

	"github.com/johneliud/authentication_project/backend/handlers"
	"github.com/johneliud/authentication_project/backend/middleware"
)

// InitRoutes initializes the routes for the server.
func InitRoutes() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/frontend/", http.StripPrefix("/frontend", fs))

	http.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		middleware.SessionMiddleware(http.HandlerFunc(handlers.SignupHandler)).ServeHTTP(w, r)
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		middleware.SessionMiddleware(http.HandlerFunc(handlers.VerifyHandler)).ServeHTTP(w, r)
	})

	http.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		middleware.SessionMiddleware(http.HandlerFunc(handlers.SigninHandler)).ServeHTTP(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.SessionMiddleware(http.HandlerFunc(handlers.HomeHandler)).ServeHTTP(w, r)
	})
	
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		middleware.SessionMiddleware(http.HandlerFunc(handlers.LogoutHandler)).ServeHTTP(w, r)
	})
}
