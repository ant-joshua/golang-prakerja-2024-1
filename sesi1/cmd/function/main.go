package main

import (
	"errors"
	"fmt"
)

const (
	ErrorFirstAndLastNameEmpty = "firstname and lastname is empty"
)

func GenerateFullName(firstname, lastname string) (string, error) {

	if firstname == "" && lastname == "" {
		return "", errors.New(ErrorFirstAndLastNameEmpty)
	}

	return firstname + " " + lastname, nil
}

func main() {
	fullname, err := GenerateFullName("Budi", "Santoso")

	if err != nil {

		if err == errors.New(ErrorFirstAndLastNameEmpty) {
			fmt.Println("Error:", ErrorFirstAndLastNameEmpty)
			return
		}

		fmt.Println("Error:", err)
	}

	fmt.Println(fullname)
}
