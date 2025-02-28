package main

import (
	"fmt"
)

type Profile struct {
	Nama string
	Usia int
}

func editData(val *Profile) *Profile {

	return &Profile{
		Nama: "Deni",
		Usia: 18,
	}
}

func main() {

	val := Profile{}
	data := editData(&val)

	fmt.Println(data)

	datas := 10

	var hasil *int = &datas

	fmt.Println(*hasil)
}
