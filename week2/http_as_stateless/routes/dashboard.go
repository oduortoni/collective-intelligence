package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"has/session"
)

func Dashboard(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the 'auth_token' cookie from the request
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized, token missing", http.StatusUnauthorized)
			return
		}

		// validate the token and retrieve the user
		username := session.Remember(cookie.Value)
		if username == "" {
			username = "Guest"
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", username)
	}
}
