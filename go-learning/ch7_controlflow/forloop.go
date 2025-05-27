package main

import (
	"fmt"
)

// forloop
func Forloop1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello World", i)
	}
}

// for like a while
func Forloop2() {
	i := 1
	for i <= 5 {
		fmt.Println("Hello World", i)
		i++
	}

}

// for infinite loop
func Forloop3() {
	for {
		fmt.Println("Hello World")
	}
}
