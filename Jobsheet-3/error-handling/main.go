package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("No data found.")

func findUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("findUser(%d): %w", id, ErrNotFound)
	}
	return nil
}

type ValidationError struct {
	Field string
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func ageValidation (age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "Age cannot be negative."}
	}
	return nil
}

func main() {
	if err := findUser(0); errors.Is(err, ErrNotFound) {
		fmt.Println("Handled as 404: ", err)
	}

	var ve *ValidationError
	if err := ageValidation(-5); errors.As(err, &ve) {
		fmt.Println("Field error: ", ve.Field)
	}
}