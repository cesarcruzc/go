package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
		}
	}()

	fmt.Println("Llamando g()")
	g(5)
	fmt.Println("g() retorna normalmente")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Aaaaa")
		panic("El número no puede ser mayor a 3")
	}

	defer fmt.Println("Defer en g()")
	fmt.Println("Imprimiendo en g()", i)
}
