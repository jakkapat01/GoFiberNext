package main

import (
	"fmt"
)

func MathOperator() {
	// การใช้ operator
	// การใช้ operator + - * / %
	var a int = 10
	var b int = 20
	var c int = 30
	var d int = 40
	var e int = 50
	var f int = 60
	var g int = 70
	var h int = 80
	var i int = 90
	var j int = 100

	fmt.Println("a + b =", a+b)
	fmt.Println("c - d =", c-d)
	fmt.Println("e * f =", e*f)
	fmt.Println("g / h =", g/h)
	fmt.Println("i % j =", i%j)
	// การใช้ operator ++ --
	a++
	b--
	fmt.Println("a++ =", a)
	fmt.Println("b-- =", b)
}

// 	fmt.Println("str1 = ", str1)
