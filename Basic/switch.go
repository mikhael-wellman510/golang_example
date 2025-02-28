package main

import "fmt"

func main() {

	var angka int = 8

	// Switch Standar
	switch angka {

	case 9:
		fmt.Println("Perfect ")

	case 8:
		fmt.Println("Okay")
	case 5:
		fmt.Println("Not Bad")

	}

	// Switch dengan menggunakan Kondisi
	switch {

	case angka > 7:
		fmt.Println("Perfect")
	case angka <= 7:
		fmt.Println("Maaf anda Remedial")

	default:
		fmt.Println("Angka tidak di ketahui")

	}

	// menggunakan Fallthrought
	// Dia akan membuat kode berikut nya tereksekusi tanpa melihat kondisi nya

	var nilai int = 90

	switch {

	case nilai >= 90:
		fmt.Println("Nilai anda cukup Besar")
		fallthrough
	case nilai >= 80:
		fmt.Println("Nilai anda cukup bagus")

	case nilai >= 70:
		fmt.Println("Nilai anda cukup baik")

	case nilai >= 60:
		fmt.Println("Nilai anda masih kurang")

	default:
		fmt.Println("Selesai pengecekan nilai")

	}
}
