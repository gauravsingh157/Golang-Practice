package main

import "fmt"


func Gaurav() func() int {

	i := 10000
	return func ()int  {
		i ++
		return i
	}
}

func main() {

	result := Gaurav()

	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	fmt.Println("Result :", result())
	
	


}
