package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/model"
	"github.com/navakub/boardlog/backend/core/internal/service"
)

var authService service.AuthService

func SetAuthService(service service.AuthService) {
	authService = service
}

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
	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	err := authService.Register(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	user.Password = ""
	return c.Status(201).JSON(user)
}

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

	user.Password = ""
	return c.Status(200).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Assuming userID is obtained from a middleware that sets it in locals
	user := c.Locals("user").(*model.User)
	err := authService.Logout(uint(user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "logged out successfully"})
}
