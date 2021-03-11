package main

import (
	"testing"
	"time"

	us "github.com/rpagliuca/go-uniswap-summary/pkg/unisummary"
	"github.com/stretchr/testify/assert"
)

func TestEnvToList(t *testing.T) {
	lpConfig := "DAI_WETH_LP,DAI,111.11,WETH,0.22,2021-01-30T09:30:00Z|DAI_USDC_LP,DAI,57.383,USDC,58.384,2021-01-30T09:30:00Z"
	liquidityProviderTokens := parseConfig(lpConfig)
	dateStart, _ := time.Parse(time.RFC3339, "2021-01-30T09:30:00Z")
	assert.Equal(t, []us.LiquidityProviderPosition{
		{us.TOKEN_DAI_WETH_LP, us.TOKEN_DAI, 111.11, us.TOKEN_WETH, 0.22, dateStart},
		{us.TOKEN_DAI_USDC_LP, us.TOKEN_DAI, 57.383, us.TOKEN_USDC, 58.384, dateStart},
	}, liquidityProviderTokens)
}
