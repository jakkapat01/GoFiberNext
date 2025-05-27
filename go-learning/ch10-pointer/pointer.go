package main

import "fmt"

func SamplePointer1() {
	//บ้านเลขที่ 10 (เก็บค่าจริง)
	value := 10

	//ป้ายเลขที่บ้าน = ตัวแปร pointer ที่เก็บ address ของบ้าน
	pointer := &value

	//แสดง address ของบ้าน
	fmt.Println("Address of value:", pointer) // 0xc0000120b8 (example)
	//แสดงค่าที่อยู่ในบ้าน
	fmt.Println("Value:", value) // 10
	//แสดงค่าที่อยู่ในบ้านผ่าน pointer
	fmt.Println("Value through pointer:", *pointer) // 10

	//เปลี่ยนค่าในบ้านผ่าน pointer
	*pointer = 20
	fmt.Println("Value:", value) // 20
	fmt.Println("----------------------")
	//การสร้าง pointer กับ slice
	// สร้าง slice ที่มีค่าเริ่มต้น
	slice := []int{1, 2, 3}
	// สร้าง pointer ที่ชี้ไปยัง slice
	pointerSlice := &slice
	// แสดง address ของ slice
	fmt.Println("Address of slice:", pointerSlice)
	fmt.Println("Slice value:", *pointerSlice) // [1 2 3]

}
func changeAge(age *int) { // ฟังก์ชันที่รับ pointer ของตัวแปร age
	*age += 1                                 // เพิ่มค่า age ที่ชี้ไปยัง pointer
	fmt.Println("อายุหลังเปลี่ยนแปลง:", *age) // แสดงค่าอายุที่เปลี่ยนแปลงแล้ว
}

// กรณีใช้ pointer
// กำหนด struct สำหรับห้อง
type Room struct {
	RoomNumber int
	BedColor   string
}

// function ที่ใช้ pointer เพื่อเปลี่ยนแปลงสีเตียง
func ChangeBedColorWithPointer(keyToRoom *Room, newColor string) {
	keyToRoom.BedColor = newColor // เปลี่ยนแปลงสีเตียงโดยใช้ pointer
}
func SamplePointer2() {
	room101 := Room{
		RoomNumber: 101,
		BedColor:   "เหลือง",
	}
	fmt.Println("ก่อนเปลี่ยนสี")
	fmt.Println("หมายเลขห้อง:", room101.RoomNumber)
	fmt.Println("สีเตียง:", room101.BedColor)

	// เรียกใช้ฟังก์ชันเพื่อเปลี่ยนแปลงสีเตียง
	ChangeBedColorWithPointer(&room101, "เขียว") // ส่ง pointer ของ room101
	fmt.Println("หลังเปลี่ยนสี")
	fmt.Println("หมายเลขห้อง:", room101.RoomNumber)
	fmt.Println("สีเตียง:", room101.BedColor) // แสดงสีเตียงที่เปลี่ยนแปลงแล้ว
}
