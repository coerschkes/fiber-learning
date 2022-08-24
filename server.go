package main

import (
	"github.com/coerschkes/fiber-learning/database"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
