package main

import "fmt"

/*
interface adalah tipe data abstract, dia tidak memiliki
implementasi langsung
Sebuah interface berisikan definisi - definisi method
Biasanya interface digunakan sebagai kontrak
*/

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type hasName interface {
	GetName() string
}

func SayHello(hN hasName) {
	fmt.Println("Hello", hN.GetName())
}

func SayAnimal(name hasName) {
	fmt.Println("Name Animal", name.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	cat := Animal{
		Name: "Rabbit",
	}

	SayAnimal(cat)
}
