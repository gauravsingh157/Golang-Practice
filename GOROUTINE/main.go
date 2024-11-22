package main

import (
	"fmt"
	"time"
)

func VISHAL(){

	fmt.Println("GAURAV TUM AAYE NHU ABHI TAK")

}
func GAURAV(){
	fmt.Println("BHAIYA MAI AA GYA HU KAB KA")
}

func main() {

	fmt.Println("LEARN GOROUTINE....")
	VISHAL()
	time.Sleep(1000*time.Millisecond)
	GAURAV()
}