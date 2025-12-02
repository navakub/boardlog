package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/model"
	"github.com/navakub/boardlog/backend/core/internal/service"
	"github.com/navakub/boardlog/backend/core/internal/utils"
)

var authService service.AuthService

func SetAuthService(service service.AuthService) {
	authService = service
}

// ---------------------- REGISTER ----------------------
func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := authService.Register(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user.Password = "" // hide password
	return c.Status(201).JSON(user)
}

// ---------------------- LOGIN ----------------------
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	user, err := authService.Login(input.Email, input.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Create tokens
	accessToken, accessExp, err := utils.CreateAccessToken(uint(user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create access token"})
	}

	refreshToken, refreshExp, err := utils.CreateRefreshToken(uint(user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create refresh token"})
	}

	user.Password = "" // hide password
	return c.Status(200).JSON(fiber.Map{
		"user":          user,
		"access_token":  accessToken,
		"access_exp":    accessExp,
		"refresh_token": refreshToken,
		"refresh_exp":   refreshExp,
	})
}

// ---------------------- ME ----------------------
func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	fetchedUser, err := authService.GetCurrentUser(uint(user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	fetchedUser.Password = ""
	return c.Status(200).JSON(fetchedUser)
}

// ---------------------- LOGOUT ----------------------
func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.User)

	// Optionally: invalidate refresh token in DB/cache
	if err := authService.Logout(uint(user.ID)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "logged out successfully"})
}

// ---------------------- REFRESH TOKEN ----------------------
func RefreshToken(c *fiber.Ctx) error {
	type RefreshInput struct {
		RefreshToken string `json:"refresh_token"`
	}

	var input RefreshInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Validate refresh token
	userID, exp, err := utils.ValidateRefreshTokenWithExpiry(input.RefreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if refresh token is expired
	if time.Now().Unix() > exp {
		return c.Status(401).JSON(fiber.Map{"error": "refresh token expired"})
	}

	// Create new access token
	accessToken, accessExp, err := utils.CreateAccessToken(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create access token"})
	}

	response := fiber.Map{
		"access_token": accessToken,
		"access_exp":   accessExp,
	}

	// Optionally: issue new refresh token if it will expire soon
	if exp-time.Now().Unix() < 24*3600 { // less than 1 day left
		refreshToken, refreshExp, err := utils.CreateRefreshToken(userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create refresh token"})
		}
		response["refresh_token"] = refreshToken
		response["refresh_exp"] = refreshExp
	}

	return c.Status(200).JSON(response)
}
