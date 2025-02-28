package main

import "fmt"

func changePoint(point *int) {
	*point = 200
}

type Biodata struct {
	Name, Alamat string
}

func main() {
	point := 100

	changePoint(&point)

	fmt.Println(point)

	bio := Biodata{Name: "Mik", Alamat: "bgr"}
	bio2 := &bio
	bio2.Alamat = "bandung"

	fmt.Println(bio)
	fmt.Println(bio2)

}
