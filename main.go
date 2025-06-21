package main

import (
        "log"
        "os"

        "github.com/gofiber/fiber/v2"
        "github.com/gofiber/fiber/v2/middleware/cors"
        "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
        app := fiber.New()

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
                        "description": "Advanced Medical Treatment Platform with 95% Efficacy",
                        "status":      "running",
                        "features": []string{
                                "Multi-user authentication",
                                "Medical assessment forms", 
                                "Appointment booking",
                                "Photo progress tracking",
                                "Treatment protocols",
                                "Doctor consultations",
                        },
                })
        })

        // Serve static files - simple approach
        app.Static("/", "./frontend/dist")

        // Fallback for SPA routing
        app.Get("*", func(c *fiber.Ctx) error {
                return c.SendFile("./frontend/dist/index.html")
        })

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("Server starting on port %s", port)
        log.Fatal(app.Listen(":" + port))
}
