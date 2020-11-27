package utils

// ConflictError struct
type ConflictError struct {
	message string
}

// NewConflictError func
func NewConflictError() *ConflictError {
	return &ConflictError{message: "409"}
}

func (e ConflictError) Error() string {
	return e.message
}
