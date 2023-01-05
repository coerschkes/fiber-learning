package testutils

import (
	"net/http/httptest"
	"testing"

	"github.com/coerschkes/fiber-learning/internal"
	"github.com/coerschkes/fiber-learning/repository"
	"github.com/coerschkes/fiber-learning/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Description     string
	Route           string
	Method          string
	BodyObj         interface{}
	ExpectedCode    int
	ExpectedContent interface{}
	Repository      repository.NoteRepository
	Timeout         int
}

func RunTestCases[T any](t *testing.T, testCases []TestCase, assertions ...func(TestCase, internal.JsonResponse[T])) {
	for _, test := range testCases {
		app := fiber.New()
		router.SetupRoutes(app, test.Repository)

		req := httptest.NewRequest(test.Method, test.Route, internal.MarshalResponseBody(test.BodyObj))
		req.Header.Add("Content-Type", "application/json")

		resp, _ := app.Test(req, test.Timeout)
		assert.Equalf(t, test.ExpectedCode, resp.StatusCode, test.Description)
		parsedResponse := internal.UnmarshalResponseBody[T](resp.Body)
		for _, assertion := range assertions {
			assertion(test, parsedResponse)
		}
	}
}
