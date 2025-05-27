package main

import (
	"fmt"
)

func AssignmentOperator() {
	// =, +=, -=, *=, /=, %=
	var a int = 10
	var b int = 20
	var c int

	c = a + b // 30
	fmt.Println("a + b : ", c)

	c += a // 40
	fmt.Println("c += a : ", c)

	c -= a // 30
	fmt.Println("c -= a : ", c)

	c *= a // 300
	fmt.Println("c *= a : ", c)

	c /= a // 30
	fmt.Println("c /= a : ", c)

	c %= a // 0
	fmt.Println("c %= a : ", c)

	// การใช้ operator &=, |=, ^=, <<=, >>=
	var d int = 10
	var e int = 20
	var f int
	f = d & e // 0
	fmt.Println("d & e : ", f)
	f = d | e // 30
	fmt.Println("d | e : ", f)
	f = d ^ e // 30
	fmt.Println("d ^ e : ", f)
	f = d << 2 // 40
	fmt.Println("d << 2 : ", f)
	f = d >> 2 // 2
	fmt.Println("d >> 2 : ", f)

}
