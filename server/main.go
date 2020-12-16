package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

// WiseSaying struct (Model)
type WiseSaying struct {
	ID         string `json:"id"`
	Text       string `json:"text"`
	PersonName string `json:"person_name"`
}

// Person struct
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Init wiseSayings var as a slice WiseSaying struct
var wiseSayings []WiseSaying

var db *sql.DB
var err error

// Get all wise sayings
func getWiseSayings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var wiseSayings []WiseSaying

	result, err := db.Query("SELECT ws.id, ws.text, p.name FROM wise_saying ws INNER JOIN person p on ws.person_id = p.id")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var wiseSaying WiseSaying
		err := result.Scan(&wiseSaying.ID, &wiseSaying.Text, &wiseSaying.PersonName)
		if err != nil {
			panic(err.Error())
		}
		wiseSayings = append(wiseSayings, wiseSaying)
	}

	json.NewEncoder(w).Encode(wiseSayings)
}

// Get single wise sayings
func getWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT ws.id, ws.text, p.name FROM wise_saying ws INNER JOIN person p on ws.person_id = p.id WHERE ws.id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var wiseSaying WiseSaying

	for result.Next() {
		err := result.Scan(&wiseSaying.ID, &wiseSaying.Text, &wiseSaying.PersonName)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(wiseSaying)
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

// Update wise saying
func updateWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range wiseSayings {
		if item.ID == params["id"] {
			var wiseSaying WiseSaying
			_ = json.NewDecoder(r.Body).Decode(&wiseSaying)
			wiseSaying.ID = item.ID
			copy(wiseSayings[index:], []WiseSaying{wiseSaying})
			json.NewEncoder(w).Encode(wiseSaying)
			return
		}
	}
}

// Delete wise saying
func deleteWiseSaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	for index, item := range wiseSayings {
		if item.ID == params["id"] {
			wiseSayings = append(wiseSayings[:index], wiseSayings[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(wiseSayings)
}

// Main function
func main() {

	dataSource := di.user + ":" + di.pwd + "@tcp(" + di.url + ")/" + di.database
	db, err = sql.Open(di.engine, dataSource)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// // Init router
	router := mux.NewRouter()

	// // Hardcoded data - @todo: add database
	// wiseSayings = append(wiseSayings, WiseSaying{ID: "1", Text: "Text One", Person: &Person{Name: "John"}})
	// wiseSayings = append(wiseSayings, WiseSaying{ID: "2", Text: "Text Two", Person: &Person{Name: "Steve"}})

	// // Route handles & endpoints
	router.HandleFunc("/wise-sayings", getWiseSayings).Methods("GET")
	// router.HandleFunc("/wise-sayings", createWiseSaying).Methods("POST")
	router.HandleFunc("/wise-sayings/{id}", getWiseSaying).Methods("GET")
	// router.HandleFunc("/wise-sayings/{id}", updateWiseSaying).Methods("PUT")
	// router.HandleFunc("/wise-sayings/{id}", deleteWiseSaying).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))

	// TODO: 이거 보고 마저 만들기!
	// https://medium.com/@hugo.bjarred/rest-api-with-golang-mux-mysql-c5915347fa5b
	// https://www.youtube.com/watch?v=SonwZ6MF5BE
}
