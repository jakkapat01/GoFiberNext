package main

import (
	"fmt"
	"strconv"
)

func TypeConvert1() {
	var i int = 42
	var f float64 = float64(i) // Convert int to float64
	fmt.Printf("i = %d, f = %f\n", i, f)

}
func TypeConvert2() {
	var f float64 = 42.56
	var i int = int(f)
	fmt.Printf("f = %f, i = %d\n", f, i)
}
func TypeConvert3() {
	var i int = 42
	var s string = strconv.Itoa(i) // Convert int to string (integer to ASCII)
	fmt.Printf("i = %d, s = %s\n", i, s)
}
func TypeConvert4() {
	var s string = "42"
	var i int
	var err error
	i, err = strconv.Atoi(s) // Convert string to int (ASCII to integer)
	//ถ้าแปลงไม่สำเร็จ err จะไม่เป็น nil
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("s = %s, i = %d\n", s, i)

}

// แปลงจาก boolean เป็น int เพราะ go ไม่มีการแปลงอัตโนมัติ
func TypeConvert5() {
	var b bool = true
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Printf("b = %t, i = %d\n", b, i)
}
