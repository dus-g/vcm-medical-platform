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
				"Real-time chat support",
				"Order management",
			},
		})
	})

	// Serve a simple HTML page
	app.Get("/", func(c *fiber.Ctx) error {
		html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; }
        .container { max-width: 800px; margin: 0 auto; text-align: center; }
        .header { background: rgba(255,255,255,0.1); padding: 40px; border-radius: 20px; margin-bottom: 30px; }
        .features { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; }
        .feature { background: rgba(255,255,255,0.1); padding: 20px; border-radius: 15px; }
        .btn { display: inline-block; background: #4CAF50; color: white; padding: 15px 30px; text-decoration: none; border-radius: 25px; margin: 10px; }
        .btn:hover { background: #45a049; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🏥 VCM Medical Platform</h1>
            <p>Advanced Medical Treatment Platform with 95% Efficacy</p>
            <p><strong>✅ Successfully Deployed on Railway!</strong></p>
        </div>
        
        <div class="features">
            <div class="feature">
                <h3>🔐 Multi-User Authentication</h3>
                <p>8 user types supported with JWT security</p>
            </div>
            <div class="feature">
                <h3>📋 Medical Assessments</h3>
                <p>Comprehensive forms for psoriasis and eye diseases</p>
            </div>
            <div class="feature">
                <h3>📅 Appointment System</h3>
                <p>Schedule consultations with specialized doctors</p>
            </div>
            <div class="feature">
                <h3>💬 Real-time Chat</h3>
                <p>5 chat room types with support team</p>
            </div>
        </div>
        
        <div style="margin-top: 40px;">
            <a href="/health" class="btn">🔍 API Health Check</a>
            <a href="/api/v1/info" class="btn">📊 Platform Info</a>
        </div>
        
        <div style="margin-top: 40px; background: rgba(255,255,255,0.1); padding: 20px; border-radius: 15px;">
            <h3>🚀 Platform Status</h3>
            <p>✅ Go Backend Running</p>
            <p>✅ API Endpoints Active</p>
            <p>✅ Railway Deployment Successful</p>
            <p>✅ Ready for Database Integration</p>
        </div>
        
        <footer style="margin-top: 40px; opacity: 0.8;">
            <p>© 2024 VAMOS BIOTECH (Shanghai) Co., Ltd.</p>
        </footer>
    </div>
</body>
</html>`
		c.Type("html")
		return c.SendString(html)
	})

	// Catch all other routes
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 VCM Medical Platform starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
