package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"has/session"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// try parsing multipart form first
		err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
		if err != nil {
			// if not multipart, try regular form
			err = r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form data", http.StatusBadRequest)
				return
			}
		}

		// try getting username from different sources
		username := r.FormValue("username")
		if username == "" {
			// if still empty, try to read JSON body
			var data struct {
				Username string `json:"username"`
			}
			if err := json.NewDecoder(r.Body).Decode(&data); err == nil {
				username = data.Username
			}
		}
		// fmt.Printf("Username received: %s\n", username)

		if username == "" {
			http.Error(w, "Username cannot be empty", http.StatusBadRequest)
			return
		}

		token, ok := session.Save(username)
		if !ok {
			fmt.Fprintf(w, "Could not log user in")
			return
		}

		// set the token as a cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			HttpOnly: true, // prevent javascript from accessing the cookie for security reasons
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
			MaxAge:   3600,
		})

		// send the response to javascript
		response := map[string]string{
			"status":   "success",
			"message":  "Login successful",
			"redirect": "/dashboard", // the page that javascript will redirect to after successful login
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// If the request method is not POST, respond with an error
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
