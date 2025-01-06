package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"has/session"
)

func Dashboard(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the 'auth_token' cookie from the request
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized, token missing", http.StatusUnauthorized)
			return
		}

		fmt.Println("User: ", cookie.Name)
		fmt.Println("Value: ", cookie.Value)
		// Validate the token and retrieve the user
		username := session.Remember(cookie.Value)
		if username == "" {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", username)
	}
}
