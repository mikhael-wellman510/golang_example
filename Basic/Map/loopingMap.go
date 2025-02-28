package main

import "fmt"

func main() {
	var warteg map[string]int = map[string]int{
		"nasi":         3000,
		"Ayam Goreng":  10000,
		"Sayur Bayem":  4000,
		"Es Teh Manis": 2000,
	}

	for key, value := range warteg {
		fmt.Println("Harga ", key, " Adalah : ", value)
	}

	// Mencari isi map
	val, isExist := warteg["nasi"]
	fmt.Println("val : ", val, "isExist : ", isExist)

	// Menentukan panjang map

	fmt.Println(len(warteg))

	// menghapus
	delete(warteg, "nasi")

	for key, val := range warteg {
		fmt.Println("Nama : ", key, " harga : ", val)
	}

}
