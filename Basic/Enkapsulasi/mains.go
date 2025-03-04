package main

import (
	"fmt"
	data "latihan-golang/Basic/Enkapsulasi/Data"
)

func main() {
	instanceObject := data.Data{}

	// Panggil constructor
	tes := data.NewData()
	fmt.Println("cek name ", tes.GetName())

	// Ini buat setter getter
	instanceObject.SetName("Mikhael")
	instanceObject.SetAlamat("Bogor")

	fmt.Println(instanceObject.GetName())
	fmt.Println(instanceObject.GetAlamat())

}
