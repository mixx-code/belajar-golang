package main

import "fmt"

type Customer struct {
	nama, alamat string
	usia         int
}

func (customer Customer) sayHi(nama string) {
	fmt.Println("hello", nama, "nama saya", customer.nama)
}

func (a Customer) sayHuu() {
	fmt.Println("from huu", a.nama)
}

func main() {
	var rizki Customer
	rizki.nama = "Rizki"
	rizki.alamat = "indonesia"
	rizki.usia = 20

	rizki.sayHi("kiki")
	rizki.sayHuu()

	// fmt.Println(rizki)
	// fmt.Println(rizki.nama)
	// fmt.Println(rizki.alamat)
	// fmt.Println(rizki.usia)

	// //stract literal
	// kiki := Customer{
	// 	nama:   "kiki",
	// 	alamat: "indonesia",
	// 	usia:   20,
	// }
	// fmt.Println(kiki)
}
