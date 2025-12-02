package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/config"
	"github.com/navakub/boardlog/backend/core/internal/database"
	"github.com/navakub/boardlog/backend/core/internal/handler"
	"github.com/navakub/boardlog/backend/core/internal/repository"
	"github.com/navakub/boardlog/backend/core/internal/router"
	"github.com/navakub/boardlog/backend/core/internal/service"
)

func SetupApp() *fiber.App {
	cfg := config.LoadConfig()
	database.Connect(cfg)

	userRepo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserService(userRepo)
	handler.SetUserService(userService)
	authService := service.NewAuthService(userRepo)
	handler.SetAuthService(authService)

	app := fiber.New()
	router.SetupRoutes(app, userRepo)

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Server is running!",
		})
	})

	app.Listen(":" + cfg.AppPort)

	return app
}
