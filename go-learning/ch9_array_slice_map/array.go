package main

import (
	"fmt"
)

func SampleArray1() {

	// ประกาศตัวแปร array ขนาด 3 ช่อง เก็บข้อมูลประเภท string
	var fruits [3]string

	// กำหนดค่าให้กับ array fruits
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	// fruits[3] = "Grapes" // Error: index out of range

	// แสดงผลข้อมูลใน array fruits
	fmt.Println("Fruits:", fruits) // [Apple Banana Orange]

	index := 0

	// เข้าถึงสมาชิกใน Array โดยใช้ index
	fmt.Println("Fruit at index 1:", fruits[1])     // Banana
	fmt.Println("Fruit at index 2:", fruits[1+1])   // Orange
	fmt.Println("Fruit at index 1:", fruits[index]) // Apple

	// นับจำนวนสมาชิกใน Array
	fmt.Println("Number of fruits:", len(fruits)) // 3

	fmt.Println("-------------------")

	// แสดงผลข้อมูลใน array fruits โดยใช้ loop
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index", i, ":", fruits[i])
	}

	fmt.Println("-------------------")

	// แสดงผลข้อมูลใน array fruits โดยใช้ range
	for i, fruit := range fruits {
		fmt.Println("Fruit at index", i, ":", fruit)
	}

}

// เรียนรู้การใช้ array เพิ่มเติม การสร้าง Array และกำหนดสมาชิก

func SampleArray2() {

	// ประกาศตัวแปร array ขนาด 5 ช่อง เก็บข้อมูลประเภท string
	var carsbrands = [5]string{"Toyota", "Honda", "Nissan", "Mazda", "Subaru"}

	// แสดงผลข้อมูลใน array carsbrands
	fmt.Println("Cars brands:", carsbrands) // [Toyota Honda Nissan Mazda Subaru]

	// กรณีกำหนดสมาชิกไม่ครบ
	var laptopbrands = [5]string{"Apple", "Dell", "HP"}

	// แสดงผลข้อมูลใน array laptopbrands
	fmt.Println("Laptop brands:", laptopbrands) // [Apple Dell HP ] // ช่องที่เหลือจะถูกกำหนดเป็นค่าว่าง (empty string)

	// กรณี Array ขนาด 10 ช่อง แต่กำหนดสมาชิก 5 ช่อง
	var amount = [10]int{1, 2, 3, 4, 5}

	// แสดงผลข้อมูลใน array amount
	fmt.Println("Amount:", amount) // [1 2 3 4 5 0 0 0 0 0] // ช่องที่เหลือจะถูกกำหนดเป็นค่า 0 (zero value)

	// นอกจากนั้นยังสามารถกำหนดสมาชิกใน array โดยใช้ index ได้
	var numbers = [5]int{1: 10, 3: 20}

	// แสดงผลข้อมูลใน array numbers
	fmt.Println("Numbers:", numbers) // [0 10 0 20 0] // ช่องที่เหลือจะถูกกำหนดเป็นค่า 0 (zero value)

	// สามารถสร้าง array ที่มีขนาดไม่แน่นอนได้ โดยใช้ ... (ellipsis)
	var array1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// แสดงผลข้อมูลใน array array1
	fmt.Println("Array1:", array1) // [1 2 3 4 5 6 7 8 9 10]

	// เพิ่มสมาชิกใน Array โดยการใช้ append ไม่ได้
	// array1 = append(array1, 16) // Error: cannot use append(array1, 16) (value of type []int) as [15]int value in assignment
}
