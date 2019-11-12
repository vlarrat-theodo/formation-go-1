package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/formation-go/", teachToTheodoers)
	_ = http.ListenAndServe(":8080", nil)
}

type Theodoer struct {
	FirstName string
	LastName  string
}

func (t Theodoer) GetWelcomeMessage() string {
	return fmt.Sprintf("Hello %s %s from Theodo :)", t.FirstName, t.LastName)

}

func teachToTheodoers(w http.ResponseWriter, r *http.Request) {
	messageLine1, messageLine2 := getFullMessage()
	_, _ = fmt.Fprintln(w, messageLine1)
	_, _ = fmt.Fprintln(w, messageLine2)
}

func getCourseMessage(courseNumber int) string {
	return fmt.Sprintf("Welcome to GO course #%v!", courseNumber)
}

func getFullMessage() (welcomeMessage, courseMessage string) {
	theodoer := Theodoer{
		FirstName: "John",
		LastName:  "Doe",
	}
	courseNumber := 1
	welcomeMessage = theodoer.GetWelcomeMessage()
	courseMessage = getCourseMessage(courseNumber)

	return
}

// Browser output : Hello John Doe from Theodo :)
//					Welcome to GO course #1!
