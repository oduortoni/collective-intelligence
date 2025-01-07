package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"has/routes"
)

var (
	tmpl *template.Template
	err  error
	port int
)

func init() {
	// parse all html files in the frontend and its subdirectories beforehand // optimization
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port = 9000
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Server running on http://localhost:%d\n", port)

	http.HandleFunc("/", routes.Index(tmpl))
	http.HandleFunc("/login", routes.Login(tmpl))
	http.HandleFunc("/authenticate", routes.Authenticate)
	http.HandleFunc("/dashboard", routes.Dashboard(tmpl))
	http.HandleFunc("/blog", routes.Blog(tmpl))
	http.HandleFunc("/logout", routes.Logout(tmpl))

	http.ListenAndServe(address, nil)
}

