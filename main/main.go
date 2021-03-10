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
var userWalletAddress = os.Getenv("USER_WALLET_ADDRESS")
var liquidityProviderPoolPositions = os.Getenv("LIQUIDITY_PROVIDER_POOL_POSITIONS")

func Handler(ctx context.Context, request Request) (Response, error) {

	data := run(etherscanApiKey, userWalletAddress, liquidityProviderPoolPositions)

	// Return JSON response
	return createSuccessResponse(data)
}

// Allow mocking
var lambdaStart = lambda.Start

func main() {
	lambdaStart(Handler)
}
