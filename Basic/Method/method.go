package main

import "fmt"

type Person struct {
	Name string
}

// Getter
func (p Person) GetName() string {

	return p.Name
}

// Setter
func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {

	data := Person{}

	data.SetName("Mikhael")
	fmt.Println(data.GetName())
}
