package internal

import (
	"bytes"
	"encoding/json"
	"io"
)

func UnmarshalResponseBody[T any](body io.ReadCloser) JsonResponse[T] {
	var responseObj JsonResponse[T]
	err := json.Unmarshal(readResponseBody(body), &responseObj)
	if err != nil {
		panic(err)
	}
	return responseObj
}

func MarshalResponseBody(obj interface{}) io.Reader {
	encodedObj, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	reader := bytes.NewReader(encodedObj)
	return reader
}

func readResponseBody(body io.ReadCloser) []byte {
	defer body.Close()
	readBody, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	return readBody
}
