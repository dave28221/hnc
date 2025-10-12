package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

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
