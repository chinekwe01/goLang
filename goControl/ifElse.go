package main

import (
	"fmt"
)

func main() {
	/* local variable definition */
	var a int = 10
	// check the boolean condition using if statement

	/*
		Go If
		The if statement in Go is used to test the condition. If it evaluates to true, the body of the statement is executed. If it evaluates to false, if block is skipped.

		Syntax :

		if(boolean_expression) {

			   // statement(s) got executed only if the expression results in true
			}
	*/
	if a%2 == 0 {
		fmt.Println(a, "is an even number")
	}
}
