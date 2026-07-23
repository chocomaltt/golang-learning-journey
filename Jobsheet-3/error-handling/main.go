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
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func ageValidation(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "Age cannot be negative."}
	} else if age > 150 {
		return &ValidationError{Field: "age", Msg: "Age cannot be over 150 (It's not possible, is it?)."} //Challenge No.3
	}
	return nil
}

func main() {
	if err := findUser(0); errors.Is(err, ErrNotFound) {
		fmt.Println("Handled as 404: ", err)
	}

	var ve *ValidationError
	fmt.Println("\nValidating Age: -5")
	if err := ageValidation(-5); errors.As(err, &ve) {
		fmt.Println("Field error: ", ve.Field)
	}

	// Challenge No.3
	fmt.Println("\nValidating Age: 151")
	if err := ageValidation(151); errors.As(err, &ve) {
		fmt.Printf("Field error: %s\nMessage: %s", ve.Field, ve.Msg)
	}
}
