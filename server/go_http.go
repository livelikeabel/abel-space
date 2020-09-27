package main

import (
	"log"
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
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Person *Person `json:"person"`
}

// Person struct
type Person struct {
	Name string `json:"name"`
}

// Get all wise sayings
func getWiseSayings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

// Get wise saying
func getWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wise saying"))
}

// Create wise saying
func createWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create wise saying"))
}

// Update wise saying
func updateWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update wise saying"))
}

// Delete wise saying
func deleteWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete wise saying"))
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/wise-sayings", getWiseSayings).Methods("GET")
	r.HandleFunc("/api/wise-sayings/{id}", getWiseSaying).Methods("GET")
	r.HandleFunc("/api/wise-sayings", createWiseSaying).Methods("POST")
	r.HandleFunc("/api/wise-sayings/{id}", updateWiseSaying).Methods("PUT")
	r.HandleFunc("/api/wise-sayings/{id}", deleteWiseSaying).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
