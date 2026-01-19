// auth.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store *sessions.CookieStore

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sessionSecret := os.Getenv("SESSION_TOKEN")
	if sessionSecret == "" {
		log.Fatal("SESSION_TOKEN not set in environment")
	}

	// start session
	store = sessions.NewCookieStore([]byte(sessionSecret))

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24,
		HttpOnly: true,
		Secure:   false, // change in prod
		SameSite: http.SameSiteLaxMode,
	}
}

func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	session.Values["authenticated"] = false
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		log.Println("Error destroying session:", err)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
