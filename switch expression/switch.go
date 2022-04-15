package main

import "fmt"

func main() {
	nama := "rizkirizki"

	switch nama {
	case "rizki":
		fmt.Println("nama : rizki")
		fmt.Println("umur : 20")
		fmt.Println("JK   : laki-laki")
	case "kiki":
		fmt.Println("nama : kiki")
		fmt.Println("umur : 19")
		fmt.Println("JK   : perempuan")
	default:
		fmt.Println("masukan nama yang benar")
	}

	//switch dengan short statement
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama nama sudah benar")
	}

	//switch tanpa kondisi
	panjang := len(nama)
	switch {
	case panjang > 10:
		fmt.Println("nama terlalu panjang")
	case panjang > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
