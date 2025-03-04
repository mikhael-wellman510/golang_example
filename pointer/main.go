package main

import (
	"fmt"
	example "latihan-golang/pointer/Example"
)

func main() {

	// tes := example.NewData()

	cek := example.NewData()

	data := example.Product{
		Nama: "Beras",
		Qty:  1,
	}

	cek.AddProduct(data)
	cek.AddProduct(data)
	fmt.Println("cek data : ", cek.FindAll())
	num := 1

	ubah := &num
	*ubah = 8
	fmt.Println(num)
}
