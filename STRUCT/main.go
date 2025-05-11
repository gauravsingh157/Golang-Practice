package main

import "fmt"

type Empoloyee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
	Email    string
	Address  string
}

func main() {
	Gaurav := Empoloyee{
		Name:     "Gaurav",
		Age:      25,
		Position: "Software Engineer",
		Salary:   50000.0,
		Email:    "gaurav@gmail.com",
		Address:  "Delhi",
	}
	fmt.Println(Gaurav)
}