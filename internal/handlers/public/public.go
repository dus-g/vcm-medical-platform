package public

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name":        "VCM Medical Platform",
		"version":     "1.0.0",
		"description": "Advanced Medical Treatment Platform with 95% Efficacy",
		"features": []string{
			"Multi-user authentication (8 user types)",
			"Medical assessment forms",
			"Appointment booking system",
			"Real-time chat support",
			"Order management",
			"Progress tracking",
		},
		"medical_specialties": []string{
			"Cancer Immunotherapy",
			"Autoimmune Disorders", 
			"Ophthalmology",
			"Neurological Sciences",
			"Respiratory Medicine",
			"Infectious Diseases",
		},
	})
}
