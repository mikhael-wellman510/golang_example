package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countingSort(arr []int) []int {

	max := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("hasil max : ", max)
	result := make([]int, 100)

	for i := 0; i < len(arr); i++ {
		var cek int = arr[i]

		result[cek] = result[cek] + 1

	}
	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	total, _ := reader.ReadString('\n')
	total = strings.TrimSpace(total)

	name, _ := reader.ReadString('\n')

	num, _ := strconv.Atoi(total)
	nameArr := strings.Fields(name)

	arrData := []int{}

	for i := 0; i < num; i++ {

		nums, _ := strconv.Atoi(nameArr[i])
		arrData = append(arrData, nums)
	}
	res := countingSort(arrData)

	var hasil strings.Builder

	for _, i := range res {
		hasil.WriteString(strconv.Itoa(i) + " ")
	}

	fmt.Println(hasil.String())

}
