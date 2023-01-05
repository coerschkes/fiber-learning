package noteHandlerTest

import (
	"io"

	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/google/uuid"
)

type testCase struct {
	description     string
	route           string
	method          string
	body            io.Reader
	expectedCode    int
	expectedContent interface{}
	repository      repository.NoteRepository
	timeout         int
}

var testNote model.Note = model.Note{ID: uuid.New(), Title: "testTitle", SubTitle: "testsubtitle", Text: "testText"}
var emptyNote model.Note = model.Note{}
var emptyNotes []model.Note = []model.Note{}
var testNotes []model.Note = []model.Note{testNote}
