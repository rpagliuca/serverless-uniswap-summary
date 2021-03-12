package main

import (
	us "github.com/rpagliuca/go-uniswap-summary/pkg/unisummary"
)

func run(etherscanApiKey, userWalletAddress string) interface{} {

	if etherscanApiKey == "" || userWalletAddress == "" {
		panic("Environment variables ETHERSCAN_API_KEY and USER_ADDRESS are required.")
	}

	req := us.NewUniswapSummaryRequest(
		etherscanApiKey,
		userWalletAddress,
		[]us.LiquidityProviderPosition{},
	)

	req.LiquidityProviderTokens = us.FromWalletAddress(req)

	resp := req.Do()

	return resp
}
