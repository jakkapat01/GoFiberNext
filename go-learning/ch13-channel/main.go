package main

import (
	"fmt"
	"time"
)

func main() {
	// สร้าง channel สำหรับส่งขข้อมูลประเภท string
	ch := make(chan string)
	// เรียกใช้ Goroutine เพื่อส่งข้อมูลไปยัง channel
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from Goroutine!"
	}()
	// รอรับข้อมูลจาก channel
	msg := <-ch
	fmt.Println(msg)
}
