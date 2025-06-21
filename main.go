package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "VCM Medical Platform",
			"version":   "2.0.0",
			"time":      time.Now(),
			"platform":  "Railway",
		})
	})

	app.Get("/api/v1/countries", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"countries": []fiber.Map{
				{"id": 1, "name": "United States", "code": "US"},
				{"id": 2, "name": "Canada", "code": "CA"},
				{"id": 3, "name": "United Kingdom", "code": "UK"},
				{"id": 4, "name": "Germany", "code": "DE"},
				{"id": 5, "name": "France", "code": "FR"},
				{"id": 6, "name": "Japan", "code": "JP"},
				{"id": 7, "name": "Australia", "code": "AU"},
				{"id": 8, "name": "India", "code": "IN"},
			},
		})
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(`
<!DOCTYPE html>
<html>
<head>
    <title>VCM Medical Platform</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            text-align: center;
            background: rgba(255,255,255,0.1);
            padding: 2rem;
            border-radius: 20px;
            backdrop-filter: blur(10px);
            max-width: 500px;
            margin: 1rem;
        }
        h1 { font-size: 2.5rem; margin-bottom: 1rem; }
        .success { 
            background: rgba(76, 175, 80, 0.3);
            padding: 1rem 2rem;
            border-radius: 25px;
            margin: 1rem 0;
            border: 2px solid rgba(76, 175, 80, 0.5);
        }
        .endpoint {
            background: rgba(0,0,0,0.2);
            padding: 0.5rem 1rem;
            margin: 0.5rem 0;
            border-radius: 8px;
            font-family: monospace;
        }
        a { color: #87CEEB; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🏥 VCM Medical Platform</h1>
        <p>Advanced Medical Treatment Platform</p>
        
        <div class="success">
            ✅ Successfully Deployed on Railway!
        </div>
        
        <div style="background: rgba(255,255,255,0.1); padding: 1rem; border-radius: 10px; margin: 1rem 0;">
            <h3>🔗 Test Endpoints:</h3>
            <div class="endpoint"><a href="/health">/health</a> - API Health Check</div>
            <div class="endpoint"><a href="/api/v1/countries">/api/v1/countries</a> - Countries Data</div>
        </div>
        
        <p style="font-size: 0.9rem; opacity: 0.8; margin-top: 2rem;">
            🎉 Your VCM Medical Platform is now live on Railway!<br>
            Ready for frontend integration and database connection.
        </p>
    </div>
</body>
</html>`)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
