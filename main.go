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
	messageLine1, messageLine2 := getFullMessage()
	_, _ = fmt.Fprintln(w, messageLine1)
	_, _ = fmt.Fprintln(w, messageLine2)
}

func getWelcomeMessage(firstName, lastName string) string {
	return fmt.Sprintf("Hello %s %s :)", firstName, lastName)

}

func getCourseMessage(courseNumber int) string {
	return fmt.Sprintf("Welcome to GO course #%v!", courseNumber)
}

func getFullMessage() (welcomeMessage, courseMessage string) {
	firstName, lastName := "John", "Doe"
	courseNumber := 1
	welcomeMessage = getWelcomeMessage(firstName, lastName)
	courseMessage = getCourseMessage(courseNumber)

	return
}

// Browser output : Hello John Doe :)
//					Welcome to GO course #1!
