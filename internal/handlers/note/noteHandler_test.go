package noteHandler_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestNoteHttpHandler_FindNotes(t *testing.T) {
	tests := []struct {
		description     string
		route           string
		expectedCode    int
		expectedContent interface{}
		repository      repository.NoteRepository
	}{
		{
			description:     "get should return status 200 with empty content",
			route:           "/api/note/",
			expectedCode:    200,
			expectedContent: []model.Note{},
			repository:      &TestNoteRepository{},
		},
		{
			description:     "get should return status 200 with one entry",
			route:           "/api/note/",
			expectedCode:    200,
			expectedContent: []model.Note{{Title: "test"}},
			repository: &TestNoteRepository{findAllFn: func() []model.Note {
				var notes []model.Note
				notes = append(notes, model.Note{Title: "test"})
				return notes
			}},
		},
	}

	// note := router.Group("/note")
	// note.Post("/", n.CreateNote)
	// note.Get("/", n.FindNotes)
	// note.Get("/:noteId", n.FindNote)
	// note.Put("/:noteId", n.UpdateNote)
	// note.Delete("/:noteId", n.DeleteNote)

	for _, test := range tests {
		app := fiber.New()
		router.SetupRoutes(app, test.repository)

		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		parsed := parseNotesFromResponseBody(resp.Body)
		assert.Equal(t, test.expectedContent, parsed)
	}
}

// todo: refactor
func parseNotesFromResponseBody(body io.ReadCloser) []model.Note {
	defer body.Close()
	readBody, _ := io.ReadAll(body)
	dataAsString := gjson.Get(
		string(readBody),
		"data",
	).Raw

	var notes []model.Note
	err := json.Unmarshal([]byte(dataAsString), &notes)
	if dataAsString == "null" || err != nil {
		return []model.Note{}
	}
	return notes
}

type TestNoteRepository struct {
	findAllFn    func() []model.Note
	findByIdFn   func(id string) model.Note
	createFn     func(note model.Note) error
	updateFn     func(note model.Note) error
	deleteByIdFn func(id string) error
	existsFn     func(id string) bool
}

func (m *TestNoteRepository) FindAll() []model.Note {
	if m.findAllFn != nil {
		return m.findAllFn()
	}
	return []model.Note{}
}

func (m *TestNoteRepository) FindById(id string) model.Note {
	if m.findAllFn != nil {
		return m.findByIdFn(id)
	}
	return model.Note{}
}

func (m *TestNoteRepository) Create(note model.Note) error {
	if m.createFn != nil {
		return m.createFn(note)
	}
	return nil
}

func (m *TestNoteRepository) Update(note model.Note) error {
	if m.updateFn != nil {
		return m.updateFn(note)
	}
	return nil
}

func (m *TestNoteRepository) DeleteById(id string) error {
	if m.deleteByIdFn != nil {
		return m.deleteByIdFn(id)
	}
	return nil
}

func (m *TestNoteRepository) Exists(id string) bool {
	if m.existsFn != nil {
		return m.existsFn(id)
	}
	return false
}
