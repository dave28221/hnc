package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var tmpl *template.Template
var err error

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_TOKEN")))

func templateParse() {
	tmpl, err = template.ParseGlob("ui/html/*.html")
	if err != nil {
		log.Println("Error parsing templates:", err)
		// maybe make fatal error
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

	// Check authentication FIRST
	session, err := store.Get(r, "sessionCreation")
	if err != nil {
		session.Values["authenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username, _ := session.Values["username"].(string)

	userData := struct {
		Username        string
		isAuthenticated bool
	}{
		Username:        username,
		isAuthenticated: true,
	}

	err = tmpl.ExecuteTemplate(w, "index", userData)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "login", nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("Username")
		password := r.FormValue("Password")

		users, err := existingUser(db, Users{Username: username})
		if err != nil {
			log.Println("Database error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		user := users[0]
		if !CheckPasswordHash(password, user.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		session, err := store.Get(r, "sessionCreation")
		if err != nil {
			log.Println("Session error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		// Set session options when creating the session
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600 * 8, // 8 hours
			HttpOnly: true,
			Secure:   false, // set to true in prod
			SameSite: http.SameSiteStrictMode,
		}

		session.Values["authenticated"] = true
		session.Values["username"] = username

		// Save session
		err = session.Save(r, w)
		if err != nil {
			log.Println("Session save error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		// Redirect to homepage
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sessionCreation")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1 // Delete the cookie
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("createUser")
		password := r.FormValue("createPassword")

		users, err := existingUser(db, Users{Username: username})
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if len(users) > 0 {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		hashedPassword, err := HashPassword(password)
		if err != nil {
			log.Println("Error hashing password:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		user := Users{
			Username: username,
			Password: hashedPassword,
		}

		_, err = Insert(db, user)
		if err != nil {
			log.Println("Error inserting user:", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}
