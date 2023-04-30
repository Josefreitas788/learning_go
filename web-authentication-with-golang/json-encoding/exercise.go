package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	FirstName string
}

func main() {

	println("Listening on port 8000...")
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8000", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		FirstName: "Jul",
	}
	p2 := person{
		FirstName: "Lian",
	}
	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)

	if err != nil {
		log.Println("Bad data to encode", err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {
	people := []person{}

	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Bad data to decode", err)
	}

	fmt.Println(people)
}
