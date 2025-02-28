package main

import (
	"fmt"
	"strings"
)

func main() {

	var hasil float64 = nilaiRataRata(3, 4, 6, 8, 5, 1)
	fmt.Println("Hasil : ", hasil)
	var hobby []string = []string{"futsal", "gym", "basket"}
	var res string = myHobby("Mike", hobby...)
	fmt.Println(res)
}
func myHobby(nama string, hobby ...string) string {

	return "Hallo , My name " + nama + " hobi saya : " + strings.Join(hobby, "-")
}

func nilaiRataRata(nilai ...int) float64 {
	panjang := len(nilai)

	fmt.Println("Params : ", nilai)

	val := 0

	for _, num := range nilai {
		val = val + num
	}

	return float64(val) / float64(panjang)
}
