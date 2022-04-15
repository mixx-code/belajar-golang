package main

import "fmt"

func main() {
	//membandingkan nilai string
	nama1 := "rizki"
	nama2 := "rizki febriansyah"

	var hasil1 bool = nama1 == nama2

	fmt.Println("hasil 1", hasil1)
	// membandingkan panjang string
	nama3 := len("rizki")
	nama4 := len("rizki febriansyah")

	fmt.Println("panjang nama3 > nama 4  =", nama3 > nama4)
	fmt.Println("panjang nama3 < nama 4  =", nama3 < nama4)
	fmt.Println("panjang nama3 == nama 4  =", nama3 == nama4)
	fmt.Println("panjang nama3 != nama 4  =", nama3 != nama4)

	//membandingkan bilangan
	nilai1 := 80
	nilai2 := 85

	fmt.Println("panjang nilai 1 > nilai 2  =", nilai1 > nilai2)
	fmt.Println("panjang nilai 1 < nilai 2  =", nilai1 < nilai2)
	fmt.Println("panjang nilai 1 == nilai 2  =", nilai1 == nilai2)
	fmt.Println("panjang nilai 1 != nilai 2  =", nilai1 != nilai2)

}
