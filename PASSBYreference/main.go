package main

import "fmt"

func main() {
	f := 5
	g := 10

	addByOne(&f, g)
	fmt.Println("f =", f, "g =", g)
}


func addByOne(a *int, b int) {
	*a = *a + 1 // f will be changed (pass by reference)
	b = b + 1   // g will NOT be changed (pass by value)
}
