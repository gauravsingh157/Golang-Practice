package main

import "fmt"

func main() {

	fmt.Println("Variadic Sum value ",add(25 ,25 ,32))
}

func add( G ... int) int {
	Sum := 0
	for _  , val := range G{
		//Sum := Sum
		Sum = Sum + val


	}
	return Sum
}