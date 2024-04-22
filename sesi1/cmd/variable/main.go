package main

import "fmt"

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

func main() {
	learnVariable()
}
