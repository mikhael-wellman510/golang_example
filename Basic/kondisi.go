package main

import "fmt"

func main() {

	/*
		Variable Temporary di blok kondisi
	*/

	var uas int = 92
	var un int = 70

	if val := (uas + un) / 2; val > 90 {
		fmt.Printf("Nilai tinggi :  %d \n", val)
	} else if val > 70 {
		fmt.Printf("Nilai sedang : %d \n", val)
	} else {
		fmt.Println("gagal , Karena nilai nya : %d \n", val)
	}

	var point int = 8

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")

	} else {
		fmt.Println("Tidak lulus")
	}
}
