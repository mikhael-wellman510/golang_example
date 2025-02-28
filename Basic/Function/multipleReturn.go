package main

import "fmt"

func main() {

	luas, keliling := persegiPanjang(4, 5)

	fmt.Printf("Hasil luas :%d , Hasil Keliling : %d ", luas, keliling)
}

func persegiPanjang(p int, l int) (int, int) {

	var keliling int = 2 * (p + l)
	var luas int = p * l

	return luas, keliling

}
