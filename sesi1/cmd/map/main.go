package main

import "fmt"

func main() {
	var person = map[string]any{
		"name":    "Budi",
		"address": "Jakarta",
		"age":     20,
		"hobbies": []string{"Reading", "Coding"},
	}

	fmt.Printf("Name: %s, Address: %s, Age: %d\n", person["name"], person["address"], person["age"])
}
