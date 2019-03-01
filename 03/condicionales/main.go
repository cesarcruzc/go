package main

import "fmt"

func main() {

	isValid := true

	if isValid {
		fmt.Println("Verdadero")
		fmt.Printf("%T\n", isValid)
		fmt.Println(isValid)

		if 10 < 10 {
			fmt.Println("Si :D")
		} else {
			fmt.Println("No :P")
		}

	} else {
		fmt.Println("False")
	}

	fmt.Println(isValid)
	fmt.Println("Fin del programa")

}
