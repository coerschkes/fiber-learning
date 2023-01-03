package router

import (
	noteRoutes "github.com/coerschkes/fiber-learning/internal/routes/note"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, database *gorm.DB) {
	api := app.Group("/api", logger.New())

	noteRouter := noteRoutes.NewNoteRouter(repository.NewPostgresNoteRepository(database))

	noteRouter.SetupNoteRoutes(api)
}
