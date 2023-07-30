package models

import (
	"errors"
	"fmt"
)

type Model interface {
	Validate() error
	JSON() (string, error)
}

type ModelError struct {
	model string
	err   error
}

func (e *ModelError) Error() string {
	return fmt.Sprintf("(Model) --- %s --- %v", e.model, e.err)
}

func (e *ModelError) Wrap(err error) error {
	if err == nil {
		return nil
	}

	e.err = err
	return e
}
func (e *ModelError) From(text string) error {
	e.err = errors.New(text)
	return e
}

func newModelError(model string) *ModelError {
	return &ModelError{
		model: model,
	}
}
