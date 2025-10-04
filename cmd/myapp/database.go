package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func dbSetup() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// https://www.sqlitetutorial.net/sqlite-go/insert/.   ---- link
	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'users' created successfully")

	// Create user instance
	user := Users{
		Username: "James",
		Password: "lol1234",
	}

	// Insert user
	userID, err := Insert(db, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print inserted user info
	fmt.Printf("The user %s was inserted with ID: %d\n", user.Username, userID)
}

// Insert inserts a user into the database and returns the inserted ID
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
