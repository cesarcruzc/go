package main

import "fmt"

func main() {

	switch a := 3; {
	case a >= 0 && a <= 5:
		fmt.Println("Estás entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estás en fin de semana")
	default:
		fmt.Println("No es un día valido")

	}
}
