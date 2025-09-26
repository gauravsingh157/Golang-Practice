package main

import "fmt"

func SenderNumber() {
	for i := 0; i <= 10; i++ {

	}
}

func ReciverNumber() {
	for i := 0; i <= 10; i++ {

	}
}
func main() {
	var ch1 chan string = make(chan string)
	var ch2 chan string

	go func(){
		ch1 <- "Hello Vishal"
		ch2 <- "Hello Aman"

	}()
	
	
	fmt.Printf("ch1 = %v , ch2 = %v",<-ch1 , <-ch2)
}