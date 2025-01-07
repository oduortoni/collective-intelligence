package routes

import (
	"html/template"
	"net/http"

	"has/session"
)

func Logout(tmpl *template.Template) http.HandlerFunc {
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
		} else {
			session.Delete(cookie.Value)
			username = "logged_out_user"
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", username)
	}
}
