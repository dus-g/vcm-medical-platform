package api

import (
	"vcm-medical-platform/internal/handlers/auth"
	"vcm-medical-platform/internal/handlers/patient"
	"vcm-medical-platform/internal/handlers/public"
	"vcm-medical-platform/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "VCM Medical Platform API",
			"version":   "1.0.0",
			"timestamp": fiber.Now(),
		})
	})

	// Initialize handlers
	authHandler := auth.NewHandler(db)
	patientHandler := patient.NewHandler(db)
	publicHandler := public.NewHandler(db)

	// API v1 group
	api := app.Group("/api/v1")

	// Public routes
	public := api.Group("/public")
	public.Get("/info", publicHandler.GetInfo)

	// Auth routes
	authGroup := api.Group("/auth")
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/verify-otp", authHandler.VerifyOTP)
	authGroup.Post("/forgot-password", authHandler.ForgotPassword)

	// Protected routes
	protected := api.Group("/", middleware.JWTMiddleware(db))
	protected.Get("/dashboard", patientHandler.GetDashboard)
	protected.Get("/profile", patientHandler.GetProfile)
	protected.Put("/profile", patientHandler.UpdateProfile)

	// Patient routes
	patients := protected.Group("/patient")
	patients.Get("/assessments", patientHandler.GetAssessments)
	patients.Post("/assessments", patientHandler.CreateAssessment)
	patients.Get("/appointments", patientHandler.GetAppointments)
	patients.Post("/appointments", patientHandler.BookAppointment)
	patients.Get("/orders", patientHandler.GetOrders)
	patients.Post("/orders", patientHandler.CreateOrder)
}
