package main

import (
	"log"
	"os"
	"path/filepath"

	"vcm-medical-platform/internal/api"
	"vcm-medical-platform/internal/config"
	"vcm-medical-platform/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Serve static files from frontend build
	publicPath := filepath.Join("frontend", "dist")
	if _, err := os.Stat(publicPath); err == nil {
		app.Static("/", publicPath)
		// SPA fallback
		app.Get("/*", func(c *fiber.Ctx) error {
			return c.SendFile(filepath.Join(publicPath, "index.html"))
		})
	}

	// Setup API routes
	api.SetupRoutes(app, db)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Printf("🌐 Access at: http://localhost:%s", port)
	log.Printf("💊 API Health: http://localhost:%s/health", port)
	log.Fatal(app.Listen(":" + port))
}
