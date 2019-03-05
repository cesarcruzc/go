package main

import "fmt"

func main() {
	// Forma 1
	var nombres []string

	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Nombre 1")
	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Nombre 2")
	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Nombre 3")
	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Nombre 4")
	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Nombre 5")
	fmt.Printf("Nombres: Tamaño del Slice: %d Capacidad del Slice %d\n", len(nombres), cap(nombres))

	fmt.Println(nombres)

	// Forma 2
	apellidos := make([]string, 0)
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Nombre 1")
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Nombre 2")
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Nombre 3")
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Nombre 4")
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))
	apellidos = append(apellidos, "Nombre 5")
	fmt.Printf("Apellidos: Tamaño del Slice: %d Capacidad del Slice %d\n", len(apellidos), cap(apellidos))

	fmt.Println(apellidos)

	edades := []int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Printf("Edades: Tamaño del Slice: %d Capacidad del Slice %d\n", len(edades), cap(edades))

	fmt.Println(edades)

}
