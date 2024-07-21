package usecase

import (
	"awesomeProject/internal/service/example"
)

type SayHello struct {
	exampleService example.Service
}

func NewSayHello(exampleService example.Service) *SayHello {
	return &SayHello{exampleService: exampleService}
}

func (s *SayHello) Execute(name string) (string, error) {
	hello, err := s.exampleService.SayHello(name)
	if err != nil {
		return "", err
	}

	return hello, nil
}
