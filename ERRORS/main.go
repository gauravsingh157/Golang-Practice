// package main

// import "fmt"

// func main() {
// 	b := []int{
// 		1, 2, 3, 4,
// 	}
// 	fmt.Println(b)
// 	fmt.Println(len(b))

// }
package main

import (
	"errors"
	"fmt"
)


const (
	ErrCodeNotFound     = 1001
	ErrCodeUnauthorized = 1002
	ErrCodeInvalidInput = 1003
	ErrCodeInternal     = 1004
)


type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}


func NewCustomError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}


func getUser(id int) error {
	if id == 0 {
		return NewCustomError(ErrCodeInvalidInput, "User ID must not be zero")
	}
	if id == 42 {
		return NewCustomError(ErrCodeNotFound, "User not found")
	}
	return nil // success
}

func main() {
	err := getUser(0)
	if err != nil {
		var customErr *CustomError
		if errors.As(err, &customErr) {
			fmt.Println("Handled Custom Error:", customErr)
		} else {
			fmt.Println("Generic error:", err)
		}
	}
}
