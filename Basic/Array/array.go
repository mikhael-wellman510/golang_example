package main

import "fmt"

func main() {

	var fruits = [4]string{"nanas", "mangga", "rambutan", "melon"}

	// panjang arrays
	var panjang int = len(fruits)
	fmt.Printf("Panjang array %d \n", panjang)

	/* jika ingin membuat tanpa panjang nya gunakan ... */

	var animal = [...]string{"Kuda", "gajah", "Ikan", "Jerapah"}
	var panjangAnimal int = len(animal)
	fmt.Println(animal)
	fmt.Println(panjangAnimal)

	/*Array multidimensi*/

	var multidimensi = [2][3]int{[3]int{1, 2, 3}, [3]int{2, 4, 5}}
	fmt.Println("Multidimensi : ", multidimensi)

	/* perulangan di array */
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("hasil buah %s \n", fruits[i])
	}

	/* menggunakan range */
	for i, val := range fruits {
		fmt.Printf("No %d.%s \n", i+1, val)
	}

}
