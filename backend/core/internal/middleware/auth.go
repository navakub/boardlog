package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/repository"
)

func AuthMiddleware(userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}
		token := tokenParts[1]
		userID, err := validateTokenAndGetUserID(token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		user, err := userRepo.GetByID(userID)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}
		c.Locals("user", user)
		return c.Next()
	}
}

func validateTokenAndGetUserID(token string) (uint, error) {
	// Dummy implementation for illustration purposes.
	// Replace with actual token validation logic.
	if token == "valid-token" {
		return 1, nil // Assume user ID 1 for valid token
	}
	return 0, fiber.ErrUnauthorized
}
