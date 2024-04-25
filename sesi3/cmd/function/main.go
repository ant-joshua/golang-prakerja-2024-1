package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func calculate(a int, b ...int) int {
	sum := a
	for _, value := range b {
		sum += value
	}
	return sum
}

func calculateCircle(r float64) (float64, float64, error) {
	if r <= 0 {
		return 0, 0, fmt.Errorf("radius must be greater than 0")
	}
	area := 3.14 * r * r
	circumference := 2 * 3.14 * r
	return area, circumference, nil
}

func calculateRectacle(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return area, perimeter
}

func main() {
	fmt.Println(greet("Joshua"))

	fmt.Println(calculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	var area, circumference, err = calculateCircle(10.5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)

	fmt.Println(calculateRectacle(10, 5))
}
