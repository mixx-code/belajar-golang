package main

import "fmt"

func main() {
	name := "rizki"
	counter := 0

	increment := func() {
		name := "kiki"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}
	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
