package main

import "fmt"

func main() {
	var n1, n2 uint8
	n1 = 15
	n2 = 5

	r := suma(n1, n2)

	fmt.Println(tipoEdad(r))
}

func suma(a, b uint8) uint8 {
	return a + b
}

func tipoEdad(edad uint8) string {

	var tipo string

	switch {
	case edad < 12:
		tipo = "Niñez"
	case edad < 18:
		tipo = "Adolescencia"
	default:
		tipo = "Madurez"
	}

	return tipo
}
