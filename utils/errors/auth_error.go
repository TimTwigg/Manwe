package error_utils

type AuthErrorType interface {
	error
	Error() string
}

type AuthError struct {
	error
	Message string
}

func (e AuthError) Error() string {
	return "AuthError: " + e.Message
}
