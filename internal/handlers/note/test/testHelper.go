package noteHandlerTest

import (
	"encoding/json"
	"io"

	"github.com/coerschkes/fiber-learning/internal"
	"github.com/coerschkes/fiber-learning/model"
	"github.com/coerschkes/fiber-learning/repository"
)

type testCase struct {
	description     string
	route           string
	method          string
	expectedCode    int
	expectedContent interface{}
	repository      repository.NoteRepository
	timeout         int
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
