package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"nama":          "rizki",
		"NIM":           "2010129121",
		"jenis kelamin": "laki-laki",
	}
	//kalo mau nambah data
	mahasiswa["umur"] = "20"

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["nama"])
	fmt.Println(mahasiswa["NIM"])
	fmt.Println(mahasiswa["jenis kelamin"])
	fmt.Println(mahasiswa["umur"])

	//menghapus data
	book := make(map[string]string)
	book["title"] = "kisah kisah"
	book["harga"] = "20.000"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
