package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	// auth routes
	// app.Post("/register", handler.Register)
	// app.Post("/login", handler.Login)
	// app.Post("/logout", handler.Logout)

	// user routes
	app.Get("/user", handler.GetUsers)
	app.Get("/user/:id", handler.GetUser)
	app.Post("/user", handler.CreateUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)

	// playlog routes

	// boardgame routes

}
