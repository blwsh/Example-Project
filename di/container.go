package di

import (
	"awesomeProject/internal/service/example"
	"awesomeProject/lib/logger"
)

type Container struct {
	Logger         logger.Logger
	ExampleService example.Service
}

type Option func(*Container)

func New(opts ...Option) *Container {
	c := &Container{
		Logger:         logger.New(false),
		ExampleService: example.New(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithLogger(l logger.Logger) Option {
	return func(c *Container) {
		c.Logger = l
	}
}

func WithExampleService(s example.Service) Option {
	return func(c *Container) {
		c.ExampleService = s
	}
}
