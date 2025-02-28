package main

import (
	"fmt"
	"strings"
)

func filterName(name []string, callback func(string) bool) []string {

	var result []string = []string{}

	for _, n := range name {

		if callback(n) {
			result = append(result, n)
		}
	}

	return result
}

func main() {

	var name []string = []string{"Mike ", "aldo", "Benny ", "Naomi"}

	var res = filterName(name, func(n string) bool {
		return strings.Contains(n, "o")
	})

	fmt.Println(res)
}
