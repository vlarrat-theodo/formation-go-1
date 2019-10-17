package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/formation-go/", teachToTheodoers)
	_ = http.ListenAndServe(":8080", nil)
}

func teachToTheodoers(w http.ResponseWriter, r *http.Request) {
	var welcomeWord string
	var firstName, lastName string = "John", "Doe"
	var courseNumber int = 1

	welcomeWord = "Hello"
	language := "GO"

	_, _ = fmt.Fprintf(w, "%s %s %s :)\n", welcomeWord, firstName, lastName)
	_, _ = fmt.Fprintf(w, "Welcome to %s course #%v!\n", language, courseNumber)
}

// Browser output : Hello John Doe :)
//					Welcome to GO course #1!
