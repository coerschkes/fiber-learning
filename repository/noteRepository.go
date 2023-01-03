package repository

import (
	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository/err"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository interface {
	FindAll() []model.Note
	FindById(id string) model.Note
	Create(note model.Note) error
	Update(note model.Note) error
	DeleteById(id string) error
	Exists(id string) bool
}

type PostgresNoteRepository struct {
	database *gorm.DB
}

func NewPostgresNoteRepository(database *gorm.DB) *PostgresNoteRepository {
	return &PostgresNoteRepository{database}
}

func (h PostgresNoteRepository) FindAll() []model.Note {
	var notes []model.Note

	h.database.Find(&notes)

	return notes
}

func (h PostgresNoteRepository) FindById(id string) model.Note {
	var note model.Note
	h.database.Find(&note, "id = ?", id)
	return note
}

func (h PostgresNoteRepository) Create(note model.Note) error {
	if !h.Exists(note.ID.String()) {
		return createIdError(note.ID.String())
	}
	return h.database.Create(&note).Error
}

func (h PostgresNoteRepository) DeleteById(id string) error {
	if !h.Exists(id) {
		return createIdError(id)
	}
	note := h.FindById(id)

	return h.database.Delete(&note, "id = ?", id).Error
}

func (h PostgresNoteRepository) Update(note model.Note) error {
	if !h.Exists(note.ID.String()) {
		return createIdError(note.ID.String())
	}
	return h.database.Save(&note).Error
}

func (h PostgresNoteRepository) Exists(id string) bool {
	return h.FindById(id).ID == uuid.Nil
}

func createIdError(id string) error {
	return err.NewObjectNotFoundError("Object with id '" + id + "' does not exist.")
}
