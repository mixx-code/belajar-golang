package main

import "fmt"

func endApp() {
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("applikasi error")
	}
	fmt.Println("aplikasi berjalan")
}
func main() {
	runApp(false)
}
