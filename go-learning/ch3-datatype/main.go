package main

import (
	"fmt"
	"html"
)

func main() {
	// Data Type in go
	//มี 2 ประเภทหลักๆ
	// 1. Basic Type (int, float, string, bool)
	// 2. Composite Type (array, slice, map, struct, interface)

	// Basic Type
	//numberic type
	var a int8 = 10   //range -128 to 127
	var b int16 = 20  //range -32768 to 32767
	var c int32 = 30  //range -2147483648 to 2147483647
	var d int64 = 40  //range -9223372036854775808 to 9223372036854775807
	var e int = 50    //range -2147483648 to 2147483647 (32 bit)
	var f uint8 = 60  //range 0 to 255
	var g uint16 = 70 //range 0 to 65535
	var h uint32 = 80 //range 0 to 4294967295
	var i uint64 = 90 //range 0 to 18446744073709551615
	var j uint = 100  //range 0 to 4294967295 (32 bit)

	//float type
	var k float32 = 1.1 //range -3.402823466E+38 to 3.402823466E+38
	var l float64 = 2.2 //range -1.7976931348623157E+308 to 1.7976931348623157E+308

	//string type
	var m string = "Hello, World!"

	//boolean type
	var n bool = true

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)
	fmt.Println("m = ", m)
	fmt.Println("n = ", n)

	//String data type
	//Interpret string
	var str1 string = "Welcome to Go!/nfor the \"first\" \ttime"
	//Raw string
	var str2 string = `Welcome to Go!/nfor the first time
	this is a raw string
	you can write anything here
	`
	/*
		Escape character
		\n = new line
		\t = tab
		\\ = backslash
		\" = double quote
		\' = single quote
		\r = carriage return
		\b = backspace
	*/

	fmt.Println("str1 = ", str1)
	fmt.Println("str2 = ", str2)

	var str string = `<a href="https://www.google.com">Google</a>`
	fmt.Println("str = ", str)
	//Escape html
	var str_escape string = html.EscapeString(str)
	fmt.Println("str_escape = ", str_escape)
	//Unescape html
	var str_unescape string = html.UnescapeString(str_escape)
	fmt.Println("str_unescape = ", str_unescape)
}
