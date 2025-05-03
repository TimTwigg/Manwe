package error_utils

type ParseErrorType interface {
	error
	Error() string
}

type ParseError struct {
	error
	Message string
}

func (e ParseError) Error() string {
	return "ParseError: " + e.Message
}
