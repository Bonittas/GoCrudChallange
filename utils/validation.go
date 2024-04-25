package utils

import (
	"errors"
	"regexp"

	"github.com/Bonittas/GoCrudChallenge/models"
)

// ValidatePerson validates the fields of a Person object.
func ValidatePerson(person *model.Person) error {
	if person.Name == "" {
		return errors.New("name is required")
	}

	if person.Age <= 0 {
		return errors.New("age must be greater than 0")
	}

	// Optionally, validate hobbies if needed
	for _, hobby := range person.Hobbies {
		if !isValidHobby(hobby) {
			return errors.New("invalid hobby")
		}
	}

	return nil
}

// isValidHobby validates a hobby string.
func isValidHobby(hobby string) bool {
	hobbyRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	return hobbyRegex.MatchString(hobby)
}
