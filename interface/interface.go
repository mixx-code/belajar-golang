package main

import "fmt"

type HashName interface {
	GetName() string
}

func SayHello(hashname HashName) {
	fmt.Println("Hello", hashname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name

}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name

}

func main() {
	var rizki Person
	rizki.Name = "rizki"

	SayHello(rizki)

	cat := Animal{
		Name: "pusi",
	}
	SayHello(cat)
}
