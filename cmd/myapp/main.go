package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	// ✅ Serve static files from ui/static/ at /static/ URL path
	fs := http.FileServer(http.Dir("ui/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// ✅ Serve the HTML template
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
