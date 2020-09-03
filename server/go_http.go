package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println(r.URL.Path)
    fmt.Println(r.Form["test_param"])

    for k, v := range r.Form {
	fmt.Println(k)
	fmt.Println(strings.Join(v, ""))
    }

    fmt.Fprintf(w, "Golang WebServer Working!")
}

func main() {
    http.HandleFunc("/", defaultHandler)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
	log.Fatal("ListenAndServe: ", err)
    } else {
	fmt.Println("ListenAndServe Started! -> Port: 9090")
    }
}
