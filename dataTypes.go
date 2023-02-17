// Go Data Types
// Variables can be of different types like int, float, struct, slice or it can be of the interface.

// The general form for declaring a variable uses the keyword var:

// Syntax:-

// var identifier type
// Variables can be of different types like int, float, struct, slice or it can be of the interface.

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%T %T %T %T\n", i, f, b, s)              // Prints type of the variable
	fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s) //prints initial value of the variable
}
