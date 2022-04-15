package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke- ", counter)
		counter++
	}
	// for dengan statement

	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan ke- ", i)
	}

	//perulangan dengan array
	slice := []string{"rizki", "kiki", "iki"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	for i, value := range slice {
		fmt.Println("index ke- ", i, "= ", value)
	}
	//tanpa index, untuk mengabaikan variabel kasih underscore " _ "
	for _, value := range slice {
		// fmt.Println("index ke- ", i, "= ", value)
		fmt.Println(value)
	}
	//cara menggunakan di MAP
	mahasiswa := make(map[string]string)
	mahasiswa["nama"] = "rizki"
	mahasiswa["umur"] = "20"
	for key, value := range mahasiswa {
		// fmt.Println("index ke- ", i, "= ", value)
		fmt.Println(key, " = ", value)
	}

}
