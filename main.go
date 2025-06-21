package main

import (
        "log"
        "os"
        "path/filepath"

        "github.com/gofiber/fiber/v2"
        "github.com/gofiber/fiber/v2/middleware/cors"
        "github.com/gofiber/fiber/v2/middleware/logger"
        "github.com/gofiber/fiber/v2/middleware/filesystem"
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

        // Serve static files from frontend/dist
        distPath := filepath.Join("frontend", "dist")
        if _, err := os.Stat(distPath); err == nil {
                log.Printf("Serving frontend from: %s", distPath)
                app.Use("/", filesystem.New(filesystem.Config{
                        Root: distPath,
                        Browse: false,
                        Index: "index.html",
                }))
                
                // Handle SPA routing - serve index.html for all routes not matched above
                app.Use("*", func(c *fiber.Ctx) error {
                        return c.SendFile(filepath.Join(distPath, "index.html"))
                })
        } else {
                log.Printf("Frontend dist folder not found at %s, serving basic HTML", distPath)
                
                // Fallback HTML if frontend dist is not available
                app.Get("*", func(c *fiber.Ctx) error {
                        return c.Type("html").SendString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; margin: 0; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); min-height: 100vh; display: flex; align-items: center; justify-content: center; }
        .container { background: white; padding: 40px; border-radius: 20px; box-shadow: 0 20px 40px rgba(0,0,0,0.1); text-align: center; max-width: 600px; }
        .logo { background: linear-gradient(135deg, #667eea, #764ba2); color: white; width: 80px; height: 80px; border-radius: 50%; display: flex; align-items: center; justify-content: center; margin: 0 auto 20px; font-size: 24px; font-weight: bold; }
        h1 { color: #333; margin-bottom: 10px; }
        .subtitle { color: #666; margin-bottom: 30px; }
        .features { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 20px; margin: 30px 0; }
        .feature { background: #f8f9fa; padding: 20px; border-radius: 15px; }
        .btn { background: linear-gradient(135deg, #667eea, #764ba2); color: white; padding: 15px 30px; border: none; border-radius: 25px; font-size: 16px; font-weight: 600; text-decoration: none; display: inline-block; margin: 10px; transition: transform 0.3s; }
        .btn:hover { transform: translateY(-2px); }
        .status { background: #e7f5e7; color: #2d5a2d; padding: 10px 20px; border-radius: 20px; display: inline-block; margin: 20px 0; }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">VCM</div>
        <h1>VCM Medical Platform</h1>
        <p class="subtitle">Advanced Medical Treatment Platform</p>
        <div class="status">✅ API Running Successfully</div>
        
        <div class="features">
            <div class="feature">
                <h3>🏥 Medical Excellence</h3>
                <p>95% treatment efficacy with breakthrough therapies</p>
            </div>
            <div class="feature">
                <h3>🔬 World's First</h3>
                <p>Clinical trials for antibiotic-resistant infections</p>
            </div>
            <div class="feature">
                <h3>🌍 Global Access</h3>
                <p>24/7 platform with Shanghai headquarters</p>
            </div>
            <div class="feature">
                <h3>⚡ Advanced Tech</h3>
                <p>Life-cell based therapies and assessment systems</p>
            </div>
        </div>
        
        <a href="/api/v1/info" class="btn">View API Info</a>
        <a href="/health" class="btn">Health Check</a>
        
        <p style="margin-top: 30px; color: #888; font-size: 14px;">
            VAMOS BIOTECH (Shanghai) Co., Ltd.<br>
            Bio-pharmaceutical innovation platform
        </p>
    </div>
</body>
</html>`)
                })
        }

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("Server starting on port %s", port)
        log.Fatal(app.Listen(":" + port))
}
