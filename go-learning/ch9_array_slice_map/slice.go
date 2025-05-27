package main

import "fmt"

func SampleSlice1() {

	// ประกาศตัวแปรแบบ slice
	var fruits []string

	// กำหนดค่าให้กับ slice fruits
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Orange")

	// แสดงผลข้อมูลใน slice fruits
	fmt.Println("Fruits:", fruits) // [Apple Banana Orange]

	fmt.Println("----------------------")

	// เข้าถึงสมาชิกใน Slice โดยใช้ index
	fmt.Println("Fruit at index 1:", fruits[1]) // Banana

	fmt.Println("----------------------")

	// เพิ่มสมาชิกใน Slice
	fruits = append(fruits, "Grapes")
	fruits = append(fruits, "Mango")

	// แสดงผลข้อมูลใน slice fruits
	fmt.Println("Fruits:", fruits) // [Apple Banana Orange Grapes Mango]

	fmt.Println("----------------------")

	// สร้าง Slice โดยใช้ literal syntax
	numbers := []int{1, 2, 3, 4, 5}

	// แสดงผลข้อมูลใน slice numbers
	fmt.Println("Numbers:", numbers) // [1 2 3 4 5]

	// เข้าถึงสมาชิกใน Slice โดยใช้ index
	fmt.Println("Number at index 2:", numbers[2]) // 3

	// เพิ่มสมาชิกใน Slice
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)

	// แสดงผลข้อมูลใน slice numbers
	fmt.Println("Numbers:", numbers) // [1 2 3 4 5 6 7]

	fmt.Println("----------------------")

	// ตัด slice โดยใช้ slice operator
	subSlice1 := numbers[1:4] // ตัด slice ตั้งแต่ index 1 ถึง index 3
	subSlice2 := numbers[:3]  // ตัด slice ตั้งแต่ index 0 ถึง index 2
	subSlice3 := numbers[2:]  // ตัด slice ตั้งแต่ index 2 ถึง index 6

	// แสดงผลข้อมูลใน slice subSlice1
	fmt.Println("Sub slice:", subSlice1) // [2 3 4]
	// แสดง length ของ subSlice1
	fmt.Println("Length of sub slice:", len(subSlice1)) // 3
	// แสดง capacity ของ subSlice1
	fmt.Println("Capacity of sub slice:", cap(subSlice1)) // 5

	fmt.Println("----------------------")

	// แสดงผลข้อมูลใน slice subSlice2
	fmt.Println("Sub slice:", subSlice2) // [1 2 3]
	// แสดง length ของ subSlice2
	fmt.Println("Length of sub slice:", len(subSlice2)) // 3
	// แสดง capacity ของ subSlice2
	fmt.Println("Capacity of sub slice:", cap(subSlice2)) // 5

	fmt.Println("----------------------")

	// แสดงผลข้อมูลใน slice subSlice3
	fmt.Println("Sub slice:", subSlice3) // [3 4 5 6 7]
	// แสดง length ของ subSlice3
	fmt.Println("Length of sub slice:", len(subSlice3)) // 5
	// แสดง capacity ของ subSlice3
	fmt.Println("Capacity of sub slice:", cap(subSlice3)) // 5

	// cap เป็นค่า 10 เนื่องจาก
	// เวลาเรา ประกาศค่าในตอนต้น
	// numbers := []int{1, 2, 3, 4, 5}
	// จะมีค่า len ที่ 5 และ cap ที่ 5
	// พอเรา append ค่าเข้า slice มันทำให้ ค่าที่เรา append เกินค่า cap ตอนเราประกาศค่าตอนแรกทำให้ go ต้องขยาย memory เพิ่มอีก 1 เท่าจากเดิม ( 5*2 = 10) ครับ
}
