package err

type ObjectNotFoundError struct {
	msg string
}

func NewObjectNotFoundError(msg string) *ObjectNotFoundError {
	return &ObjectNotFoundError{msg}
}

func (m *ObjectNotFoundError) Error() string {
	return m.msg
}
