package main

import "fmt"

type Users struct {
	Id   int
	Name string
}

type Products struct {
	ProductName string
	Users       Users
}

func main() {

	data1 := Users{
		Id:   1,
		Name: "Mike",
	}

	data2 := Products{
		ProductName: "Indomie",
		Users:       data1,
	}

	fmt.Println(data2)
}
