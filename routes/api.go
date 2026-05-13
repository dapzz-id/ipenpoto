package routes

import (
	"ipenpoto/app/controllers"
	"ipenpoto/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctrl controllers.Controllers) {
	app.Use(middleware.SecurityLogger())

	api := app.Group("/api")
	api.Use(middleware.RateLimit(100, 60))

	authPublic := api.Group("/auth")
	authPublic.Use(middleware.RateLimit(5, 300))
	authPublic.Post("/login", ctrl.Auth.Login)
	authPublic.Post("/register", ctrl.Auth.Register)

	// Protected routes
	userProtected := api.Group("/user", middleware.JWTProtected())
	userProtected.Get("/profile", ctrl.User.GetProfile)
}
