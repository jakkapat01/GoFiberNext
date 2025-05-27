package main

import (
	"fmt"
)

func LogicalOperator() {
	// &&, ||, !
	var a bool = true
	var b bool = false
	var c bool

	c = (a && b) // false
	fmt.Println("a && b : ", c)

	c = (a || b) // true
	fmt.Println("a || b : ", c)

	c = !a // false
	fmt.Println("!a : ", c)
}
