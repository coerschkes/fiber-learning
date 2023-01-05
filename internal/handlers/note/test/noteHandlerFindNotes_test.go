package noteHandlerTest

import (
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNoteHttpHandlerFindNotes(t *testing.T) {
	tests := []testCase{
		{
			description:     "get should return status 200 with empty notes slice",
			route:           "/api/note/",
			method:          "GET",
			expectedCode:    200,
			expectedContent: emptyNotes,
			repository:      &TestNoteRepository{},
			timeout:         -1,
		},
		{
			description:     "get should return status 200 with notes slice with single entry",
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
		responseBodyParser := testResponseBodyParser[[]model.Note]{}
		parsedResponse := responseBodyParser.unmarshalResponseBody(resp.Body)
		assert.Len(t, test.expectedContent.([]model.Note), len(parsedResponse.Data))
		assert.Equal(t, test.expectedContent.([]model.Note), parsedResponse.Data)
	}
}
