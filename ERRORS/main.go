package main

import (
	"fmt"
)

const (
	ErrCodeInvalidInput = 1
	ErrCodeNotFound     = 2
)

func checkAge(age int) (int, string) {
	if age < 0 {

		return ErrCodeInvalidInput, "Age cannot be negative"
	}
	if age < 15 {

		return ErrCodeNotFound, "User is under 18"
	}
	return 0, "Age is valid"
}

func main() {
	code, msg := checkAge(15)

	if code != 0 {
		fmt.Printf("Error [%d]: %s\n", code, msg)
	} else {
		fmt.Println("Success:", msg)
	}
}
