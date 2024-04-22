package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var argsRaw = os.Args

	// fmt.Printf("-> %#v\n", argsRaw)
	// var firstArg = argsRaw[1]

	// fmt.Println("First Argument:", firstArg)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter firstname: ")

	firstname, _ := reader.ReadString('\n')

	fmt.Println("My Firstname:", firstname)

	fmt.Print("Enter lastname: ")

	lastname, _ := reader.ReadString('\n')

	fmt.Println("My Lastname:", lastname)

}
