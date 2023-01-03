package noteHandler

import (
	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//todo: encapsulate err obj and map to fiber err to always return err if necessary

type NoteHandler interface {
	GetNotes(c *fiber.Ctx) error
	GetNote(c *fiber.Ctx) error
	CreateNote(c *fiber.Ctx) error
	UpdateNote(c *fiber.Ctx) error
	DeleteNote(c *fiber.Ctx) error
}
type NoteHttpHandler struct {
	repository repository.NoteRepository
}
type updateNote struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Text     string `json:"Text"`
}

func NewNoteHttpHandler(repository repository.NoteRepository) *NoteHttpHandler {
	return &NoteHttpHandler{repository}
}

func (h NoteHttpHandler) GetNotes(c *fiber.Ctx) error {
	notes := h.repository.FindAll()
	if len(notes) == 0 {
		return h.createFiberError(c, 404, "No notes found")
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes}) //todo: success msg func
}

func (h NoteHttpHandler) GetNote(c *fiber.Ctx) error {
	id := h.getNoteIdParam(c)
	if !h.repository.Exists(id) {
		return h.createFiberError(c, 404, "Note with id '"+id+"' not found.")
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": h.repository.FindById(id)})
}

func (h NoteHttpHandler) CreateNote(c *fiber.Ctx) error {
	note, err := h.parseNoteFromBody(c)
	if err != nil {
		return err
	}
	err = h.repository.Create(note)
	if err != nil {
		return h.createFiberErrorWithData(c, 500, "Could not create note", err)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func (h NoteHttpHandler) UpdateNote(c *fiber.Ctx) error {
	data, err := h.parseUpdateNoteFromBody(c)
	if err != nil {
		return err
	}
	note := h.toNote(h.getNoteIdParam(c), data)
	err = h.repository.Update(note)
	if err != nil {
		return h.createFiberError(c, 404, err.Error())
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func (h NoteHttpHandler) DeleteNote(c *fiber.Ctx) error {
	id := h.getNoteIdParam(c)
	if !h.repository.Exists(id) {
		return h.createFiberError(c, 404, "No note with id '"+id+"' present")
	}
	err := h.repository.DeleteById(id)
	if err != nil {
		return h.createFiberErrorWithData(c, 404, "Failed to delete note", err)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Note deleted."})
}

func (h NoteHttpHandler) createFiberError(c *fiber.Ctx, status int, msg string) error {
	return h.createFiberErrorWithData(c, status, msg, nil)
}

func (h NoteHttpHandler) createFiberErrorWithData(c *fiber.Ctx, status int, msg string, data error) error {
	return c.Status(status).JSON(fiber.Map{"status": "error", "message": msg, "data": data})
}

func (h NoteHttpHandler) parseNoteFromBody(c *fiber.Ctx) (model.Note, error) {
	var note model.Note
	err := c.BodyParser(&note)
	note.ID = uuid.New()
	if err != nil {
		err = h.createFiberError(c, 400, "Invalid input")
	}
	return note, err
}

func (h NoteHttpHandler) parseUpdateNoteFromBody(c *fiber.Ctx) (updateNote, error) {
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		err = h.createFiberError(c, 400, "Invalid input")
	}
	return updateNoteData, err
}

func (h NoteHttpHandler) toNote(id string, data updateNote) model.Note {
	var note model.Note
	note.ID = uuid.MustParse(id)
	note.Title = data.Title
	note.SubTitle = data.SubTitle
	note.Text = data.Text
	return note
}

func (h NoteHttpHandler) getNoteIdParam(c *fiber.Ctx) string {
	return c.Params("noteId")
}
