package main

import "fmt"

type Profile struct {
	Name    string
	Address string
}

func changeData(data *Profile, change string) {
	data.Address = change
}

func main() {
	data := Profile{
		Name:    "Mikhael",
		Address: "Bogor",
	}

	changeData(&data, "Depok")
	fmt.Print(data)
}
