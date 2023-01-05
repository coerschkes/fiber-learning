package noteHandlerTestUtils

import (
	"github.com/coerschkes/fiber-learning/model"
	"github.com/google/uuid"
)

var TestNote model.Note = model.Note{ID: uuid.New(), Title: "testTitle", SubTitle: "testsubtitle", Text: "testText"}
var EmptyNote model.Note = model.Note{}
var EmptyNotes []model.Note = []model.Note{}
var TestNotes []model.Note = []model.Note{TestNote}

type TestNoteRepository struct {
	FindAllFn    func() []model.Note
	FindByIdFn   func(id string) model.Note
	CreateFn     func(note model.Note) error
	UpdateFn     func(note model.Note) error
	DeleteByIdFn func(id string) error
	ExistsFn     func(id string) bool
}

func (m *TestNoteRepository) FindAll() []model.Note {
	if m.FindAllFn != nil {
		return m.FindAllFn()
	}
	return make([]model.Note, 0)
}

func (m *TestNoteRepository) FindById(id string) model.Note {
	if m.FindByIdFn != nil {
		return m.FindByIdFn(id)
	}
	return model.Note{}
}

func (m *TestNoteRepository) Create(note model.Note) error {
	if m.CreateFn != nil {
		return m.CreateFn(note)
	}
	return nil
}

func (m *TestNoteRepository) Update(note model.Note) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(note)
	}
	return nil
}

func (m *TestNoteRepository) DeleteById(id string) error {
	if m.DeleteByIdFn != nil {
		return m.DeleteByIdFn(id)
	}
	return nil
}

func (m *TestNoteRepository) Exists(id string) bool {
	if m.ExistsFn != nil {
		return m.ExistsFn(id)
	}
	return false
}
