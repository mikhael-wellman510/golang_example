package main

import "fmt"

func main() {

	var mahasiswa []map[string]string = []map[string]string{
		map[string]string{"nama": "mikhael", "usia": "20"},
		map[string]string{"nama": "deni", "usia": "25"},
		map[string]string{"nama": "alvaro", "usia": "25"},
	}

	for _, data := range mahasiswa {
		fmt.Println(data["nama"])
	}

	
}
