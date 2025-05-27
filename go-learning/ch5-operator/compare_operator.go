package main

import (
	"fmt"
)

func CompareOperator() {
	// ==, !=, >, <, >=, <=
	var a int = 10
	var b int = 20
	var c bool

	c = (a == b) // false
	fmt.Println("a == b : ", c)

	c = (a != b) // true
	fmt.Println("a != b : ", c)

	c = (a > b) // false
	fmt.Println("a > b : ", c)

	c = (a < b) // true
	fmt.Println("a < b : ", c)

	c = (a >= b) // false
	fmt.Println("a >= b : ", c)

	c = (a <= b) // true
	fmt.Println("a <= b : ", c)
}
