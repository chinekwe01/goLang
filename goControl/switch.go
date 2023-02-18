package main

import "fmt"

func main() {
	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 10:
		fmt.Print("The value is 10\n")
	case 20:
		fmt.Print("The value is 20\n")
	case 30:
		fmt.Print("The value is 30\n")
	case 40:
		fmt.Print("The value is 40\n")
	default:
		fmt.Print("The value is not 10,20,30,or 40\n")
	}
}
