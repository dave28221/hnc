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

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	app := &application{
		templateCache: make(map[string]*template.Template),
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /", app.home)

	log.Fatal(http.ListenAndServe(":8080", router))

}
