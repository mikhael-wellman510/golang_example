package main

import "fmt"

func main() {

	/*Slice ini kalau di java seperti array list */
	var fruits = []string{"Nanas", "Melon", "Durian", "Semangka"}

	/* jika membuat slice baru , dia tetap akan merubah refrensi
	slice lama nya */

	var newSlice []string = fruits[1:3]
	fmt.Println("Slice baru ", newSlice)

	// Refrensi awal slice fruits juga akan berubah
	newSlice[1] = "Pisang"
	fmt.Println(fruits)

}
