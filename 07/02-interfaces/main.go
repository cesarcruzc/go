package main

import "github.com/cesarcruzc/07/02-interfaces/animals"

func main() {
	var d animals.Dog
	var c animals.Cat

	d.Name = "Goofy"
	c.Name = "Rainbow"

	/*
		AdoptDog(d)
		AdoptCat(c)
	*/

	AdoptPet(d)
	AdoptPet(c)

}

/*
func AdoptDog(d animals.Dog) {
	d.Comunicate()
}

func AdoptCat(c animals.Cat) {
	c.Comunicate()
}
*/

func AdoptPet(p animals.Pet) {
	p.Comunicate()
}
