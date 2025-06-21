package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "VCM Medical Platform API",
			"version":   "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api/v1")
	
	api.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":        "VCM Medical Platform",
			"description": "Advanced Medical Treatment Platform",
			"status":      "running",
		})
	})

	// Serve frontend files
	app.Static("/", "./frontend/dist")
	
	// SPA fallback
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/dist/index.html")
	})

	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
