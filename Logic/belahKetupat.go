package main

import "fmt"

func main() {

	var num int = 11

	for i := 1; i < num/2; i++ {

		for j := 0; j < (num/2)-i; j++ {
			fmt.Print("*")
		}

		for k := 0; k <= (num%i)*2; k++ {
			fmt.Print("x")
		}

		for l := 0; l < (num/2)-i; l++ {
			fmt.Print("*")
		}
		fmt.Print("\n")

	}

	for i := num / 2; i > 2; i-- {

		for k := 1; k >= i-(num/2); k-- {
			fmt.Print("*")
		}
		for j := 0; j < ((((num) % (i - 1)) * 2) - 1); j++ {
			fmt.Print("x")
		}

		for k := 1; k >= i-(num/2); k-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	fmt.Println(11 % 4)
}
