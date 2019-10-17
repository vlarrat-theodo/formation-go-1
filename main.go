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
	_, _ = fmt.Fprint(w, "Hello Theodoers!")
}

// Browser output : Hello Theodoers!
