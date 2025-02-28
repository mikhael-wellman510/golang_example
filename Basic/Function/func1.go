package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10) + 1
	var nama []string = []string{"mike", "wellman"}
	cekKesempatan(random)
	sayHello("Hallo ", nama)

}

func sayHello(hello string, val []string) {

	fmt.Println(hello, " , nama saya adalah : ", strings.Join(val, "-"))
}

func cekKesempatan(val int) {
	fmt.Printf("mike mempunyai %d kesempatan \n", val)
}
