package main

import "fmt"

func halloWorld() {
	fmt.Println("Hallo World")
}

func main() {
	halloWorld() //cara memanggil function

	//bisa di kombinasikan dengan looping juga
	for i := 0; i < 10; i++ {
		halloWorld()
	}
}
