package noteHandler

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/internal"
	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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

	for _, test := range tests {
		app := fiber.New()
		router.SetupRoutes(app, test.repository)

		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		parsed := mapJsonResponse(unmarshalResponseBody(resp.Body))
		assert.Equal(t, test.expectedContent, parsed.Data)
	}
}

func unmarshalResponseBody(body io.ReadCloser) internal.JsonResponse[[]model.Note] {
	var responseObj internal.JsonResponse[[]model.Note]
	err := json.Unmarshal(readResponseBody(body), &responseObj)
	if err != nil {
		panic(err)
	}
	return responseObj
}

func mapJsonResponse(jsonResponse internal.JsonResponse[[]model.Note]) internal.JsonResponse[[]model.Note] {
	if jsonResponse.Data == nil {
		return internal.NewJsonResponse(999, "", []model.Note{})
	}
	return jsonResponse
}

func readResponseBody(body io.ReadCloser) []byte {
	defer body.Close()
	readBody, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return readBody
}
