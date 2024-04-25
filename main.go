package main

import (
	"log"
	"net/http"

	"github.com/Bonittas/GoCrudChallenge/api"
	"github.com/Bonittas/GoCrudChallenge/database"
)

func main() {
	// Create a new instance of the database
	db := database.NewInMemoryDB()

	// Create a new instance of the API handler with the database
	handler := api.NewHandler(db)

	// Set up routes
	http.HandleFunc("/person", handler.GetPersons)
	http.HandleFunc("/person/", handler.GetPerson)
	http.HandleFunc("/person/create", handler.CreatePerson)
	http.HandleFunc("/person/update", handler.UpdatePerson)
	http.HandleFunc("/person/delete", handler.DeletePerson)

	// Apply CORS middleware to the router
	routerWithCORS := api.CORS(http.DefaultServeMux)

	// Start the server with the CORS-enabled router
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", routerWithCORS))
}
