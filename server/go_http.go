package main

import (
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

// Route Struct
type Route struct {
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

// WiseSaying struct
type WiseSaying struct {
	Saying string `json:"saying"`
	Name   string `json:"name"`
}

// HomeHandler write gorilla
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8000", r)
}
