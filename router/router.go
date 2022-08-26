package router

import (
	noteRoutes "github.com/coerschkes/fiber-learning/internal/routes/note"
	userRoutes "github.com/coerschkes/fiber-learning/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, database *gorm.DB) {
	api := app.Group("/api", logger.New())

	noteRouter := noteRoutes.NewNoteRouter(database)
	userRouter := userRoutes.NewUserRouter(database)

	noteRouter.SetupNoteRoutes(api)
	userRouter.SetupNoteRoutes(api)
}
