package router

import (
	noteRoutes "github.com/coerschkes/fiber-learning/internal/routes/note"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, repository repository.NoteRepository) {
	api := app.Group("/api", logger.New())

	noteRouter := noteRoutes.NewNoteRouter(repository)

	noteRouter.SetupNoteRoutes(api)
}
