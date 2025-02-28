package main

import "fmt"

func main() {

	var num int = 5

	for i := 0; i < num; i++ {

		for j := num; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	for a := 0; a < num; a++ {

		for c := 0; c < a; c++ {
			fmt.Print(" ")
		}
		for b := num; b > a; b-- {
			fmt.Print("*")
		}

		fmt.Print("\n")
	}

}
