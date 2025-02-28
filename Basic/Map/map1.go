package main

import "fmt"

func main() {
	/* map itu menggunakan key dan value
	key nya harus bersifat uniq , dan dia akan me-replace jika sama
	*/
	// Cara 1
	fruits := map[string]int{}

	fruits["melon"] = 10000
	fruits["semangka"] = 20000

	fmt.Println("Fruits : ", fruits)
	fmt.Println("harga melon :", fruits["melon"])

	// Cara 2

	var baju map[string]int = map[string]int{}

	baju["kemeja"] = 50000
	baju["batik"] = 100000

	fmt.Println("Baju : ", baju["kemeja"])

}
