package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

var etherscanApiKey = os.Getenv("ETHERSCAN_API_KEY")

func Handler(ctx context.Context, request Request) (Response, error) {

	userWalletAddress := request.PathParameters["walletAddress"]

	data := run(etherscanApiKey, userWalletAddress)

	// Return JSON response
	return createSuccessResponse(data)
}

// Allow mocking
var lambdaStart = lambda.Start

func main() {
	lambdaStart(Handler)
}
