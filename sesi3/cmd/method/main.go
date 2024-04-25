package main

import "fmt"

type Employee struct {
	Person   Person `json:"person"` // public field
	Division string `json:"division"`
	Salary   int    `json:"salary"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Company struct {
	Employees []Employee `json:"employees"`
}

func (s *Company) AddEmployee(employee Employee) {
	s.Employees = append(s.Employees, employee)
}

func (s *Company) RemoveEmployee(index int) {
	s.Employees = append(s.Employees[:index], s.Employees[index+1:]...)
}

func (s *Company) CountEmployee() int {
	return len(s.Employees)
}

func (s *Company) CalculateSalaryWithTax(tax float64) float64 {
	totalSalary := 0.0
	for _, employee := range s.Employees {
		totalSalary += float64(employee.Salary) - (float64(employee.Salary) * tax)
	}
	return float64(totalSalary)
}

func init() {
	fmt.Println("Initializing...")
}

func main() {
	employee := Employee{
		Person: Person{
			Name: "Joshua",
			Age:  25,
		},
		Division: "Engineering",
		Salary:   1000000,
	}

	employee2 := Employee{
		Person: Person{
			Name: "Michael",
			Age:  30,
		},
		Division: "Marketing",
		Salary:   2000000,
	}

	employee3 := &Employee{
		Person: Person{
			Name: "John",
			Age:  35,
		},
		Division: "Human Resource",
		Salary:   3000000,
	}

	company := Company{}

	company.AddEmployee(employee)
	company.AddEmployee(employee2)
	company.AddEmployee(*employee3)

	fmt.Println(company.CountEmployee())

	fmt.Printf("total company pay : %f", company.CalculateSalaryWithTax(0.1))

}
