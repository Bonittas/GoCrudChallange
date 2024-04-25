package database

import (
	"errors"
	"sync"

	"github.com/Bonittas/GoCrudChallenge/models"
	"github.com/google/uuid"
)

// InMemoryDB represents an in-memory database.
type InMemoryDB struct {
	mu       sync.RWMutex
	persons  map[string]*model.Person
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		persons: make(map[string]*model.Person),
	}
}

func (db *InMemoryDB) GetAllPersons() ([]model.Person, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	persons := make([]model.Person, 0, len(db.persons))
	for _, person := range db.persons {
		persons = append(persons, *person)
	}
	return persons, nil
}

func (db *InMemoryDB) GetPersonByID(id string) (*model.Person, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	person, ok := db.persons[id]
	if !ok {
		return nil, nil
	}
	return person, nil
}

func (db *InMemoryDB) CreatePerson(person *model.Person) error {
	if person.ID == "" {
		person.ID = uuid.New().String()
	}

	if err := db.validatePerson(person); err != nil {
		return err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	// Check if person ID already exists
	if _, ok := db.persons[person.ID]; ok {
		return errors.New("person ID already exists")
	}

	// Add the person to the database
	db.persons[person.ID] = person
	return nil
}

func (db *InMemoryDB) UpdatePerson(id string, person *model.Person) (*model.Person, error) {
	if err := db.validatePerson(person); err != nil {
		return nil, err
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.persons[id]; !ok {
		return nil, errors.New("person not found")
	}

	person.ID = id
	db.persons[id] = person
	return person, nil
}

func (db *InMemoryDB) DeletePerson(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.persons[id]; !ok {
		return errors.New("person not found")
	}

	delete(db.persons, id)
	return nil
}

func (db *InMemoryDB) validatePerson(person *model.Person) error {
	if person.Name == "" {
		return errors.New("name is required")
	}

	if person.Age <= 0 {
		return errors.New("age must be greater than 0")
	}

	return nil
}
