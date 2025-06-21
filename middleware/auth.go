package middleware

import (
	"strings"
	"vcm-medical-platform/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Authorization header required",
		})
	}

	// Check if it's a Bearer token
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid authorization format",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Set user info in context
	c.Locals("userID", claims.UserID)
	c.Locals("userEmail", claims.Email)
	c.Locals("userType", claims.UserType)

	return c.Next()
}

func RequireUserType(allowedTypes ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Locals("userType").(int)
		
		for _, allowedType := range allowedTypes {
			if userType == allowedType {
				return c.Next()
			}
		}
		
		return c.Status(403).JSON(fiber.Map{
			"error": "Insufficient permissions",
		})
	}
}
