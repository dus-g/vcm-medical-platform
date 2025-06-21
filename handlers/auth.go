package handlers

import (
	"log"
	"time"
	"vcm-medical-platform/database"
	"vcm-medical-platform/models"
	"vcm-medical-platform/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	UserType int    `json:"userType" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6"`
}

type CompleteProfileRequest struct {
	Email         string    `json:"email" validate:"required,email"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name" validate:"required"`
	Gender        string    `json:"gender" validate:"required"`
	DateOfBirth   time.Time `json:"date_of_birth" validate:"required"`
	PhoneNumber   string    `json:"phone_number" validate:"required"`
	WechatId      string    `json:"wechat_id" validate:"required"`
	HeightCm      int       `json:"height_cm" validate:"required,min=50,max=300"`
	WeightKg      int       `json:"weight_kg" validate:"required,min=20,max=500"`
	MaritalStatus string    `json:"marital_status" validate:"required"`
	NoChildren    int       `json:"no_children"`
	Languages     string    `json:"languages"`
	Occupation    string    `json:"occupation"`
	Religion      string    `json:"religion"`
	CdCountry     int       `json:"cd_country" validate:"required"`
	CdState       int       `json:"cd_state" validate:"required"`
	CdCity        int       `json:"cd_city"`
	CdDistrict    int       `json:"cd_district"`
	StreetAddress string    `json:"street_address" validate:"required"`
	PostalCode    string    `json:"postal_code" validate:"required"`
}

// Register - Initial registration with email/password
func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to process password",
		})
	}

	// Generate OTP
	otpCode := utils.GenerateOTP()

	// Create user with pending status
	user := models.User{
		Email:        req.Email,
		Password:     hashedPassword,
		TyUser:       req.UserType,
		UserStatus:   "Registered", // Will be updated to Active after OTP verification
		OtpCode:      otpCode,
		OtpCreatedAt: time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Send OTP email
	if err := utils.SendOTPEmail(req.Email, otpCode); err != nil {
		log.Printf("Error sending OTP email: %v", err)
		// Don't fail registration if email fails
	}

	return c.JSON(fiber.Map{
		"message": "Registration successful. Please check your email for verification code.",
		"user_id": user.CdUser,
	})
}

// VerifyOTP - Verify email with OTP
func VerifyOTP(c *fiber.Ctx) error {
	var req VerifyOTPRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Check OTP validity (10 minutes)
	if time.Since(user.OtpCreatedAt) > 10*time.Minute {
		return c.Status(400).JSON(fiber.Map{
			"error": "OTP has expired. Please request a new one.",
		})
	}

	if user.OtpCode != req.OTP {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid OTP code",
		})
	}

	// Clear OTP and mark as verified
	user.OtpCode = ""
	user.OtpCreatedAt = time.Time{}
	
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to verify OTP",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email verified successfully",
		"user_id": user.CdUser,
	})
}

// CompleteProfile - Complete user profile after OTP verification
func CompleteProfile(c *fiber.Ctx) error {
	var req CompleteProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Update user profile
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Gender = req.Gender
	user.DateOfBirth = req.DateOfBirth
	user.PhoneNumber = req.PhoneNumber
	user.WechatId = req.WechatId
	user.HeightCm = req.HeightCm
	user.WeightKg = req.WeightKg
	user.MaritalStatus = req.MaritalStatus
	user.NoChildren = req.NoChildren
	user.Languages = req.Languages
	user.Occupation = req.Occupation
	user.Religion = req.Religion
	user.CdCountry = req.CdCountry
	user.CdState = req.CdState
	user.CdCity = req.CdCity
	user.CdDistrict = req.CdDistrict
	user.StreetAddress = req.StreetAddress
	user.PostalCode = req.PostalCode
	user.UserStatus = "Active" // Now fully registered

	if err := database.DB.Save(&user).Error; err != nil {
		log.Printf("Error updating user profile: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile completed successfully",
		"token":   token,
		"user": fiber.Map{
			"id":        user.CdUser,
			"email":     user.Email,
			"name":      user.GetFullName(),
			"userType":  user.TyUser,
			"status":    user.UserStatus,
		},
	})
}

// Login - User login
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.Preload("UserType").Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Check if user is active
	if user.UserStatus != "Active" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Account not activated. Please complete your registration.",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"user": fiber.Map{
			"id":        user.CdUser,
			"email":     user.Email,
			"name":      user.GetFullName(),
			"userType":  user.TyUser,
			"status":    user.UserStatus,
			"profileComplete": user.IsProfileComplete(),
		},
	})
}

// ResendOTP - Resend OTP code
func ResendOTP(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Generate new OTP
	otpCode := utils.GenerateOTP()
	user.OtpCode = otpCode
	user.OtpCreatedAt = time.Now()

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate new OTP",
		})
	}

	// Send OTP email
	if err := utils.SendOTPEmail(req.Email, otpCode); err != nil {
		log.Printf("Error sending OTP email: %v", err)
	}

	return c.JSON(fiber.Map{
		"message": "New verification code sent",
	})
}

// GetMe - Get current user info
func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := database.DB.Preload("UserType").Where("cd_user = ?", userID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":            user.CdUser,
			"email":         user.Email,
			"name":          user.GetFullName(),
			"firstName":     user.FirstName,
			"lastName":      user.LastName,
			"userType":      user.TyUser,
			"userTypeName":  user.UserType.UserTypeName,
			"status":        user.UserStatus,
			"profileComplete": user.IsProfileComplete(),
			"phone":         user.PhoneNumber,
			"gender":        user.Gender,
			"dateOfBirth":   user.DateOfBirth,
		},
	})
}
