package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/formation-go/", teachToTheodoers)
	http.HandleFunc("/formation-go/public/", teachToPublic)
	_ = http.ListenAndServe(":8080", nil)
}

type User interface {
	GetWelcomeMessage() (string, error)
}

type DefaultUser struct{}

func (d DefaultUser) GetWelcomeMessage() (string, error) {
	return "Hello random user :)", nil
}

type Theodoer struct {
	FirstName string
	LastName  string
}

func (t Theodoer) GetWelcomeMessage() (string, error) {
	var (
		welcomeMessage string
		err            error
	)
	if t.FirstName == "" || t.LastName == "" {
		err = errors.New("Incomplete information for Theodoer")
	} else {
		welcomeMessage = fmt.Sprintf("Hello %s %s from Theodo :)", t.FirstName, t.LastName)
	}
	return welcomeMessage, err

}

func teachTo(w http.ResponseWriter, r *http.Request, u User) {
	messageLine1, messageLine2 := getFullMessage(u)
	_, _ = fmt.Fprintln(w, messageLine1)
	_, _ = fmt.Fprintln(w, messageLine2)
}

func teachToTheodoers(w http.ResponseWriter, r *http.Request) {
	theodoer := Theodoer{
		FirstName: "John",
		LastName:  "Doe",
	}
	teachTo(w, r, theodoer)
}

func teachToPublic(w http.ResponseWriter, r *http.Request) {
	teachTo(w, r, DefaultUser{})
}

func getCourseMessage(courseNumber int) string {
	return fmt.Sprintf("Welcome to GO course #%v!", courseNumber)
}

func getFullMessage(u User) (welcomeMessage, courseMessage string) {
	courseNumber := 1
	tempWelcomeMessage, err := u.GetWelcomeMessage()
	if err != nil {
		welcomeMessage = fmt.Sprintf("An error occured: %s", err)
	} else {
		welcomeMessage = tempWelcomeMessage
	}
	courseMessage = getCourseMessage(courseNumber)

	return
}

// Browser output : Hello John Doe from Theodo :)
//					Welcome to GO course #1!
//
// Browser output /public : Hello random user :)
//							Welcome to GO course #1!
