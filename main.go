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
                        "status":  "ok",
                        "message": "VCM Medical Platform API",
                        "version": "2.0.0",
                })
        })

        // API info
        app.Get("/api/v1/info", func(c *fiber.Ctx) error {
                return c.JSON(fiber.Map{
                        "name":        "VCM Medical Platform",
                        "description": "Advanced Medical Treatment Platform",
                        "status":      "running",
                        "version":     "2.0.0",
                })
        })

        // Serve the modern HTML for all other routes
        app.Get("*", func(c *fiber.Ctx) error {
                return c.Type("html").SendString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform - Advanced Treatment Solutions</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body { font-family: system-ui, -apple-system, sans-serif; }
        .gradient-bg { background: linear-gradient(135deg, #3B82F6 0%, #8B5CF6 100%); }
        .pulse { animation: pulse 2s infinite; }
        @keyframes pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.5; } }
    </style>
</head>
<body class="bg-gray-50">
    <!-- Header -->
    <header class="fixed w-full top-0 z-50 bg-white/95 backdrop-blur border-b border-gray-200 shadow-sm">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex items-center justify-between h-20">
                <div class="flex items-center">
                    <div class="h-16 w-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                        <span class="text-white font-bold text-xl">VCM</span>
                    </div>
                    <span class="ml-3 text-xl font-bold text-gray-900">VCM Medical</span>
                </div>
                <nav class="hidden lg:flex items-center space-x-8">
                    <a href="#" class="text-blue-600 bg-blue-50 font-medium px-3 py-2 rounded-lg">Home</a>
                    <a href="#therapies" class="text-gray-700 hover:text-blue-600 font-medium">Therapies</a>
                    <a href="#about" class="text-gray-700 hover:text-blue-600 font-medium">About</a>
                    <a href="#contact" class="text-gray-700 hover:text-blue-600 font-medium">Contact</a>
                </nav>
            </div>
        </div>
    </header>

    <!-- Hero Section -->
    <main class="pt-20">
        <section class="pt-16 pb-16 px-6 bg-gradient-to-br from-blue-50 to-cyan-50">
            <div class="max-w-7xl mx-auto text-center">
                <div class="inline-flex items-center px-4 py-2 bg-blue-100 border border-blue-200 rounded-full text-blue-700 text-sm font-medium mb-6">
                    <div class="w-2 h-2 bg-blue-500 rounded-full mr-2 pulse"></div>
                    VAMOS BIOTECH - Bio-Pharmaceutical Innovation
                </div>
                
                <h1 class="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
                    Advanced Medical
                    <span class="block text-blue-600">Treatment Platform</span>
                </h1>
                
                <p class="text-xl text-gray-600 mb-8 max-w-4xl mx-auto">
                    Breakthrough life-cell based therapies for cancer, viral infections, autoimmune disorders, 
                    and antibiotic-resistant bacterial infections with proven 95% efficacy rates.
                </p>
                
                <div class="flex gap-4 justify-center mb-12">
                    <button class="px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg shadow-lg">
                        Start Treatment
                    </button>
                    <button class="px-8 py-3 bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium rounded-lg">
                        Explore Therapies
                    </button>
                </div>

                <!-- Highlights -->
                <div class="grid grid-cols-1 md:grid-cols-4 gap-6 max-w-6xl mx-auto">
                    <div class="bg-white/80 rounded-xl p-6 shadow border">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🏆</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">95% Treatment Efficacy</h3>
                        <p class="text-sm text-gray-600">Breakthrough results in melanoma and cancer treatment</p>
                    </div>
                    <div class="bg-white/80 rounded-xl p-6 shadow border">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🔬</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">World's First</h3>
                        <p class="text-sm text-gray-600">Clinical trials for antibiotic-resistant infections</p>
                    </div>
                    <div class="bg-white/80 rounded-xl p-6 shadow border">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🌍</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">24/7 Platform Access</h3>
                        <p class="text-sm text-gray-600">Global operations with Shanghai headquarters</p>
                    </div>
                    <div class="bg-white/80 rounded-xl p-6 shadow border">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">⚡</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Advanced Therapies</h3>
                        <p class="text-sm text-gray-600">Cutting-edge medical treatments</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Treatment Process -->
        <section class="py-16 px-6 bg-white">
            <div class="max-w-7xl mx-auto">
                <div class="text-center mb-12">
                    <h2 class="text-3xl font-bold text-gray-900 mb-4">Your Treatment Journey</h2>
                    <p class="text-xl text-gray-600">Simple steps to access our breakthrough medical treatments</p>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
                    <div class="bg-white rounded-2xl p-6 shadow-lg border text-center">
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">1</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Register</h3>
                        <p class="text-gray-600 text-sm">Sign up and schedule consultation</p>
                    </div>
                    <div class="bg-white rounded-2xl p-6 shadow-lg border text-center">
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">2</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Assessment</h3>
                        <p class="text-gray-600 text-sm">Complete medical forms</p>
                    </div>
                    <div class="bg-white rounded-2xl p-6 shadow-lg border text-center">
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">3</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Protocol</h3>
                        <p class="text-gray-600 text-sm">Receive treatment plan</p>
                    </div>
                    <div class="bg-white rounded-2xl p-6 shadow-lg border text-center">
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">4</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Purchase</h3>
                        <p class="text-gray-600 text-sm">Buy prescribed products</p>
                    </div>
                    <div class="bg-white rounded-2xl p-6 shadow-lg border text-center">
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">5</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Monitor</h3>
                        <p class="text-gray-600 text-sm">Track your progress</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Medical Conditions -->
        <section id="therapies" class="py-16 px-6 bg-gray-50">
            <div class="max-w-7xl mx-auto">
                <div class="text-center mb-12">
                    <h2 class="text-3xl font-bold text-gray-900 mb-4">Medical Conditions We Treat</h2>
                    <p class="text-xl text-gray-600">Comprehensive treatment protocols across major medical specialties</p>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
                    <div class="bg-white rounded-lg shadow border-l-4 border-l-red-500 p-6">
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
                    <div class="bg-white rounded-lg shadow border-l-4 border-l-blue-500 p-6">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">❤️</span>
                            <h3 class="text-xl font-bold text-gray-900">Autoimmune</h3>
                        </div>
                        <ul class="space-y-2 text-sm text-gray-700">
                            <li>• Psoriasis Vulgaris</li>
                            <li>• Rheumatoid arthritis</li>
                            <li>• Lupus, Hashimoto's</li>
                        </ul>
                    </div>
                    <div class="bg-white rounded-lg shadow border-l-4 border-l-cyan-500 p-6">
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
                    <div class="bg-white rounded-lg shadow border-l-4 border-l-green-500 p-6">
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

        <!-- CTA -->
        <section class="py-16 px-6 bg-gradient-to-r from-blue-600 to-blue-700">
            <div class="max-w-4xl mx-auto text-center">
                <h2 class="text-3xl font-bold text-white mb-6">Ready to Begin Your Treatment?</h2>
                <p class="text-xl text-blue-100 mb-8">Join thousands of patients who have experienced breakthrough results.</p>
                <div class="flex gap-4 justify-center">
                    <button class="px-8 py-3 bg-white text-blue-600 hover:bg-gray-100 font-medium rounded-lg">Start Your Journey</button>
                    <button class="px-8 py-3 border-2 border-white text-white hover:bg-white hover:text-blue-600 font-medium rounded-lg">Contact Our Team</button>
                </div>
            </div>
        </section>
    </main>

    <!-- Footer -->
    <footer id="contact" class="bg-gray-900 text-white py-12">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
                <div>
                    <div class="flex items-center mb-4">
                        <div class="h-8 w-8 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                            <span class="text-white font-bold text-sm">VCM</span>
                        </div>
                        <span class="ml-2 text-white font-bold text-lg">VCM Medical</span>
                    </div>
                    <p class="text-gray-300 text-sm">VAMOS BIOTECH (Shanghai) Co., Ltd. - Bio-pharmaceutical company specializing in advanced life-cell based therapies with proven 95% treatment efficacy.</p>
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
                        <p>+86 (21) 1234-5678</p>
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

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("🚀 VCM Medical Platform v2.0 starting on port %s", port)
        log.Fatal(app.Listen(":" + port))
}
