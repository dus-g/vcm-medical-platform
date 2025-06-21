package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed frontend/dist
var embedDirStatic embed.FS

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "VCM Medical Platform is running",
		})
	})

	// API routes
	app.Get("/api/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is working!",
		})
	})

	// Serve static files
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix: "frontend/dist",
		Browse: true,
	}))

	// SPA fallback
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("frontend/dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
