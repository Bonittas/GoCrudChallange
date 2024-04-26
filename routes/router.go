package routes

import (
    "net/http"

    "github.com/Bonittas/GoCrudChallenge/api"
)

// RegisterRoutes registers all routes with the provided handler.
func RegisterRoutes(handler *api.Handler) {
    http.HandleFunc("/person", handler.GetPersons)
    http.HandleFunc("/person/", handler.GetPerson)
    http.HandleFunc("/person/create", handler.CreatePerson)
    http.HandleFunc("/person/update/", handler.UpdatePerson)
    http.HandleFunc("/person/delete/", handler.DeletePerson)
}
