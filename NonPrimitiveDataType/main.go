package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	RollNumber float32
	SchoolName string
	Email      string
	Address    string
	MobilNo    int
}

func main() {
	var Gaurav Student
	//Gaurav.class = 12

	//Gaurav.Name = "Gaurav Singh"
	Shiva := Student{
		Name:       "Gaurav Singh",
		Age:        20,
		RollNumber: 12345.67,
		
	}
	fmt.Println(" Student Data ....", Shiva)
	fmt.Println(" Student Data ....", Gaurav)

}
