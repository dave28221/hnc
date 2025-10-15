package main

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite" // SQLite driver
)

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// / create table if not exists - set this up
func dbSetup() {
	var err error
	db, err = sql.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func Insert(db *sql.DB, user Users) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// search for existing users - setup in handlers.go
func existingUser(db *sql.DB, user Users) ([]Users, error) {
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// slice for the usernames
	var matchingUsers []Users

	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		matchingUsers = append(matchingUsers, user)
	}
	if err = rows.Err(); err != nil {
		return matchingUsers, err
	}
	return matchingUsers, nil

}
