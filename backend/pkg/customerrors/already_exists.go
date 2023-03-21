package customerrors

import (
	"fmt"
)

type AlreadyExistsError struct {
	message string
}

func NewAlreadyExistsError(message string) *AlreadyExistsError {
	return &AlreadyExistsError{
		message: message,
	}
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s already exists", e.message)
}
