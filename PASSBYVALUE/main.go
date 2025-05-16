package main

import "fmt"

func main() {
	a := 4
	b := 5

	fmt.Println("Main func calling =", a+b)
	fmt.Println("Adding value  =",add(a , b))
	fmt.Println("Subtraction value = ",sub(a , b))
	fmt.Println("Multipal  value = ",mul(a , b))


}
func add(e, f int) int {
	c := e + f
	return c
}

func sub(e, f int) int {
	return e - f
}
func mul(e, f int) int {
	return e * f
}
// package main

// import "fmt"


// func increment(num int) {
//     num = num + 1
//     fmt.Println("Inside function:", num) 
// }

// func main() {
//     original := 10
//     increment(original)
//     fmt.Println("Outside function:", original) 
// }
