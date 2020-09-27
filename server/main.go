package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// WiseSaying struct (Model)
type WiseSaying struct {
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Person *Person `json:"person"`
}

// Person struct
type Person struct {
	Name string `json:"name"`
}

// Init wiseSayings var as a slice WiseSaying struct
var wiseSayings []WiseSaying

// Get all wise sayings
func getWiseSayings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wiseSayings)
}

// Get single wise sayings
func getWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through wise sayings and find one with the id from the params
	for _, item := range wiseSayings {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&WiseSaying{})
}

// Add new wise saying
func createWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wiseSaying WiseSaying
	_ = json.NewDecoder(r.Body).Decode(&wiseSaying)
	wiseSaying.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	wiseSayings = append(wiseSayings, wiseSaying)
	json.NewEncoder(w).Encode(wiseSaying)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	wiseSayings = append(wiseSayings, WiseSaying{ID: "1", Text: "Text One", Person: &Person{Name: "John"}})
	wiseSayings = append(wiseSayings, WiseSaying{ID: "2", Text: "Text Two", Person: &Person{Name: "Steve"}})

	// Route handles & endpoints
	r.HandleFunc("/wise-sayings", getWiseSayings).Methods("GET")
	r.HandleFunc("/wise-sayings/{id}", getWiseSaying).Methods("GET")
	r.HandleFunc("/wise-sayings", createWiseSaying).Methods("POST")
	// r.HandleFunc("/wise-sayings/{id}", updateWiseSaying).Methods("PUT")
	// r.HandleFunc("/wise-sayings/{id}", deleteWiseSaying).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
