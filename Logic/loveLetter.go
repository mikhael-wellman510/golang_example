package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func theLoveLetterMystery(s string) int {

	pointer := len(s) / 2
	strS := strings.Split(s, "")
	result := 0
	for i := 0; i < pointer; i++ {

		if strS[i] != strS[(len(s)-1)-i] {

			left := int(strS[i][0])
			right := int(strS[((len(s) - 1) - i)][0])

			if left > right {
				result += left - right
			} else {
				result += right - left
			}

		}
	}

	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	numTrim := strings.TrimSpace(num)
	numLen, _ := strconv.Atoi(numTrim)

	for i := 0; i < numLen; i++ {
		str, _ := reader.ReadString('\n')
		strTrim := strings.TrimSpace(str)
		fmt.Println(theLoveLetterMystery(strTrim))
	}

}
