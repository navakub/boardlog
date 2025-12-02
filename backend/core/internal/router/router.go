package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/navakub/boardlog/backend/core/internal/handler"
	"github.com/navakub/boardlog/backend/core/internal/middleware"
	"github.com/navakub/boardlog/backend/core/internal/repository"
)

func SetupRoutes(app *fiber.App, userRepo repository.UserRepository) {
	// auth routes
	app.Post("api/auth/register", handler.Register)
	app.Post("api/auth/login", handler.Login)
	app.Get("/api/auth/me", middleware.JWTAuth(userRepo), handler.Me)
	app.Post("/api/auth/logout", middleware.JWTAuth(userRepo), handler.Logout)
	app.Post("/api/auth/refresh", handler.RefreshToken)

	// user routes
	app.Get("api/user", handler.GetUsers)
	app.Get("api/user/:id", handler.GetUser)
	app.Post("api/user", handler.CreateUser)
	app.Put("api/user/:id", handler.UpdateUser)
	app.Delete("api/user/:id", handler.DeleteUser)

	// playlog routes

	// boardgame routes

}
