package routes

import (
	"html/template"
	"net/http"

	"has/session"
)

func Login(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
		* First see if a cookie already exists.
		* No need to login if user is already logged in
		* redirect to dashboard if logged in
		 */
		// get the 'auth_token' cookie from the request
		token := ""
		cookie, err := r.Cookie("auth_token")
		if err == nil {
			token = cookie.Value
		}
		// validate the token and retrieve the user
		username := session.Remember(token)
		// fmt.Println("Username: ", username)
		if username != "" {
			// already logged in
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		tmpl.ExecuteTemplate(w, "login.html", nil)
	}
}
