package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func contohError() {
	number, err := strconv.Atoi("A10")

	if err != nil {
		println("Error:", err)

		log.Fatalf("Error: %v", err)
	}

	println(number)
}

func main() {

	defer catchError()

	var password string

	fmt.Scanln(&password)

	if valid, err := validatePassword(password); !valid {

		if errors.Is(err, ErrPasswordTooShort) {
			println("Password is too short from errors.Is")
		}

		var errPasswordShort PasswordShortError

		if errors.As(err, &errPasswordShort) {
			println("Password is too short from errors.As")
		}

		panic(err)
	}

	println("Password is valid")
}

// const (
// 	ErrPasswordTooShort = errors.New("password is too short")
// )

var ErrPasswordTooShort = errors.New("password is too short")

type PasswordShortError struct {
	Input string
}

func (e PasswordShortError) Error() string {
	return fmt.Sprintf("password %s is too short", e.Input)
}

func validatePassword(password string) (bool, error) {
	pl := len(password)

	if pl < 5 {
		return false, PasswordShortError{Input: password}
	}

	return true, nil
}

func catchError() {

	if r := recover(); r != nil {
		println("Recovered from", r)
	}
	// defer func() {

	// }()

}
