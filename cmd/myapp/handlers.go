package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// package level variable
var tmpl *template.Template
var err error

func templateParse() {
	tmpl, err = template.ParseGlob("ui/html/*.html")
	if err != nil {
		log.Println("Error parsing templates:", err)
		// nmaybe make fatal error
		return

	}
	for _, t := range tmpl.Templates() {
		fmt.Println("Parsed template:", t.Name())
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println("Template execution error:", err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("createUser")
		password := r.FormValue("createPassword")

		//fmt.Printf("Received: %s, %s\n", username, password)

		users, error := existingUser(db, Users{Username: username})
		if error != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if len(users) > 0 {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		hashedPassword, error := HashPassword(password)
		if error != nil {
			log.Println("Error hashing password:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		user := Users{
			Username: username,
			Password: hashedPassword,
		}

		_, err := Insert(db, user)
		if err != nil {
			log.Println("Error inserting user:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
