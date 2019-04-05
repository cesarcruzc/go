package animals

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Comunicate() {
	fmt.Println("Wof")
}
