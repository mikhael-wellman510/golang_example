package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findMedian(arr []int) int {

	sort.Ints(arr)
	res := arr[(len(arr) / 2)]

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)
	numInt, _ := strconv.Atoi(num)
	arrNum, _ := reader.ReadString('\n')
	arrNum = strings.TrimSpace(arrNum)

	arr := strings.Split(arrNum, " ")
	fmt.Println("cek arr : ", arr)

	var data []int
	for i := 0; i < numInt; i++ {
		nums, _ := strconv.Atoi(arr[i])
		data = append(data, nums)
	}
	fmt.Println("Cek data : ", data)
	hasil := findMedian(data)

	fmt.Println(hasil)
}
