package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //untuk mendapatkan panjang
	fmt.Println(cap(slice1)) //untuk mendapatkan kapasitas

	// dapat mengubah array
	// bulan[5] = "ubah"
	// fmt.Println(slice1)

	// mengubah dari slice dapat mengubah arraynya
	// slice1[0] = "ubah"
	// fmt.Println(bulan)

	//append()
	var slice2 = bulan[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "rizki")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)

	//membuat slice baru
	sliceBaru := make([]string, 2, 5)

	sliceBaru[0] = "rizki"
	sliceBaru[1] = "kiki"

	fmt.Println(sliceBaru)
	fmt.Println(len(sliceBaru))
	fmt.Println(cap(sliceBaru))

	//copy slice

	copySlice := make([]string, len(sliceBaru), cap(sliceBaru)) // panjang data harus sama agar tidak kepotong
	copy(copySlice, sliceBaru)

	fmt.Println(copySlice)
}
