package main

import "fmt"

type Customer struct {
	nama, alamat string
	usia         int
}

func main() {
	var rizki Customer
	rizki.nama = "Rizki"
	rizki.alamat = "indonesia"
	rizki.usia = 20

	fmt.Println(rizki)
	fmt.Println(rizki.nama)
	fmt.Println(rizki.alamat)
	fmt.Println(rizki.usia)

	//stract literal
	kiki := Customer{
		nama:   "kiki",
		alamat: "indonesia",
		usia:   20,
	}
	fmt.Println(kiki)
}
