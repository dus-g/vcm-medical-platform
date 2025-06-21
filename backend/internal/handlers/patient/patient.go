package patient

import (
	"vcm-medical-platform/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetDashboard(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Get stats
	var assessmentCount int64
	var appointmentCount int64
	var orderCount int64

	h.db.Model(&models.Assessment{}).Where("cd_user = ?", userID).Count(&assessmentCount)
	h.db.Model(&models.Appointment{}).Where("cd_user = ?", userID).Count(&appointmentCount)
	h.db.Model(&models.Order{}).Where("cd_user = ?", userID).Count(&orderCount)

	return c.JSON(fiber.Map{
		"user": user,
		"stats": fiber.Map{
			"assessments": assessmentCount,
			"appointments": appointmentCount,
			"orders": orderCount,
		},
		"message": "Dashboard data retrieved successfully",
	})
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{"user": user})
}

func (h *Handler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Languages   string `json:"languages"`
		Occupation  string `json:"occupation"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Update fields
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Languages = req.Languages
	user.Occupation = req.Occupation

	if err := h.db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully", "user": user})
}

func (h *Handler) GetAssessments(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var assessments []models.Assessment
	if err := h.db.Where("cd_user = ?", userID).Find(&assessments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch assessments"})
	}

	return c.JSON(fiber.Map{"assessments": assessments})
}

func (h *Handler) CreateAssessment(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req struct {
		DiseaseID int `json:"disease_id"`
		ProductID int `json:"product_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	assessment := models.Assessment{
		UserID:    userID,
		DiseaseID: req.DiseaseID,
		ProductID: req.ProductID,
		Status:    0, // In Progress
	}

	if err := h.db.Create(&assessment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create assessment"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Assessment created successfully", "assessment": assessment})
}

func (h *Handler) GetAppointments(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var appointments []models.Appointment
	if err := h.db.Where("cd_user = ?", userID).Preload("Doctor").Find(&appointments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch appointments"})
	}

	return c.JSON(fiber.Map{"appointments": appointments})
}

func (h *Handler) BookAppointment(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req struct {
		DoctorID        uint   `json:"doctor_id"`
		AppointmentDate string `json:"appointment_date"`
		AppointmentTime string `json:"appointment_time"`
		Notes           string `json:"notes"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	appointment := models.Appointment{
		DoctorID:        req.DoctorID,
		PatientID:       userID,
		AppointmentTime: req.AppointmentTime,
		Notes:           req.Notes,
		Status:          "scheduled",
	}

	if err := h.db.Create(&appointment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to book appointment"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Appointment booked successfully", "appointment": appointment})
}

func (h *Handler) GetOrders(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var orders []models.Order
	if err := h.db.Where("cd_user = ?", userID).Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}

	return c.JSON(fiber.Map{"orders": orders})
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req struct {
		TotalAmount    float64 `json:"total_amount"`
		OrderReference string  `json:"order_reference"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	order := models.Order{
		UserID:         userID,
		TotalAmount:    req.TotalAmount,
		OrderReference: req.OrderReference,
		Status:         "pending",
	}

	if err := h.db.Create(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create order"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Order created successfully", "order": order})
}
