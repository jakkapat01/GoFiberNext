package main

import (
	"fmt"
)

var xmlglobal string = "xmlglobal" // global variable

func main() {
	fmt.Println(xmlglobal) // global variable
	xml := "xml"           // local variable
	fmt.Println(xml)       // local variable

	//แบบหลายตัวแปร
	var (
		num   float32
		count int    = 10
		name  string = "john"
	)

	num = 10.5
	fmt.Println(num)   // local variable
	fmt.Println(count) // local variable
	fmt.Println(name)  // local variable

	// แบบ multiple short variable declaration
	var x, y, z = 1_000_000, 20, 3_0 + 5
	var p, q, r = "Hello", "World", "go"
	fmt.Println(x, y, z) // local variable
	fmt.Println(p, q, r) // local variable

	// ประกาศค่าคงที่
	// const
	const pi float32 = 3.14
	const vat = 7
	const name1, name2 = "poa", "jakkapat"
	fmt.Println(pi)    // local variable
	fmt.Println(vat)   // local variable
	fmt.Println(name1) // local variable

}
