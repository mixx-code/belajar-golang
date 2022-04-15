package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "rizki"
	nama[1] = "kiki"
	nama[2] = "iki"

	fmt.Println("nama index ke 0 = ", nama[0])
	fmt.Println("nama index ke 1 = ", nama[1])
	fmt.Println("nama index ke 2 = ", nama[2])

	//penulisan ke 2
	var nilai = [3]int{10, 20, 30}
	fmt.Println("isi array nilai = ", nilai)

	//melihat panjang nilai array
	fmt.Println("panjang data array nama = ", len(nama))

	fmt.Println("panjang nama pada index ke 0 = ", len(nama[0]))

	//mencoba perbandingan
	fmt.Println("perbandingan nama index ke 0 > nama index ke 1 = ", nama[0] > nama[1]) // bisa
}
