package main

import "fmt"

//Persona es una estructura
type Persona struct {
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	// Forma 1
	var persona1 Persona
	persona1.Nombre = "Persona1"
	persona1.Edad = 18
	persona1.Emails = []string{"p1@1", "p1@2"}

	fmt.Println(persona1.Nombre)
	fmt.Println(persona1.Edad)
	fmt.Println(persona1.Emails)

	// Forma 2
	persona2 := Persona{}
	persona2.Nombre = "Persona2"
	persona2.Edad = 20
	persona1.Emails = []string{"p2@1", "p2@2"}

	fmt.Println(persona2.Nombre)
	fmt.Println(persona2.Edad)
	fmt.Println(persona2.Emails)

	// Forma 3
	// Se pueden colocar solamente los valores sin el nombre de la propiedad
	persona3 := Persona{
		Nombre: "Persona3",
		Edad:   22,
		Emails: []string{"p3@1", "p3@2"},
	}

	fmt.Println(persona3.Nombre)
	fmt.Println(persona3.Edad)
	fmt.Println(persona3.Emails)
}
