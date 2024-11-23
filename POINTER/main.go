package main

import "fmt"

func changeNum(num int) {
	num = 10
	fmt.Println("IN CHANGENUM ",num)
}

func main() {
	num := 1
	fmt.Println("BEFORE CHANGENUM ",num)
	changeNum(num)

}