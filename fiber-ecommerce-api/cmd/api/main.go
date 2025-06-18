// @title E-commerce API
// @version 1.0.0
// @description API สำหรับระบบ E-commerce ที่พัฒนาด้วย Go Fiber
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT token สำหรับการยืนยันตัวตน ให้ใส่ token ในรูปแบบ: Bearer <token>
package main

import (
	"log"

	_ "github.com/jakkapat01/fiber-ecommerce-api/docs"

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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

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
