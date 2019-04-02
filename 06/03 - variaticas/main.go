package main

import "fmt"

func main() {
	severalGreet(18, "Nombre 1", "Nombre 2", "Nombre 3")
}

func severalGreet(edad uint8, nombres ...string) {
	for _, v := range nombres {
		fmt.Println("Hola", v)
		fmt.Println("Edad", edad)
	}
}
