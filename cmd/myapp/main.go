package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dbSetup()
	templateParse()

	router := http.NewServeMux()

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/create", createHandler)
	router.HandleFunc("/logout", logoutHandler)

	//parse files from static
	fs := http.FileServer(http.Dir("ui/static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
