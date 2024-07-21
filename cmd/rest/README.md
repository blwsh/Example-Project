To update the API:

1. Update openapi.yaml 
2. Run `cmd/rest/api/generate.go` (or `make gen`)
3. This will update `ServerInterface` which `cmd/rest/api/server.go` implements. Implement the newly generated method(s) in `cmd/rest/api/server.go`
