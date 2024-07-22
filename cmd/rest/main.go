package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"awesomeProject/cmd/rest/api"
	"awesomeProject/di"
)

func main() {
	// initialize outside of lambda.Start to avoid reinitializing on every invocation
	var (
		container = di.New()
		server    = api.NewServer(container)
		adapter   = httpadapter.New(api.HandlerFromMux(server, http.NewServeMux()))
	)

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		container.Logger.With("request", req).Info("request received")

		return adapter.ProxyWithContext(ctx, req)
	})
}
