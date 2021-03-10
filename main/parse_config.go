package main

import (
	"fmt"
	"strconv"
	"strings"

	us "github.com/rpagliuca/go-uniswap-summary/pkg/unisummary"
)

var allTokens = []us.Token{
	us.TOKEN_DAI_WETH_LP,
	us.TOKEN_DAI_USDC_LP,
	us.TOKEN_DAI,
	us.TOKEN_WETH,
	us.TOKEN_USDC,
}

func parseConfig(c string) []us.LiquidityProviderToken {
	items := strings.Split(c, "|")
	tokens := []us.LiquidityProviderToken{}
	for _, item := range items {
		tokens = append(tokens, parseItem(item))
	}
	return tokens
}

func parseItem(item string) us.LiquidityProviderToken {
	values := strings.Split(item, ",")
	if len(values) != 5 {
		panic("Invalid configuration. Each LP item should contain 5 values.")
	}
	stringStruct := StringStruct{values[0], values[1], values[2], values[3], values[4]}
	typedStruct := convertToTyped(stringStruct)
	return typedStruct
}

func convertToTyped(s StringStruct) us.LiquidityProviderToken {
	lpToken, err := findToken(s.LpTokenId)
	handleError(err)
	token1, err := findToken(s.Token1Id)
	handleError(err)
	token2, err := findToken(s.Token2Id)
	handleError(err)
	token1Quantity, err := strconv.ParseFloat(s.Token1InitialQuantity, 64)
	handleError(err)
	token2Quantity, err := strconv.ParseFloat(s.Token2InitialQuantity, 64)
	handleError(err)
	return us.LiquidityProviderToken{
		*lpToken,
		*token1,
		token1Quantity,
		*token2,
		token2Quantity,
	}
}

func findToken(tokenId string) (*us.Token, error) {
	for _, t := range allTokens {
		if tokenId == t.Id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("Token %s not found\n", tokenId)
}

type StringStruct struct {
	LpTokenId             string
	Token1Id              string
	Token1InitialQuantity string
	Token2Id              string
	Token2InitialQuantity string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
