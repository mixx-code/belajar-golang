package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
	message := recover() // menangkap error dan apliksi tetap berjalan
	if message != nil {
		fmt.Println("error dengan message", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("applikasi error")
	}
	fmt.Println("aplikasi berjaln")
}
func main() {
	runApp(true)
	fmt.Println("rizki")
}
