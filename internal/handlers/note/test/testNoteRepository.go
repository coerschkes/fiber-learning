package noteHandler

import "github.com/coerschkes/fiber-learning/model"

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
