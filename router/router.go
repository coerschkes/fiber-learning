package router

import (
	noteRoutes "github.com/coerschkes/fiber-learning/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, database *gorm.DB) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api, database)
}
