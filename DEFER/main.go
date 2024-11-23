package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {
		 defer fmt.Println(i)
	}
	defer fmt.Println("Program start")
	fmt.Println("Program ended")

    fmt.Println("Program closed")

}
