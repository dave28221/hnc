package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/* sort out 404
if r.ULR.Path =! "/"{
	http.NotFound(w, r)
	return
}
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("ui/html/*.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing templates:", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("Template execution error:", err)
	}

	for _, t := range tmpl.Templates() {
		fmt.Println("Parsed template:", t.Name())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}
