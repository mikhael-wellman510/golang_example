package main

import "fmt"

func main() {

	var num int = 5

	for i := 0; i < num; i++ {
		for j := 0; j < num-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i*2; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
