package main

import (
	"fmt"
)

func learnVariable() {
	var firstname, lastname string
	firstname, lastname = "Budi", "Santoso"
	// var age = 20

	age := 20
	score := 80.5

	// age = "halo"

	fmt.Printf("Name: %s %s Umur: %d, Score: %.1f\n", firstname, lastname, age, score)

	// fmt.Println(name)
	// fmt.Println("umur saya", age)
}

func learnCondition() {
	var isComplete bool = true

	if isComplete {
		fmt.Println("Task is complete")
	} else {
		fmt.Println("Task is not complete")
	}
}

func learnCondition2() {
	var currentYear = 2021

	if age := currentYear - 1990; age > 30 {
		fmt.Println("You are old")
	} else {
		fmt.Println("You are young")
	}
}

func learnSwitch() {
	var day = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown")
	}
}

func learnSwitch2() {
	var score = 8.5

	switch {
	case score == 10:
		fmt.Println("Perfect")
	case score < 10 && score >= 5:
		fmt.Println("Lulus")
		fallthrough
	case score >= 6:
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Study more")
		}
	}
}

func learnMap() {

	var employee struct {
		Firstname string
		Lastname  string
		Address   string
		Age       int
	}

	employee.Firstname = "Budi"
	employee.Lastname = "Santoso"
	employee.Address = "Jakarta"
	employee.Age = 20

	// {} , Record<string, string>

	// fmt.Printf("Name: %s, Address: %s, Age: %d\n", person["name"], person["address"], person["age"])
}

func main() {

	learnMap()

	// var age = 20
	// var umur = 40

	// var age2 *int = &age

	// age = 10

	// fmt.Println(age)
	// fmt.Println(&umur)

	// fmt.Println(*age2)
	// var age2 *int = &age

	// learnVariable()

	// fmt.Print("Hello, World!\n")
	// fmt.Print("Hello, World!\n")

	// // WelcomeToLibrary()
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// // welcomeToLibrary()

	// r.Run(":5000")
}
