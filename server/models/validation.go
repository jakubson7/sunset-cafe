package models

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	model string
	err   error
}

func (ver *ValidationError) Error() string {
	return fmt.Sprintf("(Validation) --- %s --- %v", ver.model, ver.err)
}

func (ver *ValidationError) Wrap(err error) error {
	if err == nil {
		return nil
	}

	ver.err = err
	return ver
}
func (ver *ValidationError) From(text string) error {
	ver.err = errors.New(text)
	return ver
}

func newValidationError(model string) *ValidationError {
	return &ValidationError{
		model: model,
	}
}

func trim(str string) string {
	return strings.Trim(str, " \n")
}
