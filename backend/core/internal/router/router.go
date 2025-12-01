package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	// auth routes
	app.Post("api/auth/register", handler.Register)
	app.Post("api/auth/login", handler.Login)
	app.Post("api/auth/logout", handler.Logout)

	// user routes
	app.Get("api/user", handler.GetUsers)
	app.Get("api/user/:id", handler.GetUser)
	app.Post("api/user", handler.CreateUser)
	app.Put("api/user/:id", handler.UpdateUser)
	app.Delete("api/user/:id", handler.DeleteUser)

	// playlog routes

	// boardgame routes

}
