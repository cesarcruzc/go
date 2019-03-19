package main

import (
	"fmt"
)

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"}

	for i, v := range numeros {
		fmt.Println(i, v)
	}

	// No usar i, se reemplaza por _ (Ignorar el valor i)
	for _, v := range numeros {
		fmt.Println(v)
	}

	// El valor en un forrnge puede omitirse, los indices no
	for i := range numeros {
		fmt.Println(i)
	}

	nombres := map[string]string{"es": "España", "co": "Colombia", "br": "Brasil"}

	for i, v := range nombres {
		fmt.Println(i, v)
	}

	frase := "Hola mundo"
	// Imprime en v el valor en binario
	// string(v) imprime el valor en string del valor
	for i, v := range frase {
		fmt.Println(i, v, string(v))
	}

	// delcarar un slice en el for
	for _, entero := range []int{15, 36, 24, 85} {
		fmt.Println(entero)
	}

	// delcarar string en el for
	for i, v := range "Hello" {
		fmt.Println(i, v, string(v))
	}

}
