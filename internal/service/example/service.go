package example

//go:generate go run go.uber.org/mock/mockgen@latest -destination=mock_service.go -package=example . Service

type Service interface {
	SayHello(name string) (string, error)
}

type service struct{}

func New() Service {
	return &service{}
}
