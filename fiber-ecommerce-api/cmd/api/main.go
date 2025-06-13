package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/http/routes"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/adapters/persistence/repositories"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/config"
	"github.com/jakkapat01/fiber-ecommerce-api/internal/core/services"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// Setup database connection
	db := config.SetupDatabase(cfg)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
