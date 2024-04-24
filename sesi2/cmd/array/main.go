package main

import "fmt"

func multidimensionalArray() {
	balances := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for _, balance := range balances {
		fmt.Println("balance", balance)
		for _, value := range balance {
			fmt.Println(value)
		}
	}

	contoh := [2][2]interface{}{
		{1, "a"},
		{2, "b"},
	}

	fmt.Printf("contoh %#v\n", contoh)
}

func arrayExample() {
	var numbers [5]int

	numbers[0] = 1
	numbers[1] = 2

	numbers[0] = 3

	fmt.Println(numbers)
	// len
	fmt.Println(len(numbers))
	// cap
	fmt.Println(cap(numbers))

	for i := 0; i < len(numbers); i++ {
		fmt.Println("array with loop", numbers[i])
	}

	for index, value := range numbers {
		fmt.Println("array with range", index, value)
	}
}

func main() {
	// arrayExample()
	multidimensionalArray()
}
