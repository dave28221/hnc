package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//parse static files including css and js//

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("")
		tmpl.ParseGlob("html/*.html")
		for _, t := range tmpl.Templates() {
			fmt.Println(t.Name())
		}

	})

	fmt.Println("server is up and running on 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
