package noteHandlerTest

import (
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var testNote model.Note = model.Note{ID: uuid.New(), Title: "testTitle", SubTitle: "testsubtitle", Text: "testText"}
var emptyNotes []model.Note = []model.Note{}
var testNotes []model.Note = []model.Note{testNote}

func TestNoteHttpHandler(t *testing.T) {
	tests := []testCase{
		{
			description:     "get should return status 200 with empty content",
			route:           "/api/note/",
			method:          "GET",
			expectedCode:    200,
			expectedContent: emptyNotes,
			repository:      &TestNoteRepository{},
			timeout:         -1,
		},
		{
			description:     "get should return status 200 with one entry",
			route:           "/api/note/",
			method:          "GET",
			expectedCode:    200,
			expectedContent: testNotes,
			repository: &TestNoteRepository{findAllFn: func() []model.Note {
				return testNotes
			}},
			timeout: -1,
		},
	}

	for _, test := range tests {
		app := fiber.New()
		router.SetupRoutes(app, test.repository)

		req := httptest.NewRequest(test.method, test.route, nil)

		resp, _ := app.Test(req, test.timeout)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		parsed := mapJsonResponse(unmarshalResponseBody(resp.Body))
		assert.Equal(t, test.expectedContent, parsed.Data)
	}
}
