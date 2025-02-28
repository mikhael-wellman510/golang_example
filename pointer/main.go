package main

import "fmt"

func main() {

	num := 1

	ubah := &num
	*ubah = 8
	fmt.Println(num)
}
