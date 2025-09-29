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

// package level variable
var tmpl *template.Template

func templateParse() {
	var err error
	tmpl, err = template.ParseGlob("ui/html/*.html")
	if err != nil {
		log.Println("Error parsing templates:", err)
		// nmaybe make fatal error
		return

	}
	for _, t := range tmpl.Templates() {
		fmt.Println("Parsed template:", t.Name())
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println("Template execution error:", err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	err := tmpl.ExecuteTemplate(w, "login", nil)
	if err != nil {
		log.Println("Template execution error:", err)
	}

}
