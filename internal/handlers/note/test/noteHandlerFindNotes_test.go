package noteHandler_test

import (
	"testing"

	. "github.com/coerschkes/fiber-learning/internal"
	. "github.com/coerschkes/fiber-learning/internal/handlers/note/test/noteHandlerTestUtils"
	. "github.com/coerschkes/fiber-learning/internal/testutils"
	. "github.com/coerschkes/fiber-learning/model"
	"github.com/stretchr/testify/assert"
)

func TestNoteHttpHandlerFindNotes(t *testing.T) {
	tests := []TestCase{
		{
			Description:     "get should return status 200 with empty notes slice",
			Route:           "/api/note/",
			Method:          "GET",
			ExpectedCode:    200,
			ExpectedContent: EmptyNotes,
			Repository:      &TestNoteRepository{},
			Timeout:         -1,
		},
		{
			Description:     "get should return status 200 with notes slice with single entry",
			Route:           "/api/note/",
			Method:          "GET",
			ExpectedCode:    200,
			ExpectedContent: TestNotes,
			Repository: &TestNoteRepository{FindAllFn: func() []Note {
				return TestNotes
			}},
			Timeout: -1,
		},
	}

	RunTestCases(t, tests, func(tc TestCase, jsonResponse JsonResponse[[]Note]) {
		assert.Len(t, tc.ExpectedContent.([]Note), len(jsonResponse.Data))
		assert.Equal(t, tc.ExpectedContent.([]Note), jsonResponse.Data)
	})
}
