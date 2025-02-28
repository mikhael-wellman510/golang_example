package main

import "fmt"

func findMax(max int, val ...int) (int, func() []int) {

	var res []int = []int{}

	for _, num := range val {
		if max < num {
			res = append(res, num)

		}
	}

	return len(res), func() []int {
		return res
	}

}

func main() {
	var arr []int = []int{2, 4, 6, 5, 2, 1, 3, 8, 6, 3, 1}
	var max int = 3

	var length, num = findMax(max, arr...)
	fmt.Println(length)
	fmt.Println(num())

}
