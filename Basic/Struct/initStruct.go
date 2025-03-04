package main

import "fmt"

type Profile struct {
	// Jika string tidak di isi value nya  ""
	Name string

	// Jika int tidak di isi value nya 0
	Age int
}

func main() {

	// todo-> init struct
	// Cara 1
	var profile = Profile{}

	profile.Name = "Mikhael"
	profile.Age = 28

	fmt.Println("Profile 1: ", profile)

	// Cara 2
	var profile2 = Profile{
		Name: "Denis",
		Age:  25,
	}

	fmt.Println("Profile 2 : ", profile2)
}
