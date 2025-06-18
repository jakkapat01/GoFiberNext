package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/http/middleware"
)

// Setup Routes
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	// สร้าง handler สำหรับผู้ดูแลระบบ
	adminHandler := handlers.NewAdminHandler()
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	// API Routes
	api := app.Group("/api")
	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected Routes
	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())        // ใช้ middleware สำหรับตรวจสอบสิทธิ์
	user.Get("/profile", authHandler.GetProfile) // ดึงข้อมูลโปรไฟล์ของผู้ใช้

	// Admin only Routes
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())             // ใช้ middleware สำหรับตรวจสอบสิทธิ์
	admin.Use(middleware.RequireRole("admin"))         // ตรวจสอบว่าเป็นผู้ดูแลระบบ
	admin.Get("/dashboard", adminHandler.GetDashboard) // ดึงข้อมูลแดชบอร์ดของผู้ดูแลระบบ)
	admin.Post("/register", authHandler.AdminRegister) // ลงทะเบียนผู้ใช้ใหม่โดยผู้ดูแลระบบ

}
