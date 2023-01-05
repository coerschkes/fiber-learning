package noteHandlerTest

import (
	"encoding/json"
	"io"

	"github.com/coerschkes/fiber-learning/internal"
)

type testResponseBodyParser[T any] struct{}

func (m *testResponseBodyParser[T]) unmarshalResponseBody(body io.ReadCloser) internal.JsonResponse[T] {
	var responseObj internal.JsonResponse[T]
	err := json.Unmarshal(m.readResponseBody(body), &responseObj)
	if err != nil {
		panic(err)
	}
	return responseObj
}

func (m *testResponseBodyParser[T]) readResponseBody(body io.ReadCloser) []byte {
	defer body.Close()
	readBody, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return readBody
}
