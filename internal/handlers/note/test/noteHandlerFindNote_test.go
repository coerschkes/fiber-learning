package noteHandler_test

import (
	"testing"

	. "github.com/coerschkes/fiber-learning/internal"
	. "github.com/coerschkes/fiber-learning/internal/handlers/note/test/noteHandlerTestUtils"
	. "github.com/coerschkes/fiber-learning/internal/testutils"
	. "github.com/coerschkes/fiber-learning/model"
	"github.com/stretchr/testify/assert"
)

func TestNoteHttpHandlerFindNote(t *testing.T) {
	tests := []TestCase{
		{
			Description:     "get should return status 404 if note does not exist with empty note entry",
			Route:           "/api/note/1",
			Method:          "GET",
			ExpectedCode:    404,
			ExpectedContent: EmptyNote,
			Repository:      &TestNoteRepository{},
			Timeout:         -1,
		},
		{
			Description:     "get should return status 200 note with id " + TestNote.ID.String() + " if it exists",
			Route:           "/api/note/" + TestNote.ID.String(),
			Method:          "GET",
			ExpectedCode:    200,
			ExpectedContent: TestNote,
			Repository: &TestNoteRepository{
				FindByIdFn: func(id string) Note {
					return TestNote
				},
				ExistsFn: func(id string) bool {
					return true
				},
			},
			Timeout: -1,
		},
	}
	RunTestCases(t, tests, func(tc TestCase, jsonResponse JsonResponse[Note]) {
		assert.Equal(t, tc.ExpectedContent.(Note), jsonResponse.Data)
	})
}
