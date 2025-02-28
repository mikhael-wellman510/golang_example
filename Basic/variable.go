package main

import (
	"fmt"
)

func main() {
	/*
		TYPE INFERENCES
		tipe data tidak perlu di tuliskan , karena sudah diketahui secara otomatis
		kalau pakai printF , dia perlu di kasih /n , karena dia tidak auto membuat line baru
		berbedan dengan printLn yg auto membuat garis baru
	*/
	name := "mikhael"
	fmt.Printf("Nama saya adalah : %s\n", name)
	one, two, tree := "satu", "dua", "tiga"
	fmt.Printf("1: %s , 2: %s , 3: %s \n", one, two, tree)

	/*,
	nama variable ini adalah MANIFEST TYPING
	dia mendeklarasikan variable dan type data
	*/
	var firstName string = "mikhael"
	var lastname string
	lastname = "Wellman"

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Printf("Nama saya : %s \ndan nama belakang : %s\n", firstName, lastname)

	fmt.Printf("first : %s , second : %s , third : %s", first, second, third)

	/*

		VARIABLE UNDERSCORE
		di golang , type data tidak boleh kosong karena akan error
		biasa nya untuk menampung nilai balik fungsi yang tidak di gunakan
	*/

	_ = "variable yang tidak terpakai"

	nameBaru, _ := "mikhael", "wellman"
	fmt.Println("nama baru : ", nameBaru)

	/*
		Deklarasi variable dengan menggunakan new
	*/

	nama1 := new(string)
	fmt.Println(nama1)
	fmt.Println(*nama1)

}
