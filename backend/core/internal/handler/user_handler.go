package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/model"
	"github.com/navakub/boardlog/backend/core/internal/service"
)

var userService service.UserService

func SetUserService(service service.UserService) {
	userService = service
}

func GetUsers(c *fiber.Ctx) error {
	users, err := userService.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	// user.Password = hashPassword(user.Password)

	err := userService.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user.Password = ""
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	update := new(model.User)
	if err := c.BodyParser(update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	update.ID = int64(id)

	err = userService.UpdateUser(update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	update.Password = ""
	return c.JSON(update)
}

func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	err = userService.DeleteUser(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "user deleted"})
}
