package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var noKTPrizki NoKTP = "1223123213123"
	var marriedstatus married = false
	fmt.Println(noKTPrizki)
	fmt.Println(marriedstatus)

}
