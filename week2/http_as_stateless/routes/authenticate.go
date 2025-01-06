package routes

import (
	"fmt"
	"net/http"

	"has/session"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		if username == "" {
			http.Error(w, "Username cannot be empty", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Successfully authenticated user: %s", username)
		token, ok := session.Save(username)
		if !ok {
			fmt.Fprintf(w, "Colud not log user in")
			return
		}

		// set the token as a cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			MaxAge:   3600,
		})
	} else {
		// If the request method is not POST, respond with an error
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
