package main

import "fmt"

func main() {
	/*
	var nombres [4]string

	nombres[0] = "Nombre 1"
	nombres[1] = "Nombre 2"
	nombres[2] = "Nombre 3"
	nombres[3] = "Nombre 4"

	fmt.Println(nombres[3])
	*/

	nombres := [3]string{"Nombre 1", "Nombre 2", "Nombre 3"}
	nombre := nombres[1]
	size := len(nombres)

	fmt.Println(nombre)
	fmt.Println("Tamaño del arreglo", size)
	fmt.Println(nombres[0:3])

}
