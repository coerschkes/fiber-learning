package noteHandler

import (
	"strconv"

	"github.com/coerschkes/fiber-learning/internal"
	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type NoteHandler interface {
	FindNotes(c *fiber.Ctx) error
	FindNote(c *fiber.Ctx) error
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

func (h NoteHttpHandler) FindNotes(c *fiber.Ctx) error {
	notes := h.repository.FindAll()
	if len(notes) == 0 {
		return h.createJSONResponse(c, fiber.StatusNotFound, "no notes found ", nil)
	}
	return h.createJSONResponse(c, fiber.StatusOK, "found '"+strconv.Itoa(len(notes))+"' notes", notes)
}

func (h NoteHttpHandler) FindNote(c *fiber.Ctx) error {
	id := h.getNoteIdParam(c)
	if !h.repository.Exists(id) {
		return h.createJSONResponse(c, fiber.StatusNotFound, "note with id '"+id+"' not found.", nil)
	}
	return h.createJSONResponse(c, fiber.StatusOK, "note with id '"+id+"' found", h.repository.FindById(id))
}

func (h NoteHttpHandler) CreateNote(c *fiber.Ctx) error {
	note, err := h.parseNoteFromBody(c)
	if err != nil {
		return h.createJSONResponse(c, fiber.StatusBadRequest, "invalid input", err)
	}
	err = h.repository.Create(note)
	if err != nil {
		return h.createJSONResponse(c, fiber.StatusInternalServerError, "could not create note", err)
	}
	return h.createJSONResponse(c, fiber.StatusCreated, "note created", note)
}

func (h NoteHttpHandler) UpdateNote(c *fiber.Ctx) error {
	data, err := h.parseUpdateNoteFromBody(c)
	if err != nil {
		return h.createJSONResponse(c, fiber.StatusBadRequest, "invalid input", err)
	}
	note := h.convertToNote(h.getNoteIdParam(c), data)
	err = h.repository.Update(note)
	if err != nil {
		return h.createJSONResponse(c, fiber.StatusNotFound, "note not found", err)
	}
	return h.createJSONResponse(c, fiber.StatusNoContent, "note updated", note)
}

func (h NoteHttpHandler) DeleteNote(c *fiber.Ctx) error {
	id := h.getNoteIdParam(c)
	if !h.repository.Exists(id) {
		h.createJSONResponse(c, fiber.StatusNotFound, "note with id '"+id+"' not found", nil)
	}
	err := h.repository.DeleteById(id)
	if err != nil {
		return h.createJSONResponse(c, fiber.StatusNotFound, "error deleting note with id '"+id+"'", err)
	}
	return h.createJSONResponse(c, fiber.StatusNoContent, "note with id '"+id+"' deleted", nil)
}

func (h NoteHttpHandler) parseNoteFromBody(c *fiber.Ctx) (model.Note, error) {
	var note model.Note
	err := c.BodyParser(&note)
	note.ID = uuid.New()
	return note, err
}

func (h NoteHttpHandler) parseUpdateNoteFromBody(c *fiber.Ctx) (updateNote, error) {
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	return updateNoteData, err
}

func (h NoteHttpHandler) convertToNote(id string, data updateNote) model.Note {
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

func (h NoteHttpHandler) createJSONResponse(c *fiber.Ctx, status uint, msg string, data interface{}) error {
	return c.Status(int(status)).JSON(internal.NewJsonResponse(status, msg, data))
}
