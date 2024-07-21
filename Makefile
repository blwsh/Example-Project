gen:
	go generate ./...

build:
	make build-Rest build-Authorizer

build-Rest:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/rest/bootstrap cmd/rest/main.go

build-Authorizer:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/authorizer/bootstrap cmd/authorizer/main.go
