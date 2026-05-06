package main

import (
	"log"
	"os"

	"ipenpoto/config"
	"ipenpoto/routes"
	// "ipenpoto/database/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

		// log.Println("Running seeder...")
		// seeders.Run(config.DB)
		// return
	}

	config.MigrateDB()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Use(logger.New())

	routes.SetupRoutes(app)

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