package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/go-playground/form"
)

type User struct {
	Username string
	Password string
}

// single instance of Decoder - cache struct
var decoder *form.Decoder

func formData() {
	decoder = form.NewDecoder()

	// this simulates the results of http.Request's ParseForm() function
	values := parseForm()

	var user User

	err := decoder.Decode(&user, values)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", user)
}

func parseForm() url.Values {
	return url.Values{
		"Username": []string{},
		"Password": []string{},
	}
}
