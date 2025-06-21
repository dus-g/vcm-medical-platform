package main

import (
        "log"
        "os"
        "path/filepath"

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

        // Check if frontend dist exists
        distPath := "./frontend/dist"
        indexPath := filepath.Join(distPath, "index.html")
        
        if _, err := os.Stat(indexPath); err == nil {
                log.Printf("✅ Frontend dist found, serving React app from: %s", distPath)
                
                // Serve static assets
                app.Static("/assets", filepath.Join(distPath, "assets"))
                
                // Serve the React app for all other routes
                app.Get("*", func(c *fiber.Ctx) error {
                        return c.SendFile(indexPath)
                })
        } else {
                log.Printf("❌ Frontend dist not found at %s, serving fallback HTML", indexPath)
                
                // Fallback HTML with modern design
                app.Get("*", func(c *fiber.Ctx) error {
                        return c.Type("html").SendString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap" rel="stylesheet">
</head>
<body class="font-sans bg-gray-50">
    <!-- Header -->
    <header class="fixed w-full top-0 z-50 bg-white/95 backdrop-blur-md border-b border-gray-100 shadow-sm">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex items-center justify-between h-20">
                <div class="flex items-center">
                    <div class="h-16 w-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                        <span class="text-white font-bold text-xl">VCM</span>
                    </div>
                    <span class="ml-3 text-xl font-bold text-gray-900">VCM Medical</span>
                </div>
                <nav class="hidden md:flex space-x-8">
                    <a href="#" class="text-gray-700 hover:text-blue-600 font-medium">Home</a>
                    <a href="#" class="text-gray-700 hover:text-blue-600 font-medium">Therapies</a>
                    <a href="#" class="text-gray-700 hover:text-blue-600 font-medium">About</a>
                    <a href="#" class="text-gray-700 hover:text-blue-600 font-medium">Contact</a>
                </nav>
            </div>
        </div>
    </header>

    <!-- Main Content -->
    <main class="pt-20">
        <!-- Hero Section -->
        <section class="pt-16 pb-16 px-6 bg-gradient-to-br from-blue-50 via-blue-50 to-cyan-50">
            <div class="max-w-7xl mx-auto text-center">
                <div class="inline-flex items-center px-4 py-2 bg-blue-100 border border-blue-200 rounded-full text-blue-700 text-sm font-medium mb-6">
                    <div class="w-2 h-2 bg-blue-500 rounded-full mr-2 animate-pulse"></div>
                    VAMOS BIOTECH - Bio-Pharmaceutical Innovation
                </div>
                
                <h1 class="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
                    Advanced Medical
                    <span class="block text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-blue-600">
                        Treatment Platform
                    </span>
                </h1>
                
                <p class="text-xl text-gray-600 mb-8 max-w-4xl mx-auto">
                    Breakthrough life-cell based therapies for cancer, viral infections, autoimmune disorders, 
                    and antibiotic-resistant bacterial infections with proven 95% efficacy rates.
                </p>
                
                <div class="flex flex-col sm:flex-row gap-4 justify-center mb-12">
                    <button class="px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg shadow-lg transition-colors">
                        Start Treatment
                    </button>
                    <button class="px-8 py-3 bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium rounded-lg transition-colors">
                        Explore Therapies
                    </button>
                </div>

                <!-- Company Highlights -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 max-w-6xl mx-auto">
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🏆</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">95% Treatment Efficacy</h3>
                        <p class="text-sm text-gray-600">Breakthrough results in melanoma and cancer treatment</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🔬</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">World's First</h3>
                        <p class="text-sm text-gray-600">Clinical trials for antibiotic-resistant infections</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🌍</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">24/7 Platform Access</h3>
                        <p class="text-sm text-gray-600">Global operations with Shanghai headquarters</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">⚡</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Advanced Life-Cell Therapies</h3>
                        <p class="text-sm text-gray-600">Cutting-edge medical and naturopathic treatments</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Treatment Categories -->
        <section class="py-16 px-6 bg-white">
            <div class="max-w-7xl mx-auto">
                <div class="text-center mb-12">
                    <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
                        Medical Conditions We Treat
                    </h2>
                    <p class="text-xl text-gray-600 max-w-3xl mx-auto">
                        Comprehensive treatment protocols across major medical specialties
                    </p>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">🛡️</span>
                            <h3 class="text-xl font-bold text-gray-900">Cancers</h3>
                        </div>
                        <ul class="space-y-2 text-sm text-gray-700">
                            <li>• B-cell leukemia, lymphoma (CAR-T)</li>
                            <li>• Melanoma therapy</li>
                            <li>• Colorectal, pancreas cancer</li>
                        </ul>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">❤️</span>
                            <h3 class="text-xl font-bold text-gray-900">Autoimmune</h3>
                        </div>
                        <ul class="space-y-2 text-sm text-gray-700">
                            <li>• Psoriasis Vulgaris</li>
                            <li>• Rheumatoid arthritis</li>
                            <li>• Lupus, Hashimoto's disease</li>
                        </ul>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">👁️</span>
                            <h3 class="text-xl font-bold text-gray-900">Eye Diseases</h3>
                        </div>
                        <ul class="space-y-2 text-sm text-gray-700">
                            <li>• Macular degeneration</li>
                            <li>• Glaucoma</li>
                            <li>• Lazy eye (Amblyopia)</li>
                        </ul>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">🫁</span>
                            <h3 class="text-xl font-bold text-gray-900">Respiratory</h3>
                        </div>
                        <ul class="space-y-2 text-sm text-gray-700">
                            <li>• Tuberculosis</li>
                            <li>• Pneumonia & Bronchitis</li>
                            <li>• Gastritis & H. Pylori</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-900 text-white py-12">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
                <div>
                    <div class="flex items-center mb-4">
                        <div class="h-8 w-8 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                            <span class="text-white font-bold text-sm">VCM</span>
                        </div>
                        <span class="ml-2 text-white font-bold text-lg">VCM Medical</span>
                    </div>
                    <p class="text-gray-300 text-sm">
                        VAMOS BIOTECH (Shanghai) Co., Ltd. - Bio-pharmaceutical company specializing in 
                        advanced life-cell based therapies with proven 95% treatment efficacy.
                    </p>
                </div>
                
                <div>
                    <h3 class="text-lg font-semibold mb-4">Quick Links</h3>
                    <ul class="space-y-2 text-sm text-gray-300">
                        <li><a href="#" class="hover:text-blue-400">Home</a></li>
                        <li><a href="#" class="hover:text-blue-400">About Us</a></li>
                        <li><a href="#" class="hover:text-blue-400">Medical Therapies</a></li>
                        <li><a href="#" class="hover:text-blue-400">Contact</a></li>
                    </ul>
                </div>
                
                <div>
                    <h3 class="text-lg font-semibold mb-4">Specialties</h3>
                    <ul class="space-y-2 text-sm text-gray-300">
                        <li><a href="#" class="hover:text-blue-400">Cancer (CAR-T)</a></li>
                        <li><a href="#" class="hover:text-blue-400">Autoimmune Disorders</a></li>
                        <li><a href="#" class="hover:text-blue-400">Eye Diseases</a></li>
                        <li><a href="#" class="hover:text-blue-400">Viral Infections</a></li>
                    </ul>
                </div>
                
                <div>
                    <h3 class="text-lg font-semibold mb-4">Contact</h3>
                    <div class="space-y-2 text-sm text-gray-300">
                        <p>Building #5, Lin Gang Fengxian Industrial Park</p>
                        <p>Shanghai 201413, P.R. China</p>
                        <p>info@vamosbiotech.com</p>
                    </div>
                </div>
            </div>
            
            <div class="border-t border-gray-800 mt-8 pt-6 text-center text-sm text-gray-400">
                <p>© 2024 VAMOS BIOTECH (Shanghai) Co., Ltd. All rights reserved.</p>
            </div>
        </div>
    </footer>
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
