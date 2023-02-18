package main

import "fmt"

func main() {
	fmt.Print("Enter text: ")
	var input int
	fmt.Scanln(&input)
	if input < 0 || input > 100 {
		fmt.Print("Please enter valid no")
	} else if input >= 0 && input < 50 {
		fmt.Print(" Fail\n")
	} else if input >= 50 && input < 60 {
		fmt.Print(" D Grade\n")
	} else if input >= 60 && input < 70 {
		fmt.Print(" C Grade\n")
	} else if input >= 70 && input < 80 {
		fmt.Print(" B Grade\n")
	} else if input >= 80 && input < 90 {
		fmt.Print(" A Grade\n")
	} else if input >= 90 && input <= 100 {
		fmt.Print(" A+ Grade\n")
	}
}
