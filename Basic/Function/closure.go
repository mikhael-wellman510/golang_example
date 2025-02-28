package main

import "fmt"

func main() {

	data := []int{23, 40, 11, 3, 6, 77, 34, 22}

	var terendahDanTertinggi = func(val ...int) (int, int) {

		max := 0
		min := 0

		for i, num := range val {

			switch {
			case i == 0:
				max = num
				min = num
			case num > max:
				max = num

			case num < min:
				min = num
			}
		}

		return max, min
	}

	max, min := terendahDanTertinggi(data...)

	fmt.Printf("Hasil tertinggi adalah %d dan terendah adalah %d \n", max, min)

	var luasDanKeliling = func(p int, l int) (int, int) {

		var keliling int = 2 * (p + l)
		var luas int = p * l
		return luas, keliling

	}

	luas, keliling := luasDanKeliling(3, 4)
	fmt.Printf("Luas nya adalah %d \n", luas)
	fmt.Printf("dan keliling nya adalah : %d ", keliling)
}
