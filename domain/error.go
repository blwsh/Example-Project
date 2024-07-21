package domain

import "errors"

type ServiceError error

var (
	ErrBadRequest ServiceError = errors.New("bad request")
	ErrInternal   ServiceError = errors.New("internal error")
	ErrNotFound   ServiceError = errors.New("not found")
)

type Error struct {
	appErr error
	svcErr error
}

func NewError(svcErr ServiceError, appErr error) error {
	return &Error{
		appErr: appErr,
		svcErr: svcErr,
	}
}

func (e *Error) Error() string {
	return e.appErr.Error()
}

func (e *Error) SvcErr() error {
	return e.svcErr
}

func (e *Error) AppErr() error {
	return e.appErr
}
