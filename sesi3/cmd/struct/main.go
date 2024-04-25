package main

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Employee struct {
	Person   Person `json:"person"` // public field
	Division string `json:"division"`
}

type EmployeeList []Employee

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var employee Employee

	employee.Person.Name = "Joshua"
	employee.Person.Age = 25
	employee.Division = "Engineering"

	var employee2 = Employee{}
	employee2.Person.Name = "Michael"
	employee2.Person.Age = 30
	employee2.Division = "Marketing"

	var employee3 *Employee = &employee

	employee3.Person.Name = "John"
	employee3.Person.Age = 35

	fmt.Println(employee.Person.Name)
	fmt.Println(employee.Person.Age)

	employeeJson, err := json.Marshal(employee)
	if err != nil {
		fmt.Println("cannot convert to json", err)
		return
	}
	fmt.Println(string(employeeJson))

	baseResponse := BaseResponse{
		Code:    200,
		Message: "Success",
		Data:    employee,
	}

	updateDataBaseResponse := baseResponse.Data.(Employee)
	updateDataBaseResponse.Person.Name = "Doe"

	baseResponse.Data = updateDataBaseResponse

	baseResponseJson, err := json.Marshal(baseResponse)
	if err != nil {
		fmt.Println("cannot convert to json", err)
		return
	}

	fmt.Println(string(baseResponseJson))

	var company = struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}{}

	company.Name = "PT. Gojek Indonesia"
	company.Address = "Jakarta, Indonesia"

	companyJson, err := json.Marshal(company)
	if err != nil {
		fmt.Println("cannot convert to json", err)
		return
	}
	fmt.Println(string(companyJson))

	employeeList := EmployeeList{}

	employeeList = append(employeeList, employee)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, updateDataBaseResponse)

	employeeListJson, err := json.Marshal(employeeList)
	if err != nil {
		fmt.Println("cannot convert to json", err)
		return
	}

	fmt.Println(string(employeeListJson))

}
