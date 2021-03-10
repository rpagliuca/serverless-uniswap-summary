package main

import (
	"testing"

	us "github.com/rpagliuca/go-uniswap-summary/pkg/unisummary"
	"github.com/stretchr/testify/assert"
)

func TestEnvToList(t *testing.T) {
	lpConfig := "DAI_WETH_LP,DAI,111.11,WETH,0.22|DAI_USDC_LP,DAI,57.383,USDC,58.384"
	liquidityProviderTokens := parseConfig(lpConfig)
	assert.Equal(t, []us.LiquidityProviderToken{
		{us.TOKEN_DAI_WETH_LP, us.TOKEN_DAI, 111.11, us.TOKEN_WETH, 0.22},
		{us.TOKEN_DAI_USDC_LP, us.TOKEN_DAI, 57.383, us.TOKEN_USDC, 58.384},
	}, liquidityProviderTokens)
}
