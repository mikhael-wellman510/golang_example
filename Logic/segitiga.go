package main

import "fmt"

func main() {

	var num int = 5

	for i := 0; i < num; i++ {

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
