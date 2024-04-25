package main

import (
	"log"
	"net/http"

	"github.com/Bonittas/GoCrudChallenge/api"
)

func main() {
	handler := api.NewHandler()

	// Set up routes
	http.HandleFunc("/person", handler.GetPersons)
	http.HandleFunc("/person/", handler.GetPerson)
	http.HandleFunc("/person", handler.CreatePerson)
	http.HandleFunc("/person/", handler.UpdatePerson)
	http.HandleFunc("/person/", handler.DeletePerson)

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
