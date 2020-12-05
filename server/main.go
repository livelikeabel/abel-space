package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Person *Person `json:"person"`
}

// Person struct
type Person struct {
	Name string `json:"name"`
}

// User struct
type User struct {
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
	params := mux.Vars(r) // Get params
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

func dbQuery(db dbInfo, query string) (count int) {
	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	// conn.SetConnMaxLifetime(time.Minute)
	// conn.SetMaxOpenConns(10)
	// conn.SetMaxIdleConns(10)

	defer conn.Close()

	results, err := conn.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	return 0
}

// Main function
func main() {
	// Declear Database
	var db = dbInfo{"root", "", "127.0.0.1:3306", "mysql", "abelspace"}
	var query = "SELECT name FROM users"
	result := dbQuery(db, query)
	print(result)

	// // Init router
	// r := mux.NewRouter()

	// // Hardcoded data - @todo: add database
	// wiseSayings = append(wiseSayings, WiseSaying{ID: "1", Text: "Text One", Person: &Person{Name: "John"}})
	// wiseSayings = append(wiseSayings, WiseSaying{ID: "2", Text: "Text Two", Person: &Person{Name: "Steve"}})

	// // Route handles & endpoints
	// r.HandleFunc("/wise-sayings", getWiseSayings).Methods("GET")
	// r.HandleFunc("/wise-sayings/{id}", getWiseSaying).Methods("GET")
	// r.HandleFunc("/wise-sayings", createWiseSaying).Methods("POST")
	// r.HandleFunc("/wise-sayings/{id}", updateWiseSaying).Methods("PUT")
	// r.HandleFunc("/wise-sayings/{id}", deleteWiseSaying).Methods("DELETE")

	// // Start server
	// log.Fatal(http.ListenAndServe(":8000", r))
}
