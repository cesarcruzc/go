package main

import (
	"fmt"
	"github.com/cesarcruzc/go/06/07-paquetes/farewell"
	"github.com/cesarcruzc/go/06/07-paquetes/greet"
)

func main() {
	nombre := "Nombre 1"
	greet.Greet(nombre)

	greet.MeVes = "Texto asignado desde el main"
	fmt.Println(greet.MeVes)

	farewell.Farewell(nombre)
}