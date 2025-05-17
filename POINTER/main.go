package main

import "fmt"

func GAURAV(num int) {
	num = 10
	fmt.Println("IN CHANGENUM ", num)
}

func main() {
	num := 100
	fmt.Println("BEFORE CHANGENUM ", num)
	GAURAV(num)

}
