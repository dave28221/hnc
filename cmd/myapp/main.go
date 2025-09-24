package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handler)

	log.Fatal(http.ListenAndServe(":8080", router))

}
