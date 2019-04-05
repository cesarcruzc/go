package main

import "fmt"

type Person struct {
	name string
	age  int8
}

type Number int

func (p Person) greet() {
	fmt.Println("Hello I'm a Person")
}

func (n Number) introduceYourSelf() {
	fmt.Println("Hello I'm a method for Number type")
}

// Recibe un puntero
func (p *Person) assignAge(i int8) {
	if i >= 0 {
		p.age = i
	} else {
		fmt.Println("Is not valid a negative age")
	}
}

func main() {
	var person Person
	person.greet()

	var number Number
	number.introduceYourSelf()

	person.assignAge(-2)
	fmt.Println(person)
}
