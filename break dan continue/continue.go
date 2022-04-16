package main

import "fmt"

//dimana pada saat kodisi di temukan maka akan di longkap dan dilanjutkan ke data selanjutnya
func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke-", i)
	}
}
