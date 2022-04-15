package main

import "fmt"

func main() {
	var name string

	name = "rizki febriansyah"
	fmt.Println(name)

	//variabel secara langsung
	var kota = "jakarta"
	fmt.Println(kota)

	var ukuran = 50
	fmt.Println(ukuran)

	//mengganti kata var dengan :=
	// gak boleh pakai var
	//contoh var nama := "rizki" ini tidak boleh
	nama := "rizki"
	fmt.Println(nama)
	nama = "febriansyah"
	fmt.Println(nama)

	//deklarasi multiple variabel
	var (
		namadepan    = "rizki"
		namabelakang = "febriansyah"
	)
	fmt.Println(namadepan)
	fmt.Println(namabelakang)

}
