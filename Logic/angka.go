package main

import "fmt"

func main() {
	var data int = 5
	for i := data; i > 0; i-- {
		fmt.Print(i - (data - 1))
		for j := 0; j >= i-(data); j-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}
