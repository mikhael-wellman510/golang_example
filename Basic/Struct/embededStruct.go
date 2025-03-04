package main

import "fmt"

type Students struct {
	Name string
	Jurusan
}

type Jurusan struct {
	NamaJurusan string
	Prodi       string
}

func main() {
	var data Students = Students{}

	data.Name = "Mikhael"
	data.Jurusan.NamaJurusan = "Multimedia"
	data.Jurusan.Prodi = "Design Grafis"

	fmt.Println(data)
}
