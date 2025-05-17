package main

import (
	"fmt"
	"sync"
)

func a(mg *sync.WaitGroup) {
	defer mg.Done()

	fmt.Println("a is runing.. ")
}
func main() {


	wg := sync.WaitGroup{}

	wg.Add(1)
	go a(&wg)
	fmt.Println("*******************")
	wg.Wait()
	fmt.Println("--------------------")


}
