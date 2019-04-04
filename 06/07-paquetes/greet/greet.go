package greet

import "fmt"

// Variable para utilizar desde otro paquete
var MeVes string
var noMeVes string

func Greet(nombre string) {
	fmt.Println("Hola ", nombre)
}

func noVisible()  {
	fmt.Println("Funcion solamente accesible desde greet.go")
}