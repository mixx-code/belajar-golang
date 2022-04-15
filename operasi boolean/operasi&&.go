package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	// var lulusUjian = ujian > 75
	// var lulusAbsensi = absensi > 75

	// var lulus = lulusUjian && lulusAbsensi

	// fmt.Println(lulus)

	//biasanya penulisannya seperti ini

	fmt.Println(ujian >= 80 && absensi >= 80)

}
