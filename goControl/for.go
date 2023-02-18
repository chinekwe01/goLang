package main

import "fmt"

func main() {
	// declaring and initializing sum as integer with initial value of 1
	sum := 1
	// var sum int = 1
	// conditioned controlled iteration
	for sum < 100 {
		sum += sum
		fmt.Print(sum, "\n")
	}
}
