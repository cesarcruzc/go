package main

import "fmt"

func main() {
	var a int
	var b int8

	a = 10
	b = 10

	c := a + int(b)

	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("c es de tipo: %T\n", c)

	fmt.Printf("b es de tipo: %T\n", b)

	cadena := "á"
	fmt.Printf("El valor de la variable cadena es: %s\n", cadena)

	número := 2
	fmt.Printf("El valor de la variable numero es: %d\n", número)
}
