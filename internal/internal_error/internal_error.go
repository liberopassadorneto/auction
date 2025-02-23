package internal_error

type InternalError struct {
	Message string
	Error   string
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Error:   "internal_server_error",
	}
}
