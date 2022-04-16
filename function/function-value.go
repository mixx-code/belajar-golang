package main

import "fmt"

func getGoodBye(nama string) string {
	return "Good Bye " + nama
}

func main() {
	goodBye := getGoodBye

	result := goodBye("rizki")
	fmt.Println(result)
	fmt.Println(getGoodBye("kiki"))

}
