package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 3; j > 0; j-- {
			fmt.Print(i, " ", j, "\n")
		}
	}
}
