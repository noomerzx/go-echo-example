package utils

// NotFoundError struct
type NotFoundError struct {
	message string
}

// NewNotFoundError func
func NewNotFoundError() *NotFoundError {
	return &NotFoundError{message: "404"}
}

func (e NotFoundError) Error() string {
	return e.message
}
