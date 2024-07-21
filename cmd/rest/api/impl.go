package api

import (
	"awesomeProject/di"
)

type Server struct {
	container *di.Container
}

func NewServer(container *di.Container) Server {
	return Server{
		container: container,
	}
}
