package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	templateParse()
	//router
	router := http.NewServeMux()

	//routes
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login", loginHandler)

	//parse files
	fs := http.FileServer(http.Dir("ui/static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
