package main

import "fmt"

func main() {

	/* cara hapus data yg sering jika dapa idx nya */
	fruits := []string{"Apel", "Melon", "Durian", "semangka", "Kudu", "Salak"}
	idx := 2
	dataBaru := append(fruits[:idx], fruits[idx:]...)
	fmt.Println(dataBaru)

	/* Cara menambahkan data */

	var newFruits []string = append(fruits, "nanas")
	fmt.Println(newFruits)

	/* cara hapus yg paling terakhir*/
	fmt.Println("total : ", fruits)
	fruits = fruits[:len(fruits)-1]
	fmt.Println("Paling terakhir : ", fruits)

	/* Cara hapus yang paling depan*/
	fruits = fruits[1:]
	fmt.Println("paling depan ", fruits)

}
