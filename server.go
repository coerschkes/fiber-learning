package main

import (
	"github.com/coerschkes/fiber-learning/config"
	"github.com/coerschkes/fiber-learning/database"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database := database.DatabaseConnector.Connect(database.NewPostgresConnector())
	router.SetupRoutes(app, repository.NewPostgresNoteRepository(database))
	app.Listen(":" + config.LoadServerPort())
}
