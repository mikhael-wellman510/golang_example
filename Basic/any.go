package main

import "fmt"

func main() {

	/*
		Tipe data interface{} atau any adalah tipe data yang spesial
		karena bisa menampung segala jenis data baik:
		numerik , string bahkan array pointer
	*/

	var secret interface{}

	secret = "Mikhael Wellman"
	fmt.Println("String : ", secret)

	secret = 140
	fmt.Println("Umur  : ", secret)
	fmt.Println(secret)

	secret = []string{"Mike", "Denis", "Aldo"}
	fmt.Println("Array : ", secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"nama": "Mike",
		"umur": 18,
	}

	fmt.Println("hasil map : ", data)
}
