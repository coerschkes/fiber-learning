package internal

type JsonResponse struct {
	status uint
	msg    string
	data   interface{}
}

func NewJsonResponse(status uint, msg string, data interface{}) *JsonResponse {
	return &JsonResponse{status, msg, data}
}
