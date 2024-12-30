package serviceerrors

import "errors"

var (
	ErrValidationFailed = errors.New("validation failed")
)

type ServiceError error

func New(err error) ServiceError {
	return ServiceError(err)
}
