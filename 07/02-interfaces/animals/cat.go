package animals

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Comunicate() {
	fmt.Println("Miaw")
}
