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
                        "version":   "2.0.0",
                })
        })

        // API routes
        api := app.Group("/api/v1")
        
        api.Get("/info", func(c *fiber.Ctx) error {
                return c.JSON(fiber.Map{
                        "name":        "VCM Medical Platform",
                        "description": "Advanced Medical Treatment Platform with 95% Efficacy",
                        "status":      "running",
                        "version":     "2.0.0",
                        "features": []string{
                                "Modern React Frontend",
                                "Medical assessment forms", 
                                "Appointment booking",
                                "Photo progress tracking",
                                "Treatment protocols",
                                "Doctor consultations",
                        },
                })
        })

        // Serve the modern HTML for all routes
        app.Get("*", func(c *fiber.Ctx) error {
                return c.Type("html").SendString(getModernHTML())
        })

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("🚀 VCM Medical Platform v2.0 starting on port %s", port)
        log.Fatal(app.Listen(":" + port))
}

func getModernHTML() string {
        return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>VCM Medical Platform - Advanced Treatment Solutions</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Inter', sans-serif; }
        .gradient-bg { background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); }
        .pulse { animation: pulse 2s infinite; }
        @keyframes pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.5; } }
    </style>
</head>
<body class="bg-gray-50">
    <!-- Header -->
    <header class="fixed w-full top-0 z-50 bg-white/95 backdrop-blur-md border-b border-gray-100 shadow-sm">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex items-center justify-between h-20">
                <!-- Logo -->
                <div class="flex items-center">
                    <div class="h-16 w-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center transition-transform hover:scale-105">
                        <span class="text-white font-bold text-xl">VCM</span>
                    </div>
                    <span class="ml-3 text-xl font-bold text-gray-900">VCM Medical</span>
                </div>
                
                <!-- Desktop Navigation -->
                <nav class="hidden lg:flex items-center space-x-8">
                    <a href="#" class="text-blue-600 bg-blue-50 font-medium px-3 py-2 rounded-lg">Home</a>
                    <a href="#therapies" class="text-gray-700 hover:text-blue-600 font-medium px-3 py-2 rounded-lg transition-colors">Therapies</a>
                    <a href="#about" class="text-gray-700 hover:text-blue-600 font-medium px-3 py-2 rounded-lg transition-colors">About</a>
                    <a href="#contact" class="text-gray-700 hover:text-blue-600 font-medium px-3 py-2 rounded-lg transition-colors">Contact</a>
                </nav>
                
                <!-- Right Side Actions -->
                <div class="hidden lg:flex items-center space-x-3">
                    <button class="relative p-3 text-gray-700 hover:text-blue-600 hover:bg-gray-50 rounded-lg transition-colors">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-1.5 6M7 13l-1.5 6m0 0h9m-9 0V9a3 3 0 013-3h4a3 3 0 013 3v10"></path>
                        </svg>
                        <span class="absolute -top-1 -right-1 bg-blue-600 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center font-bold">2</span>
                    </button>
                    <div class="w-9 h-9 bg-gray-100 rounded-full flex items-center justify-center">
                        <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                        </svg>
                    </div>
                </div>
                
                <!-- Mobile menu button -->
                <button class="lg:hidden p-2 rounded-lg text-gray-700 hover:text-blue-600">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                    </svg>
                </button>
            </div>
        </div>
    </header>

    <!-- Main Content -->
    <main class="pt-20">
        <!-- Hero Section -->
        <section class="pt-16 pb-16 px-6 bg-gradient-to-br from-blue-50 via-blue-50 to-cyan-50">
            <div class="max-w-7xl mx-auto text-center">
                <div class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-blue-100 to-blue-100 border border-blue-200 rounded-full text-blue-700 text-sm font-medium mb-6">
                    <div class="w-2 h-2 bg-blue-500 rounded-full mr-2 pulse"></div>
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
                    <button class="px-8 py-3 bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white font-medium rounded-lg shadow-lg transition-all transform hover:scale-105">
                        Start Treatment
                    </button>
                    <button class="px-8 py-3 bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium rounded-lg transition-colors">
                        Explore Therapies
                    </button>
                </div>

                <!-- Company Highlights -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 max-w-6xl mx-auto">
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20 hover:shadow-lg transition-shadow">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🏆</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">95% Treatment Efficacy</h3>
                        <p class="text-sm text-gray-600">Breakthrough results in melanoma and cancer treatment</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20 hover:shadow-lg transition-shadow">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🔬</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">World's First</h3>
                        <p class="text-sm text-gray-600">Clinical trials for antibiotic-resistant infections</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20 hover:shadow-lg transition-shadow">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">🌍</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">24/7 Platform Access</h3>
                        <p class="text-sm text-gray-600">Global operations with Shanghai headquarters</p>
                    </div>
                    
                    <div class="bg-white/80 backdrop-blur-sm rounded-xl p-6 shadow-sm border border-white/20 hover:shadow-lg transition-shadow">
                        <div class="w-12 h-12 bg-blue-100 rounded-lg mb-4 mx-auto flex items-center justify-center">
                            <span class="text-2xl">⚡</span>
                        </div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Advanced Life-Cell Therapies</h3>
                        <p class="text-sm text-gray-600">Cutting-edge medical and naturopathic treatments</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Treatment Process -->
        <section class="py-16 px-6 bg-white">
            <div class="max-w-7xl mx-auto">
                <div class="text-center mb-12">
                    <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
                        Your Treatment Journey
                    </h2>
                    <p class="text-xl text-gray-600 max-w-3xl mx-auto">
                        Simple steps to access our breakthrough medical treatments
                    </p>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
                    <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 text-center hover:shadow-xl transition-shadow">
                        <div class="w-12 h-12 text-blue-600 mx-auto mb-4">
                            <svg fill="currentColor" viewBox="0 0 24 24"><path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/></svg>
                        </div>
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">1</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Register & Consultation</h3>
                        <p class="text-gray-600 text-sm mb-3">Sign up and schedule your initial medical consultation</p>
                        <div class="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                            Create account, book appointment
                        </div>
                    </div>

                    <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 text-center hover:shadow-xl transition-shadow">
                        <div class="w-12 h-12 text-blue-600 mx-auto mb-4">
                            <svg fill="currentColor" viewBox="0 0 24 24"><path d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/></svg>
                        </div>
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">2</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Medical Assessment</h3>
                        <p class="text-gray-600 text-sm mb-3">Complete comprehensive 50+ field assessment forms</p>
                        <div class="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                            Psoriasis or Eye Disease assessment
                        </div>
                    </div>

                    <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 text-center hover:shadow-xl transition-shadow">
                        <div class="w-12 h-12 text-blue-600 mx-auto mb-4">
                            <svg fill="currentColor" viewBox="0 0 24 24"><path d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,17L7,12L8.41,10.59L12,14.17L15.59,10.59L17,12L12,17Z"/></svg>
                        </div>
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">3</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Treatment Protocol</h3>
                        <p class="text-gray-600 text-sm mb-3">Receive personalized treatment plan</p>
                        <div class="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                            Custom protocol with schedule
                        </div>
                    </div>

                    <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 text-center hover:shadow-xl transition-shadow">
                        <div class="w-12 h-12 text-blue-600 mx-auto mb-4">
                            <svg fill="currentColor" viewBox="0 0 24 24"><path d="M7,18C5.9,18 5,18.9 5,20S5.9,22 7,22 9,21.1 9,20 8.1,18 7,18M1,2V4H3L6.6,11.59L5.24,14.04C5.09,14.32 5,14.65 5,15A2,2 0 0,0 7,17H19V15H7.42A0.25,0.25 0 0,1 7.17,14.75C7.17,14.7 7.18,14.66 7.2,14.63L8.1,13H15.55C16.3,13 16.96,12.58 17.3,11.97L20.88,5H5.21L4.27,3H1M17,18C15.9,18 15,18.9 15,20S15.9,22 17,22 19,21.1 19,20 18.1,18 17,18Z"/></svg>
                        </div>
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">4</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Purchase Products</h3>
                        <p class="text-gray-600 text-sm mb-3">Buy prescribed products for treatment</p>
                        <div class="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                            Secure payment, fast shipping
                        </div>
                    </div>

                    <div class="bg-gradient-to-br from-white to-gray-50 rounded-2xl p-6 shadow-lg border border-gray-100 text-center hover:shadow-xl transition-shadow">
                        <div class="w-12 h-12 text-blue-600 mx-auto mb-4">
                            <svg fill="currentColor" viewBox="0 0 24 24"><path d="M16,6L18.29,8.29L13.41,13.17L9.41,9.17L2,16.59L3.41,18L9.41,12L13.41,16L19.71,9.71L22,12V6H16Z"/></svg>
                        </div>
                        <div class="w-8 h-8 bg-blue-600 text-white rounded-full text-sm font-bold mb-3 mx-auto flex items-center justify-center">5</div>
                        <h3 class="text-lg font-bold text-gray-900 mb-2">Ongoing Monitoring</h3>
                        <p class="text-gray-600 text-sm mb-3">Track progress with photo uploads</p>
                        <div class="text-xs text-blue-600 font-medium bg-blue-50 px-3 py-2 rounded-lg">
                            Photo tracking, progress reports
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Treatment Categories -->
        <section id="therapies" class="py-16 px-6 bg-gray-50">
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
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all border-l-4 border-l-red-500">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">🛡️</span>
                            <h3 class="text-xl font-bold text-gray-900">Cancers</h3>
                        </div>
                        <ul class="space-y-2">
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">B-cell leukemia, lymphoma (CAR-T)</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Melanoma therapy</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Colorectal, pancreas cancer</span>
                            </li>
                        </ul>
                        <div class="mt-4">
                            <button class="w-full bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium py-2 px-4 rounded-lg transition-colors text-sm">
                                Learn More
                            </button>
                        </div>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all border-l-4 border-l-blue-500">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">❤️</span>
                            <h3 class="text-xl font-bold text-gray-900">Autoimmune</h3>
                        </div>
                        <ul class="space-y-2">
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Psoriasis Vulgaris</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Rheumatoid arthritis</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Lupus, Hashimoto's disease</span>
                            </li>
                        </ul>
                        <div class="mt-4">
                            <button class="w-full bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium py-2 px-4 rounded-lg transition-colors text-sm">
                                Learn More
                            </button>
                        </div>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all border-l-4 border-l-cyan-500">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">👁️</span>
                            <h3 class="text-xl font-bold text-gray-900">Eye Diseases</h3>
                        </div>
                        <ul class="space-y-2">
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Macular degeneration</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Glaucoma</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Lazy eye (Amblyopia)</span>
                            </li>
                        </ul>
                        <div class="mt-4">
                            <button class="w-full bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium py-2 px-4 rounded-lg transition-colors text-sm">
                                Learn More
                            </button>
                        </div>
                    </div>
                    
                    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 hover:shadow-xl transition-all border-l-4 border-l-green-500">
                        <div class="flex items-center mb-4">
                            <span class="text-2xl mr-3">🫁</span>
                            <h3 class="text-xl font-bold text-gray-900">Respiratory</h3>
                        </div>
                        <ul class="space-y-2">
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Tuberculosis</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Pneumonia & Bronchitis</span>
                            </li>
                            <li class="flex items-start">
                                <svg class="w-4 h-4 text-blue-500 mt-0.5 mr-2 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                                </svg>
                                <span class="text-gray-700 text-sm">Gastritis & H. Pylori</span>
                            </li>
                        </ul>
                        <div class="mt-4">
                            <button class="w-full bg-gray-200 hover:bg-gray-300 text-gray-900 font-medium py-2 px-4 rounded-lg transition-colors text-sm">
                                Learn More
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Call to Action -->
        <section class="py-16 px-6 bg-gradient-to-r from-blue-600 to-blue-700">
            <div class="max-w-4xl mx-auto text-center">
                <h2 class="text-3xl md:text-4xl font-bold text-white mb-6">
                    Ready to Begin Your Treatment?
                </h2>
                <p class="text-xl text-blue-100 mb-8">
                    Join thousands of patients who have experienced breakthrough results with our advanced therapies.
                </p>
                <div class="flex flex-col sm:flex-row gap-4 justify-center">
                    <button class="px-8 py-3 bg-white text-blue-600 hover:bg-gray-100 font-medium rounded-lg shadow-lg transition-colors transform hover:scale-105">
                        Start Your Journey
                    </button>
                    <button class="px-8 py-3 border-2 border-white text-white hover:bg-white hover:text-blue-600 font-medium rounded-lg transition-colors">
                        Contact Our Team
                    </button>
                </div>
            </div>
        </section>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-900 text-white py-12">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
                <!-- Company Info -->
                <div>
                    <div class="flex items-center mb-4">
                        <div class="h-8 w-8 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                            <span class="text-white font-bold text-sm">VCM</span>
                        </div>
                        <span class="ml-2 text-white font-bold text-lg">VCM Medical</span>
                    </div>
                    <p class="text-gray-300 text-sm mb-6">
                        VAMOS BIOTECH (Shanghai) Co., Ltd. - Bio-pharmaceutical company specializing in 
                        advanced life-cell based therapies with proven 95% treatment efficacy.
                    </p>
                    <div class="space-y-3">
                        <div class="flex items-center space-x-2 text-sm text-gray-300">
                            <span class="text-blue-400">🛡️</span>
                            <span>95% Treatment Efficacy</span>
                        </div>
                        <div class="flex items-center space-x-2 text-sm text-gray-300">
                            <span class="text-blue-400">⏰</span>
                            <span>World's First Antibiotic Resistance Trials</span>
                        </div>
                        <div class="flex items-center space-x-2 text-sm text-gray-300">
                            <span class="text-blue-400">👥</span>
                            <span>Led by Prof. Sergey I. Chernysh</span>
                        </div>
                        <div class="flex items-center space-x-2 text-sm text-gray-300">
                            <span class="text-blue-400">❤️</span>
                            <span>Shanghai Global Operations</span>
                        </div>
                    </div>
                </div>
                
                <!-- Quick Links -->
                <div>
                    <h3 class="text-lg font-semibold mb-6">Quick Links</h3>
                    <ul class="space-y-3">
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Home
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            About Us
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Medical Therapies
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Contact
                        </a></li>
                    </ul>
                </div>
                
                <!-- Treatment Specialties -->
                <div>
                    <h3 class="text-lg font-semibold mb-6">Treatment Specialties</h3>
                    <ul class="space-y-3">
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Cancers (CAR-T, Melanoma)
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Viral Infections (RNA/DNA)
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Autoimmune Disorders
                        </a></li>
                        <li><a href="#" class="text-gray-300 hover:text-blue-400 text-sm flex items-center group">
                            <svg class="w-3 h-3 mr-2 opacity-0 group-hover:opacity-100 transition-opacity" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                            </svg>
                            Eye Diseases
                        </a></li>
                    </ul>
                </div>
                
                <!-- Contact Info -->
                <div id="contact">
                    <h3 class="text-lg font-semibold mb-6">Contact Us</h3>
                    <div class="space-y-4">
                        <div class="flex items-start space-x-3">
                            <svg class="w-5 h-5 text-blue-400 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd"></path>
                            </svg>
                            <div class="text-sm text-gray-300">
                                <p>Building #5, Lin Gang Fengxian Industrial Park</p>
                                <p>1800 Xin Yang Road, Feng Xian District</p>
                                <p>Shanghai 201413, P.R. China</p>
                            </div>
                        </div>
                        
                        <div class="flex items-center space-x-3">
                            <svg class="w-5 h-5 text-blue-400 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                <path d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"></path>
                            </svg>
                            <span class="text-sm text-gray-300">+86 (21) 1234-5678</span>
                        </div>
                        
                        <div class="flex items-center space-x-3">
                            <svg class="w-5 h-5 text-blue-400 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                                <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"></path>
                                <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"></path>
                            </svg>
                            <span class="text-sm text-gray-300">info@vamosbiotech.com</span>
                        </div>
                    </div>

                    <!-- Company Details -->
                    <div class="mt-6">
                        <h4 class="text-sm font-semibold mb-3 text-gray-200">Company Details</h4>
                        <div class="space-y-1 text-xs text-gray-400">
                            <div>Registration: 91310000MAH3AQB3D</div>
                            <div>Founded: 2014</div>
                            <div>Startup with Global Operations</div>
                            <div>Led by Prof. Sergey I. Chernysh</div>
                        </div>
                    </div>
                </div>
            </div>
            
            <!-- Newsletter Signup -->
            <div class="mt-12 pt-8 border-t border-gray-800">
                <div class="max-w-md mx-auto text-center lg:max-w-none lg:text-left lg:flex lg:items-center lg:justify-between">
                    <div class="lg:flex-1">
                        <h3 class="text-lg font-semibold mb-2">Clinical Updates</h3>
                        <p class="text-gray-300 text-sm">
                            Get the latest updates on breakthrough treatments, clinical trials, and medical research.
                        </p>
                    </div>
                    <div class="mt-4 lg:mt-0 lg:ml-8">
                        <div class="flex flex-col sm:flex-row max-w-md">
                            <input
                                type="email"
                                placeholder="Enter your email"
                                class="px-4 py-2 bg-gray-800 border border-gray-700 rounded-l-md sm:rounded-r-none focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-sm flex-1"
                            />
                            <button class="px-6 py-2 bg-blue-600 hover:bg-blue-700 transition-colors rounded-r-md sm:rounded-l-none text-sm font-medium">
                                Subscribe
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Bottom Footer -->
            <div class="border-t border-gray-800 mt-8 pt-6">
                <div class="flex flex-col md:flex-row md:items-center md:justify-between">
                    <div class="text-sm text-gray-400">
                        <p>© 2024 VAMOS BIOTECH (Shanghai) Co., Ltd. All rights reserved.</p>
                        <p class="mt-1">
                            Bio-pharmaceutical innovation for advanced life-cell based therapies.
                        </p>
                    </div>
                    
                    <div class="mt-4 md:mt-0 flex flex-col sm:flex-row sm:items-center sm:space-x-6">
                        <!-- Legal Links -->
                        <div class="flex flex-wrap items-center space-x-4 text-xs text-gray-400">
                            <a href="#" class="hover:text-blue-400 transition-colors">Privacy Policy</a>
                            <span class="text-gray-600">•</span>
                            <a href="#" class="hover:text-blue-400 transition-colors">Terms of Service</a>
                            <span class="text-gray-600">•</span>
                            <a href="#" class="hover:text-blue-400 transition-colors">Medical Disclaimer</a>
                            <span class="text-gray-600">•</span>
                            <a href="#" class="hover:text-blue-400 transition-colors">Cookie Policy</a>
                        </div>

                        <!-- Social Links -->
                        <div class="flex items-center space-x-4 mt-3 sm:mt-0">
                            <a href="#" class="text-gray-400 hover:text-blue-400 transition-colors" aria-label="Facebook">
                                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                                    <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                                </svg>
                            </a>
                            <a href="#" class="text-gray-400 hover:text-blue-400 transition-colors" aria-label="Twitter">
                                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                                    <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
                                </svg>
                            </a>
                            <a href="#" class="text-gray-400 hover:text-blue-400 transition-colors" aria-label="LinkedIn">
                                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                                    <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
                                </svg>
                            </a>
                            <a href="#" class="text-gray-400 hover:text-blue-400 transition-colors" aria-label="Instagram">
                                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                                    <path d="M12.017 0C5.396 0 .029 5.367.029 11.987c0 6.618 5.367 11.986 11.988 11.986S24.005 18.605 24.005 11.987C24.005 5.367 18.638.001 12.017.001zm5.568 16.554c-.397.794-1.154 1.55-1.947 1.948-.793.396-1.67.396-2.464.396H8.827c-.794 0-1.671 0-2.464-.396-.794-.398-1.55-1.154-1.948-1.948-.396-.793-.396-1.67-.396-2.464V8.827c0-.794 0-1.671.396-2.464.398-.794 1.154-1.55 1.948-1.948.793-.396 1.67-.396 2.464-.396h4.347c.794 0 1.671 0 2.464.396.793.398 1.55 1.154 1.947 1.948.397.793.397 1.67.397 2.464v4.347c0 .794 0 1.671-.397 2.464z"/>
                                </svg>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </footer>
</body>
</html>`
}
