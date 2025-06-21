package middleware

import (
	"strings"
	"vcm-medical-platform/internal/models"
	"vcm-medical-platform/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func JWTMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Authorization header required"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization format"})
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Verify user exists and is active
		var user models.User
		if err := db.First(&user, claims.UserID).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "User not found"})
		}

		if user.UserStatus != "Active" {
			return c.Status(401).JSON(fiber.Map{"error": "Account not activated"})
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("user_type", claims.UserType)
		c.Locals("user", user)
		return c.Next()
	}
}

func RoleMiddleware(requiredUserType int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Locals("user_type").(int)
		if userType != requiredUserType {
			return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
		}
		return c.Next()
	}
}
