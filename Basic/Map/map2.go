package main

import "fmt"

func main() {
	var makanan map[string]int = map[string]int{
		"nasi":  4000,
		"ayam":  5000,
		"sayur": 3000,
	}

	fmt.Println("Harga makanan : ", makanan)
	fmt.Println("harga nasi : ", makanan["nasi"])

}
