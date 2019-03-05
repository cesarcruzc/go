package main

import (
	"fmt"
)

func main() {

	// Declaración

	// Forma 1
	idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Inglés"

	fmt.Println(idiomas)

	// Forma 2
	languages := map[string]string{
		"es": "Español",
		"en": "Inglés",
		"fr": "Francés",
		"pt": "Portugues",
	}

	fmt.Println(languages)
	delete(languages, "en")
	fmt.Println(languages)

	if language, ok := languages["pt"]; ok {
		fmt.Println("La posición pt si existe y vale", language)
	} else {
		fmt.Println("La posición no existe")
	}

	fmt.Println(languages)

	numeros := map[int]string{
		0:  "Cero",
		-1: "Uno",
		2:  "Dos",
		3:  "Tres",
		4:  "Cuatro",
		5:  "Cinco",
	}

	fmt.Println(numeros)
}
