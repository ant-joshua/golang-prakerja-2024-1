package main

import "fmt"

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

func main() {
	learnCondition()
	learnCondition2()
	learnSwitch()
	learnSwitch2()
}
