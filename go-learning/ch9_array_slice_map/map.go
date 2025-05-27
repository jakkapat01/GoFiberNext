package main

import (
	"fmt"
)

func SampleMap1() {
	//ประกาศ var ตัวแปร map[string]int
	// ประกาศ map ที่เก็บ key เป็น string และ value เป็น int
	scores := make(map[string]int)
	// เพิ่มข้อมูลลงใน map
	scores["Alice"] = 90
	scores["Bob"] = 85
	scores["Charlie"] = 95
	fmt.Println("Scores:", scores) //Scores: map[Alice:90 Bob:85 Charlie:95]

	fmt.Println("--------------------")
	// การเข้าถึงค่าใน map โดยใช้ key
	fmt.Println("Alice's score:", scores["Alice"]) // Alice's score: 90

	// การตรวจสอบว่ามี key อยู่ใน map หรือไม่
	if score, exists := scores["Bob"]; exists {
		fmt.Println("Bob's score:", score) // Bob's score: 85
	} else {
		fmt.Println("Bob's score not found")
	}
	fmt.Println("--------------------")
	// การลบ key-value pair ออกจาก map
	delete(scores, "Charlie")
	fmt.Println("Scores after deleting Charlie:", scores) // Scores after deleting Charlie: map[Alice:90 Bob:85]
	fmt.Println("--------------------")
	// การวนลูปผ่าน map
	for name, score := range scores {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}

	// การสร้่าง map โดยใช้ literal syntax
	grades := map[string]int{
		"Math":    95,
		"Science": 90,
		"English": 85,
		"History": 88,
	}
	fmt.Println("Grades;", grades) // Grades: map[English:85 History:88 Math:95 Science:90]
}
