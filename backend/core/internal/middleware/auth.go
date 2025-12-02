package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/repository"
	"github.com/navakub/boardlog/backend/core/internal/utils"
)

func JWTAuth(userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing access token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization format"})
		}

		tokenStr := parts[1]
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token missing after Bearer"})
		}

		now := time.Now().Unix()

		userID, accessExpired, err := utils.ValidateAccessTokenWithExpiry(tokenStr)
		if err != nil && now > accessExpired {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		if now > accessExpired {
			refreshToken := c.Cookies("refresh_token")
			if refreshToken == "" {
				refreshToken = c.Get("X-Refresh-Token")
			}
			if refreshToken == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "refresh token required"})
			}

			userID, refreshExpired, err := utils.ValidateRefreshTokenWithExpiry(refreshToken)
			if err != nil || now > refreshExpired {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired refresh token"})
			}

			newAccessToken, _, err := utils.CreateAccessToken(userID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate new access token"})
			}

			c.Set("Authorization", newAccessToken)
		}

		user, err := userRepo.GetByID(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
