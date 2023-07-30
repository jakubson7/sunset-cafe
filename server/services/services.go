package services

import (
	"errors"
	"fmt"
	"log"
)

type ServiceError struct {
	service string
	method  string
	err     error
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("(Service) --- %s --- %s --- %v", e.service, e.method, e.err)
}

func (e *ServiceError) Wrap(err error) error {
	if err == nil {
		return nil
	}

	e.err = err
	return e
}
func (e *ServiceError) From(text string) error {
	e.err = errors.New(text)
	return e
}
func (e *ServiceError) Fatal(err error) {
	log.Fatal(e.Wrap(err))
}

func newServiceError(service string, method string) *ServiceError {
	return &ServiceError{
		service: service,
		method:  method,
	}
}
