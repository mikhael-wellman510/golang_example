package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	student1 := Student{
		Name: "Mike",
	}

	var student2 *Student = &student1

	student2.Name = "deni"

	fmt.Println(student1)

}
