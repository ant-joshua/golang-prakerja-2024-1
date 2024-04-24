package main

import "fmt"

func sliceWithAppend() {
	fruits2 := make([]string, 3)

	fruits2[0] = "mengkudu"

	fruits2 = append(fruits2, "apple", "grape", "melon")

	// fruits2[0] = "apple"
	// fruits2[1] = "grape"
	// fruits2[2] = "melon"

	fmt.Printf("%#v\n", fruits2)
}

func sliceExample() {
	fruits := []string{"apple", "banana", "grape", "orange", "melon"}

	fmt.Printf("Jumlah element %d\n", len(fruits))
}

func sliceAppendElipsis() {
	fruits := []string{"apple", "banana", "grape", "orange", "melon"}
	fruits2 := []string{"durian", "kiwi"}

	// fruits = append(fruits, fruits2...)

	for _, fruit := range fruits2 {
		fruits = append(fruits, fruit)
	}

	fmt.Printf("%#v\n", fruits)
}

func sliceWithCopy() {
	fruits := []string{"apple", "banana", "grape", "orange", "melon"}
	fruits2 := make([]string, len(fruits))
	copeidElement := copy(fruits2, fruits)
	fruits[0] = "mengkudu"

	fmt.Printf("fruits %#v\n", fruits)
	fmt.Printf("fruits2 %#v\n", fruits2)
	fmt.Printf("copeidElement %d\n", copeidElement)
}

func sliceSlicing() {
	fruits := []string{"apple", "banana", "grape", "orange", "melon"}

	slice1 := fruits[1:2]
	slice2 := fruits[0:]
	slice3 := fruits[:3]
	// slice2 := fruits[0:]
	// slice3 := fruits[:3]

	fmt.Printf("slice1 %#v\n", slice1)
	fmt.Printf("slice2 %#v\n", slice2)
	fmt.Printf("slice3 %#v\n", slice3)
}

func sliceSlicingAndAppend() {
	var fruits1 = []string{"apple", "banana", "grape", "orange", "melon"}

	fruits1 = append(fruits1[:3], "rambutan")

	fmt.Printf("%#v\n", fruits1)
}

func main() {
	// sliceAppendElipsis()
	// sliceSlicing()
	sliceSlicingAndAppend()
}
