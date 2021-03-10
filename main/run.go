package main

import (
	us "github.com/rpagliuca/go-uniswap-summary/pkg/unisummary"
)

func run(etherscanApiKey, userWalletAddress, poolPositions string) interface{} {

	if etherscanApiKey == "" || userWalletAddress == "" || poolPositions == "" {
		panic("Environment variables ETHERSCAN_API_KEY and USER_ADDRESS are required.")
	}

	liquidityProviderTokens := parseConfig(poolPositions)

	req := us.NewUniswapSummaryRequest(
		etherscanApiKey,
		userWalletAddress,
		liquidityProviderTokens,
	)

	resp := req.Do()

	return resp
}
