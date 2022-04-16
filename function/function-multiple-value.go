package main

import "fmt"

func getFullName() (string, string, int) {
	return "rizki", "febriansyah", 20
}

func main() {
	firstName, lastName, umur := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(umur)

	firstName1, _, umur2 := getFullName()
	fmt.Println(firstName1)
	//fmt.Println(lastName2) kalo tidak peduli pada lastName kasih underscor pada variabel untuk posisi lastName
	fmt.Println(umur2)

}
