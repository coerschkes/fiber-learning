package noteRoutes

import (
	noteHandler "github.com/coerschkes/fiber-learning/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NoteRouter struct {
	noteHandler.NoteHandler
}

func NewNoteRouter(database *gorm.DB) *NoteRouter {
	return &NoteRouter{noteHandler.NewNoteHttpHandler(database)}
}

func (n NoteRouter) SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", n.CreateNotes)
	// Read all Notes
	note.Get("/", n.GetNotes)
	// // Read one Note
	note.Get("/:noteId", n.GetNote)
	// // Update one Note
	note.Put("/:noteId", n.UpdateNote)
	// // Delete one Note
	note.Delete("/:noteId", n.DeleteNote)
}
