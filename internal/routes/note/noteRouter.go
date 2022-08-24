package noteRoutes

import (
	noteHandler "github.com/coerschkes/fiber-learning/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupNoteRoutes(router fiber.Router, database *gorm.DB) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", noteHandler.CreateNotes)
	// Read all Notes
	note.Get("/", noteHandler.GetNotes)
	// // Read one Note
	note.Get("/:noteId", noteHandler.GetNote)
	// // Update one Note
	note.Put("/:noteId", noteHandler.UpdateNote)
	// // Delete one Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
