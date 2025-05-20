package main

import (
	"fmt"

)

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
	var Saurav Student
	//Gaurav.class = 12

	//Gaurav.Name = "Gaurav Singh"
	Gaurav := Student{
		Name:       "Gaurav Singh",
		Age:        20,
		RollNumber: 12345.67,
		SchoolName: "BHU ",
		Email: "gauravsingh@gmail",
		Address: "Uttar Pradesh (Varanasi)",
		MobilNo: 9384383733,
		
	}
	
	fmt.Println("Program start",Saurav)
	fmt.Println(" Student Data ....",Gaurav)
	

}
