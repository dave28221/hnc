package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//router
	router := http.NewServeMux()

	//routes
	router.HandleFunc("GET /", homeHandler)
	router.HandleFunc("GET /login", loginHandler)
	router.HandleFunc("POST /login", loginHandler)

	//parse files
	fs := http.FileServer(http.Dir("ui/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
