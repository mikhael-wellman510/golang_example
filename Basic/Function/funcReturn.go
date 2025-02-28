package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		var res int = generatedRandomValue()
		fmt.Printf("Hasil angka random nya adalah : %d \n", res)
	}

}

func generatedRandomValue() int {
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	angkaRandom := random.Intn(10) + 1

	return angkaRandom
}
