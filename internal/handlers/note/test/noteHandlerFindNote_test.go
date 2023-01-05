package noteHandlerTest

import (
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNoteHttpHandlerFindNote(t *testing.T) {
	tests := []testCase{
		{
			description:     "get should return status 404 if note does not exist with empty note entry",
			route:           "/api/note/1",
			method:          "GET",
			expectedCode:    404,
			expectedContent: emptyNote,
			repository:      &TestNoteRepository{},
			timeout:         -1,
		},
		{
			description:     "get should return status 200 note with id " + testNote.ID.String() + " if it exists",
			route:           "/api/note/" + testNote.ID.String(),
			method:          "GET",
			expectedCode:    200,
			expectedContent: testNote,
			repository: &TestNoteRepository{
				findByIdFn: func(id string) model.Note {
					return testNote
				},
				existsFn: func(id string) bool {
					return true
				},
			},
			timeout: -1,
		},
	}

	for _, test := range tests {
		app := fiber.New()
		router.SetupRoutes(app, test.repository)

		req := httptest.NewRequest(test.method, test.route, test.body)

		resp, _ := app.Test(req, test.timeout)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		responseBodyParser := testResponseBodyParser[model.Note]{}
		parsedResponse := responseBodyParser.unmarshalResponseBody(resp.Body)
		assert.Equal(t, test.expectedContent.(model.Note), parsedResponse.Data)
	}
}
