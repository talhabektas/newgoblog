package main

import (
	"awesomeProject2/database"
	"awesomeProject2/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error in .env file")
	}
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}
