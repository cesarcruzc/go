package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 20; i >= 10; i-- {

		if i == 15 {
			break // Continua con la siguiente iteración
		} else if i == 16 {
			continue // Rompe el ciclo for
		}

		fmt.Println(i)
	}

	array := [3][3]int{}
	value := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			value++
			array[i][j] = value
		}
	}

	fmt.Println(array)

	for line := 0; line < 3; line++ {
		for column := 0; column < 3; column++ {
			fmt.Println(array[line][column])
		}
	}

}
