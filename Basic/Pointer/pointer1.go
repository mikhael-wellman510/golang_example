package main

import "fmt"

func changePointer(num *int, angka int) {
	// Kirim alamat nya
	*num = angka
}

func main() {
	var number int = 3
	var change int = 10
	changePointer(&number, change)

	fmt.Println(number)
}
