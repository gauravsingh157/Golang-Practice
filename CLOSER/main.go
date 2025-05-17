package main

import "fmt"


func Gaurav() func() int {

	i := 0
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
