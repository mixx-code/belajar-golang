package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hallo", firstName, lastName)
}

func main() {
	sayHelloTo("rizki", "febriansyah")

	firstName := "kiki"
	sayHelloTo(firstName, "febriansyah") //dapat menggunakan variabel juga
}
