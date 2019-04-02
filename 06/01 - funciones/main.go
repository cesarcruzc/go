package main

import "fmt"

func main() {
	fmt.Println("Hello 0")
	greet("Tony", 100)
}

func greet(name string, edad uint8) {
	fmt.Printf("Hola %s tienes %d años de edad \n", name, edad)

	if edad > 30 {
		return
	}

	fmt.Printf("Eres joven")
}
