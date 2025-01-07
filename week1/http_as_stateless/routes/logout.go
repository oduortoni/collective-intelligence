package routes

import (
	"html/template"
	"net/http"

	"has/session"
)

func Logout(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := ""
		// get the 'auth_token' cookie from the request
		cookie, err := r.Cookie("auth_token")
		if err == nil {
			token = cookie.Value
		}

		// validate the token and retrieve the user
		username := session.Remember(token)
		if username == "" {
			username = "Guest"
		} else {
			session.Delete(token)
			username = "logged_out_user"
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", username)
	}
}
