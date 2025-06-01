package main

import (
	"fiber-restapi/controllers"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// สร้าง service และ controller สำหรับผู้ใช้
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// สร้าง route สำหรับลงทะเบียนผู้ใช้
	app.Post("/register", userController.Register)
	// สร้าง route สำหรับเข้าสู่ระบบ
	app.Post("/login", userController.Login)

	app.Listen(":3000")
}
