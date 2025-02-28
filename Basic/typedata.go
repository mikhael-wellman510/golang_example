package main

import "fmt"

func main() {
	/*
		STRING
	*/
	var hobby string = "Futsal"
	fmt.Printf("Hobi saya adalah %s \n", hobby)
	/*
		Integer
		type data bilangan bulat
		kompiler sudah cerdas tanpa harus mendeklarasikan nya lagi
	*/
	var data5 = -2000
	fmt.Printf("hasil bilangan bulat : %d \n", data5)
	/*
		type data numerik desimal
		%.3f = itu menghapus 0 di belakang koma sebanyak 3
	*/

	var decimalNumber = 2.52
	fmt.Printf("hasil %f \n , atau : %.3f \n", decimalNumber, decimalNumber)

	/*
		Type Data Boolean

	*/
	var exist bool = true
	fmt.Printf("Hasil dari Exist adalah : %t \n", exist)

	/*
		Constanta -> Nilai yang sudah pasti dan tidak akan berubah

	*/

	const pi = 3.14

	fmt.Printf("Pi adalah : %f \n", pi)

}
