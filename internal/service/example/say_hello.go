package example

import (
	"errors"

	"awesomeProject/domain"
)

var (
	ErrNameRequired = errors.New("name is required")
)

func (s *service) SayHello(name string) (string, error) {
	if name == "" {
		return "", domain.NewError(domain.ErrBadRequest, ErrNameRequired)

	}

	return "Hello, " + name, nil
}
