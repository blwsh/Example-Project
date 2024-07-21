package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"

	"github.com/aws/jsii-runtime-go"
)

type InfraStackProps struct {
	awscdk.StackProps
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//authFn := awslambda.NewFunction(stack, jsii.String("AuthFn"), &awslambda.FunctionProps{
	//	Handler: jsii.String("bootstrap"),
	//	Runtime: awslambda.Runtime_PROVIDED_AL2023(),
	//	Code:    awslambda.Code_FromAsset(jsii.String("../bin/authorizer"), &awss3assets.AssetOptions{}),
	//})
	//
	//auth := awsapigateway.NewTokenAuthorizer(stack, jsii.String("Authorizer"), &awsapigateway.TokenAuthorizerProps{
	//	AuthorizerName: jsii.String("Authorizer"),
	//	Handler:        authFn,
	//})

	var restFn awslambda.IFunction = awslambda.NewFunction(stack, jsii.String("Lambda"), &awslambda.FunctionProps{
		Handler: jsii.String("bootstrap"),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.Code_FromAsset(jsii.String("../bin/rest"), &awss3assets.AssetOptions{}),
	})

	gw := awsapigateway.NewRestApi(stack, jsii.String("Api"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("Api"),
	})

	gw.Root().
		AddResource(jsii.String("{proxy+}"), nil).
		AddMethod(jsii.String("ANY"), awsapigateway.NewLambdaIntegration(restFn, &awsapigateway.LambdaIntegrationOptions{
			Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
		}), &awsapigateway.MethodOptions{
			//Authorizer: auth,
		})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewInfraStack(app, "InfraStack", &InfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
