package main

import "fmt"

func main() {
	originCellar := []string{"php", "javascript", "html", "css", "java"}
	destinationCellar := []string{}

	// Canales
	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	// Gorutinas

	// Carga la fotocopiadora
	go func() {
		for _, libro := range originCellar {
			fotocopiadora <- libro
		}
	}()

	// Se deja en la bodega destino
	go func() {
		for {
			// Entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora
			fotocopiado <- libro

			// Agregar el libro a la bodega destino
			destinationCellar = append(destinationCellar, libro)

			if len(originCellar) == len(destinationCellar) {
				// Cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("** Listado de libros fotocopiados **")
	
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
}
