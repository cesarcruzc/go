package main

import (
	"fmt"
)

func main() {
	var n1, n2 uint8
	n1 = 15
	n2 = 5

	r := suma(n1, n2)

	fmt.Println(tipoEdad(r))

	n := [] int8{52, 63, 47, 5, 5, 3, 7, 6, 100, 2, -5}
	maximo, minimo := maxymin(n)

	fmt.Println("Máximo:", maximo)
	fmt.Println("Mínimo:", minimo)
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

func maxymin(numeros [] int8) (max int8, min int8) {

	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return
}
