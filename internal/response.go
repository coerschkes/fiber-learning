package internal

type JsonResponse[T any] struct {
	Status uint
	Msg    string
	Data   T
}

func NewJsonResponse[T any](status uint, msg string, data T) JsonResponse[T] {
	return JsonResponse[T]{status, msg, data}
}

func NewEmptyJsonResponse[T any]() JsonResponse[T] {
	return JsonResponse[T]{}
}
