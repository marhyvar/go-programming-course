package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Dog struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Breed string `json:"breed"`
    Age int `json:"age"`
}

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello world\n")
}

// returns http response as JSON
func dogJsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    dog1 := Dog{
        Id: 1, 
        Name: "Einstein",
        Breed: "Welsh Corgi Pembroke",
        Age: 11 }

    dog2 := Dog{
        Id: 2,
        Name: "Iita",
        Breed: "Beagle",
        Age: 14 }

    dogs := []Dog{dog1, dog2}

    json.NewEncoder(w).Encode(dogs)
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/dogs", dogJsonHandler)

    http.ListenAndServe(":8090", nil)
}
