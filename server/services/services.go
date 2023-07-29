package services

import (
	"errors"
	"fmt"
)

type ServiceError struct {
	service string
	method  string
	err     error
}

func (ser *ServiceError) Error() string {
	return fmt.Sprintf("(Service) --- %s --- %s --- %v", ser.service, ser.method, ser.err)
}

func (ser *ServiceError) Wrap(err error) error {
	if err == nil {
		return nil
	}

	ser.err = err
	return ser
}
func (ser *ServiceError) From(text string) error {
	ser.err = errors.New(text)
	return ser
}

func newServiceError(service string, method string) *ServiceError {
	return &ServiceError{
		service: service,
		method:  method,
	}
}
