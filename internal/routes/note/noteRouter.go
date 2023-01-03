package noteRoutes

import (
	noteHandler "github.com/coerschkes/fiber-learning/internal/handlers/note"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/gofiber/fiber/v2"
)

type NoteRouter struct {
	noteHandler.NoteHandler
}

func NewNoteRouter(repository repository.NoteRepository) *NoteRouter {
	return &NoteRouter{noteHandler.NewNoteHttpHandler(repository)}
}

func (n NoteRouter) SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", n.CreateNote)
	// Read all Notes
	note.Get("/", n.GetNotes)
	// // Read one Note
	note.Get("/:noteId", n.GetNote)
	// // Update one Note
	note.Put("/:noteId", n.UpdateNote)
	// // Delete one Note
	note.Delete("/:noteId", n.DeleteNote)
}
