package main

import (
	"fmt"
)

func main() {
	SamplePointer1()
	fmt.Println("====================================")
	age := 30
	fmt.Println("อายุเดิม:", age)
	changeAge(&age)
	fmt.Println("-------------------------------")
	SamplePointer2()
}
