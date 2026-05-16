package main

import (
	"log"
	"os"

	"ipenpoto/app/controllers"
	"ipenpoto/app/repositories"
	"ipenpoto/app/services"
	"ipenpoto/config"
	"ipenpoto/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	config.ConnectRedis()
	config.ConnectDB()

	args := os.Args

	switch {
	case len(args) > 1 && args[1] == "--migrate":
		log.Println("Running migrations...")
		config.MigrateDB()
		return

	case len(args) > 1 && args[1] == "--seed":
		log.Println("Running migration before seeding...")
		config.MigrateDB()

		// Seeder here
		return
	}

	config.MigrateDB()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins == "" {
		corsOrigins = "http://localhost:3000,http://localhost:5173"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	userRepo := repositories.NewUserRepository(config.DB)
	jwtService := services.NewJWTService()
	emailVerificationService := services.NewEmailVerificationService(config.RDB)
	emailService := services.NewEmailService()
	authService := services.NewAuthService(userRepo, jwtService, emailVerificationService, emailService)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userRepo)

	// Controller Container
	ctrl := controllers.Controllers{
		Auth: authController,
		User: userController,
	}

	// Routes
	routes.SetupRoutes(app, ctrl)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Ipenpoto API",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
