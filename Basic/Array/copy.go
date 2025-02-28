package main

import "fmt"

func main() {

	var fruits = make([]string, 3)

	buah := []string{"nanas", "apel", "semangka", "jeruk"}

	copy(fruits, buah)

	fmt.Println("Fruits : ", fruits)
	fmt.Println("buah : ", buah)
	fmt.Println(copy(fruits, buah))

	// dia akan mereplace jika di copy

	buah1 := []string{"apel", "jeruk", "semangka"}
	buah2 := []string{"pisang", "durian"}

	copy(buah1, buah2)

	fmt.Println(buah1) //[pisang durian semangka]
	fmt.Println(buah2) //[pisang durian]
}
