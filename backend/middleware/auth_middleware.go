package middleware

import (
	"net/http"

	"github.com/johneliud/authentication-project/backend/config"
)

// SessionMiddleware is a middleware that checks if the user is authenticated and redirects to the sign-in page if not.
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exempt signup and signin routes from requiring a session
		if r.URL.Path == "/sign-up" || r.URL.Path == "/sign-in" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := config.Store.Get(r, "session")
		if err != nil {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
