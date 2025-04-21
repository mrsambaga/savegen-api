package model

import "fmt"

type ErrInvalidInput struct {
	Field  string
	Reason string
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf("Invalid input for %s: %s", e.Field, e.Reason)
}

type ErrNotFound struct {
    Resource string
}

func (e ErrNotFound) Error() string {
    return fmt.Sprintf("%s not found", e.Resource)
}
