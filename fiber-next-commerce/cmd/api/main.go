package main

import (
	"fmt"
	"log"

	"github.com/jakkapat01/fiber-next-commerce/config"
	"github.com/jakkapat01/fiber-next-commerce/internal/adapters/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize a new Fiber instance
	app := fiber.New()

	// Connect to the database
	db.Connect(cfg)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from fiber Next Commerce!"})
	})
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.AppPort)))
	log.Println("Server is running on port", cfg.AppPort)

}
