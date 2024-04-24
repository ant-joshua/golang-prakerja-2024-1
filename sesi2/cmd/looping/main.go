package main

import "fmt"

func loopingFor() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Loop For", i)
	}
}

func loopingWhile() {
	i := 0
	for i <= 10 {
		fmt.Println("Loop While", i)
		i++
	}
}

func loopingDoWhile() {
	i := 0
	for {
		fmt.Println("Loop Do While", i)
		i++
		if i > 10 {
			break
		}
	}
}

func nestedLoop() {
	for i := 0; i > 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println(i)
	}
}

func loopWithLabel() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

func loopWithRange() {
	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Println("Index", index, "Value", value)
	}
}

func main() {
	// loopingFor()
	// loopingWhile()
	// loopingDoWhile()
	// nestedLoop()
	// loopWithLabel()
	loopWithRange()
}
