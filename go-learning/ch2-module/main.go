package main

import (
	"fmt"
	_ "math"

	"github.com/google/uuid"
	"github.com/jakkapat01/GoFiberNext/greeting"
)

func main() {
	fmt.Println("Hello, go")
	// การสร้าง uuid
	// สร้่างตัวแปร uuid
	var id string = uuid.New().String()
	fmt.Println("UUID: ", id)

	greeting.SayGreeting()
	greeting.SayMeeting()
}
