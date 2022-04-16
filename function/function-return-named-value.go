package main

import "fmt"

func getMahasiswa() (nama string, NIM int, umur int) {
	nama = "rizki"
	NIM = 201012921
	umur = 20
	return
}

func main() {
	nama, NIM, umur := getMahasiswa()
	fmt.Println(nama)
	fmt.Println(NIM)
	fmt.Println(umur)

	a, b, c := getMahasiswa()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
