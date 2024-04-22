package main

import "fmt"

type Employee struct {
	Firstname string
	Lastname  string
	Address   string
	Age       int
	Hobbies   []string
}

func (e Employee) GetFullName() string {
	return e.Firstname + " " + e.Lastname
}

func (e Employee) GetHobbies() string {
	var hobbies string
	for _, hobby := range e.Hobbies {
		hobbies += hobby + ", "
	}
	return hobbies
}

func (e Employee) CountHobbies() int {
	return len(e.Hobbies) // [].length
}

func main() {
	var employee = Employee{
		Firstname: "Budi",
		Lastname:  "Santoso",
		Address:   "Jakarta",
		Age:       20,
		Hobbies:   []string{"Reading", "Coding"},
	}

	fmt.Println(employee.CountHobbies())
}
