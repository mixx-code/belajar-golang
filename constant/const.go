package main

import "fmt"

func main() {
	const firstname string = "rizki"
	const lastname = "febriansyah"
	const value = 100

	fmt.Println(firstname)
	fmt.Println(lastname)

	//multiple const
	const (
		namadepan    string = "rizki"
		namabelakang        = "febriansyah"
		nilai               = 100
	)
	fmt.Println(namadepan)
	fmt.Println(namabelakang)
}
