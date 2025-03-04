package main

import "fmt"

type Profile struct {
	Id   int
	Name string
}

type ListProfile struct {
	ProfileList []Profile
}

func main() {

	edit := 2
	data := ListProfile{
		[]Profile{
			{Id: 1, Name: "Mikhael"},
			{Id: 2, Name: "Deni"},
			{Id: 3, Name: "Edo"},
		},
	}
	fmt.Println(data)
	for i, val := range data.ProfileList {
		if edit == val.Id {
			data.ProfileList[i].Name = "Monic"
		}
	}
	fmt.Println(data)

}
