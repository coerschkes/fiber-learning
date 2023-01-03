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

type ObjectExistsError struct {
	msg string
}

func NewObjectExistsError(msg string) *ObjectExistsError {
	return &ObjectExistsError{msg}
}

func (m *ObjectExistsError) Error() string {
	return m.msg
}
