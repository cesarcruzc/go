package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	// Esperar a que se ejecuten las gorutinas
	var wg sync.WaitGroup

	numbers := []uint32{1234, 456, 12343, 34563456, 76543, 12342221, 1234432, 346343, 56456, 123456, 123454432}

	// Esperar a que se ejecuten las gorutinas, se le pasan la cantidad de gorutinas a ejecutar
	wg.Add(len(numbers))

	fmt.Println("Comienza el proceso")

	for _, v := range numbers {
		go func(a uint32) {
			// Cuando se termine la gorutina, quite uno a la cantidad de gorutinas a ejecutar
			defer wg.Done()
			fmt.Println(a, PrimeNumber(a))
		}(v)
	}

	// Esperar a que se ejecuten las gorutinas delcaradas en el wg.Add
	wg.Wait()

	fmt.Println("Terminó el proceso")
}

func PrimeNumber(a uint32) bool {
	c := 0
	var i uint32

	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}

	if c == 2 {
		return true
	}

	return false
}
