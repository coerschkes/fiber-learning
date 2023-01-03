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
	note.Post("/", n.CreateNote)
	note.Get("/", n.FindNotes)
	note.Get("/:noteId", n.FindNote)
	note.Put("/:noteId", n.UpdateNote)
	note.Delete("/:noteId", n.DeleteNote)
}
