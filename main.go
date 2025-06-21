package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed frontend/dist/*
var embedDirStatic embed.FS

func main() {
	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "VCM Medical Platform",
			"version": "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api")
	
	// Test API route
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is working!",
			"time":    fmt.Sprintf("%v", os.Getenv("PORT")),
		})
	})

	// Serve static files from embedded filesystem
	staticFiles, err := fs.Sub(embedDirStatic, "frontend/dist")
	if err != nil {
		log.Printf("Error creating static file system: %v", err)
		// Fallback: serve a simple HTML page
		app.Get("/*", func(c *fiber.Ctx) error {
			return c.Type("html").SendString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-50">
    <div class="min-h-screen flex items-center justify-center">
        <div class="text-center">
            <div class="w-16 h-16 bg-gradient-to-br from-blue-600 to-blue-700 rounded-2xl flex items-center justify-center mx-auto mb-4">
                <span class="text-white font-bold text-2xl">V</span>
            </div>
            <h1 class="text-4xl font-bold text-gray-900 mb-4">VCM Medical Platform</h1>
            <p class="text-xl text-gray-600 mb-8">Advanced Medical Treatments for Complex Conditions</p>
            <div class="space-y-4">
                <a href="/health" class="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
                    Check Health Status
                </a>
                <br>
                <a href="/api/test" class="inline-block bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700 transition-colors">
                    Test API
                </a>
            </div>
            <div class="mt-8 text-sm text-gray-500">
                <p>Frontend build files not found. Building frontend...</p>
            </div>
        </div>
    </div>
</body>
</html>
			`)
		})
	} else {
		// Serve static files with SPA fallback
		app.Use("/", filesystem.New(filesystem.Config{
			Root: http.FS(staticFiles),
		}))

		// SPA fallback - serve index.html for all routes that don't match static files
		app.Get("/*", func(c *fiber.Ctx) error {
			path := c.Path()
			
			// Don't serve index.html for API routes
			if strings.HasPrefix(path, "/api") {
				return c.Status(404).JSON(fiber.Map{"error": "API route not found"})
			}

			// Serve index.html for all other routes (SPA routing)
			return c.SendFile("frontend/dist/index.html")
		})
	}

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
