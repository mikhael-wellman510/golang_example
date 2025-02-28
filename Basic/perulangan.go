package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// For range
	var nama string = "mikhaele"
	var splitNama []string = strings.Split(nama, "")
	for i, val := range nama {
		fmt.Println("No : ", strconv.Itoa(i+1), "huruf : ", val)
	}

	for _, val := range splitNama {
		fmt.Println("Hasil split : ", val)
	}

	// For normal
	for i := 0; i < 5; i++ {
		fmt.Println("Hasil i adalah : ", i)
	}

	// Contoh while di golang

	val := 0

	for val < 5 {
		fmt.Println("Hasil nilai while i : = ", val)
		val++
	}

	// Do While

	val2 := 0

	for {
		fmt.Println("Hasil val 2 : ", val2)
		val2++
		if val2 == 5 {
			break
		}

	}

}
