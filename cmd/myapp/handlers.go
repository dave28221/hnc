package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var tmpl *template.Template
var err error
var store = sessions.NewCookieStore()

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't get .env file")
	}

	sessionToken := os.Getenv("SESSION_TOKEN")
	if sessionToken == "" {
		log.Fatal("SESSION_TOKEN is not set")
	}

	store = sessions.NewCookieStore([]byte(sessionToken))

}

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

	session, err := store.Get(r, "sessionCreation")
	if err != nil {

		err = tmpl.ExecuteTemplate(w, "index", nil)
		if err != nil {
			log.Println("Template execution error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	auth, ok := session.Values["authenticated"].(bool)
	if ok && auth {
		username, _ := session.Values["username"].(string)
		userData := struct {
			Username        string
			IsAuthenticated bool
		}{
			Username:        username,
			IsAuthenticated: true,
		}
		err = tmpl.ExecuteTemplate(w, "index", userData)
	} else {
		err = tmpl.ExecuteTemplate(w, "index", nil)
	}

	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sessionCreation")
	auth, ok := session.Values["authenticated"].(bool)
	if ok && auth {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

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

		log.Printf("DEBUG: Login attempt for username='%s', password length=%d", username, len(password))

		users, err := existingUser(db, Users{Username: username})
		if err != nil {
			log.Println("Database error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		log.Printf("DEBUG: Found %d users", len(users))

		if len(users) == 0 {
			log.Printf("DEBUG: No user found with username '%s'", username)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		user := users[0]
		log.Printf("DEBUG: User from DB - username='%s', password hash length=%d", user.Username, len(user.Password))

		passwordMatch := CheckPasswordHash(password, user.Password)
		log.Printf("DEBUG: Password match result: %v", passwordMatch)

		if !passwordMatch {
			log.Printf("DEBUG: Password check FAILED for user '%s'", username)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		log.Printf("DEBUG: Login successful for '%s'", username)

		session, err := store.Get(r, "sessionCreation")
		if err != nil {
			log.Println("Session error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600 * 8, // 8 hours
			HttpOnly: true,
			Secure:   false, // set to true in prod
			SameSite: http.SameSiteStrictMode,
		}

		session.Values["authenticated"] = true
		session.Values["username"] = username

		err = session.Save(r, w)
		if err != nil {
			log.Println("Session save error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

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
