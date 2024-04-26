package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Bonittas/GoCrudChallenge/models"
)

// Database interface defines the methods for interacting with the database.
type Database interface {
	GetAllPersons() ([]model.Person, error)
	GetPersonByID(id string) (*model.Person, error)
	CreatePerson(person *model.Person) error
	UpdatePerson(id string, person *model.Person) (*model.Person, error)
	DeletePerson(id string) error
}

// Handler struct represents the API handler.
type Handler struct {
	db Database
}

// NewHandler creates a new instance of the API handler.
func NewHandler(db Database) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetPersons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	persons, err := h.db.GetAllPersons()
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondWithJSON(w, http.StatusOK, persons)
}

func (h *Handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract the ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/person/")
	person, err := h.db.GetPersonByID(id)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if person == nil {
		h.respondWithError(w, http.StatusNotFound, "Person not found")
		return
	}

	h.respondWithJSON(w, http.StatusOK, person)
}

func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.db.CreatePerson(&person); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondWithJSON(w, http.StatusCreated, person)
}
func (h *Handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract the ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/person/update/")
	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedPerson, err := h.db.UpdatePerson(id, &person)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to update person: "+err.Error())
		return
	}

	if updatedPerson == nil {
		h.respondWithError(w, http.StatusNotFound, "Person not found")
		return
	}

	h.respondWithJSON(w, http.StatusOK, updatedPerson)
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract the ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/person/delete/")
	if err := h.db.DeletePerson(id); err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.respondWithJSON(w, http.StatusOK, map[string]string{"message": "Person deleted successfully"})
}
func (h *Handler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
