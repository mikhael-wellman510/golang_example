package main

import "fmt"

type StudentsNew struct {
	Name   string
	School struct {
		NameSchool string
		Address    string
	}
}

func main() {

	var student StudentsNew

	student.Name = "Mike"
	student.School.Address = "Bogor"
	student.School.NameSchool = "SMPN 20 BOGOR"

	fmt.Println(student)
}
