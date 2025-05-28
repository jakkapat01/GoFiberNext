package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//สร้างแอปพลิเคชัน Fiber
	app := fiber.New()

	//กำหนดเส้นทางสำหรับการตอบสนอง
	app.Get("/", func(c *fiber.Ctx) error {
		//ส่งข้อความ "Hello, World!" กลับไปยังผู้ใช้
		return c.SendString("Hello, World!")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About page")
	})
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("Contact page")
	})

	//เริ่มเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
