package main

import "fmt"

type Product struct {
	Id    int
	Nama  string
	Price int
}
type User struct {
	Id       int
	Username string
}

type ListProduct struct {
	Product []Product
	User    []User
}

func main() {

	datas := []Product{
		{Id: 1, Nama: "Mike", Price: 10000},
		{Id: 2, Nama: "Sabun", Price: 5000},
	}

	fmt.Println(datas)
	listData1 := ListProduct{
		[]Product{
			{Id: 1, Nama: "Shampo", Price: 10000},
			{Id: 2, Nama: "Sabun", Price: 5000},
		}, []User{
			{Id: 1, Username: "Deni"},
			{Id: 2, Username: "Donny"},
		},
	}
	listData2 := ListProduct{
		[]Product{
			{Id: 3, Nama: "Beras", Price: 10000},
			{Id: 4, Nama: "Odol", Price: 5000},
		}, []User{
			{Id: 3, Username: "Alvaro"},
			{Id: 4, Username: "Faris"},
		},
	}
	data := []ListProduct{
		listData1, listData2,
	}
	for i, val := range data {
		fmt.Println("Data ", i, "adalah : ", val)
	}

}
