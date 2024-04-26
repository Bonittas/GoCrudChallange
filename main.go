package main

import (
    "log"
    "net/http"

    "github.com/Bonittas/GoCrudChallenge/api"
    "github.com/Bonittas/GoCrudChallenge/database"
    "github.com/Bonittas/GoCrudChallenge/routes"
)

func main() {
    // Create a new instance of the database
    db := database.NewInMemoryDB()

    // Create a new instance of the API handler with the database
    handler := api.NewHandler(db)

    // Register routes
    routes.RegisterRoutes(handler)

    // Apply CORS middleware to the router
    routerWithCORS := api.CORS(http.DefaultServeMux)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", routerWithCORS))
}
