package errors

type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}
