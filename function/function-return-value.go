package main

import "fmt"

func getHello(name string) string {
	result := "Hello"
	return result

}

func getName(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("rizki")
	fmt.Println(result)

	result3 := getName("rizki")
	fmt.Println(result3)

	result2 := getName("")
	fmt.Println(result2)

	fmt.Println(getName("kiki")) //bisa langsung juga tanpa membuat variabel baru lagi
}
