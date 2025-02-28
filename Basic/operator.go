package main

import "fmt"

func main() {

	/*
		Operator Perbandingan
		 != , == , < , > , <= , >=
	*/
	var num1 = 2
	var num2 = 4

	var hasil1 bool = num1 == num2
	fmt.Printf("hasil 1 : %t \n", hasil1)

	var hasil2 bool = num1 != num2
	fmt.Printf("hasil 2 : %t \n", hasil2)

	var hasil3 bool = num1 > num2
	var hasil4 bool = num1 < num2

	fmt.Printf("Hasil 3 : %t \n", hasil3)
	fmt.Printf("hasil 4 : %t \n", hasil4)

	/*
		Operator Logika && , || , !
	*/
	var left bool = true
	var right bool = false

	var hasilfalse bool = left && right
	var hasilTrue bool = left || right
	var hasilLogika bool = !right

	fmt.Printf("f : %t , t : %t , l : %t ", hasilfalse, hasilTrue, hasilLogika)
}
