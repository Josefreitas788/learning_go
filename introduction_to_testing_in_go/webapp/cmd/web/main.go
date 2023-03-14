package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {

	//set up an app config
	app := application{}

	//get appication routes
	mux := app.routes()

	//print out the messages
	log.Println("Starting server on port 8080")

	//start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
