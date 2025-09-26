package main

import "fmt"

func main() {
	num := 4

	ptr := &num

	fmt.Println("Num of value ", ptr)
	fmt.Println("poiter contains ", num)

}
