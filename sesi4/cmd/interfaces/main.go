package main

import "fmt"

func main() {
	var randomValues interface{}

	randomValues = 1
	randomValues = "Hello"
	randomValues = "World"

	fmt.Println(randomValues)

	if value, ok := randomValues.(float64); ok == true {
		fmt.Println("Value is float64", value)
	} else {
		fmt.Println("Value is not float64")
	}
}
