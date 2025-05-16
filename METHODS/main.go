// package main

// type BHOBBY struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func (user BHOBBY)addUser(a int) {

// }

// func main() {

// 	var Gaurav BHOBBY
// 	Gaurav.addUser(3)

// }
package main

import "fmt"


type Person struct {
    name string
    age  int
}


func (p Person) greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}


func (p *Person) haveBirthday() {
    p.age++
}

func main() {
    p1 := Person{name: "Amit", age: 25}

    
    p1.greet() 
    
    p1.haveBirthday()

    
    p1.greet() 
}



