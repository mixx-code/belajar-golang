package main

import "fmt"

func main() {
	var nama = "rizki"
	var R byte = nama[0]
	var Rstring string = string(R)

	fmt.Println(nama)
	fmt.Println(Rstring)
}
