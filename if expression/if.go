package main

import "fmt"

func main() {
	//kondisi false
	fmt.Println("kondisi false")
	nama := "rizki"

	if nama == "kiki" {
		fmt.Println("ya betul")
	} else {
		fmt.Println("ya salah")
	}
	//kondisi true
	fmt.Println("kondisi true")
	name := "rizki"

	if name == "rizki" {
		fmt.Println("ya betul")
	} else {
		fmt.Println("ya salah")
	}

	//if else if
	names := "rizkii"

	if names == "rizki" {
		fmt.Println("ya betul kamu rizki")
	} else if names == "kiki" {
		fmt.Println("ya betul kamu kiki")
	} else {
		fmt.Println("kamu siapa?")
	}

	//if dengan short statement
	username := "kiki"
	if length := len(username); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("kamu sudah benar")
	}
}
