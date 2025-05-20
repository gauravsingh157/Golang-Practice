package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	AccountBalance := 0
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go Paisabadhao(&AccountBalance, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println("Final Balance:", AccountBalance)
}

func Paisabadhao(balance *int, wg *sync.WaitGroup, tala *sync.Mutex) {
	defer wg.Done()
	tala.Lock()
	(*balance)++
	tala.Unlock()
}
