package routes

import (
	"ipenpoto/app/controllers"
	"ipenpoto/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	authPublic := api.Group("/auth")
	authPublic.Post("/login", controllers.Login)
	authPublic.Post("/register-account", controllers.RegisterAccount)
	authPublic.Post("/forgot-password", controllers.ForgotPassword)
}