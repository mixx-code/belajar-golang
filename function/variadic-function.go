package main

import "fmt"

func sumAll(numbers ...int) int { //variabel argumen hanya bisa di paling belakang
	total := 0
	for _, value := range numbers {
		total += value

	}
	return total
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	//data dalam slice dapat di panggil dengan variadic function
	slice := []int{10, 20, 10}
	total = sumAll(slice...)

}
