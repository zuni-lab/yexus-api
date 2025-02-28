package utils_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/zuni-lab/dexon-service/pkg/utils"
)

var testCases = []struct {
	sqrtPriceX96   string
	token0Decimals int32
	token1Decimals int32
	token0IsUSD    bool
}{
	{
		sqrtPriceX96:   "1685552085383287813036109446383368",
		token0Decimals: 6,
		token1Decimals: 18,
		token0IsUSD:    true,
	},
	{
		sqrtPriceX96:   "3715267771548011387318332",
		token0Decimals: 18,
		token1Decimals: 6,
		token0IsUSD:    false,
	},
}

func TestCalculatePrice(t *testing.T) {
	for _, testCase := range testCases {
		sqrtPriceX96Int, ok := big.NewInt(0).SetString(testCase.sqrtPriceX96, 10)
		if !ok {
			t.Fatalf("failed to convert sqrtPriceX96 to big.Int")
		}

		price := utils.CalculatePrice(sqrtPriceX96Int, testCase.token0Decimals, testCase.token1Decimals, testCase.token0IsUSD)
		fmt.Println("price", price)
		numericPrice, err := utils.BigFloatToNumeric(price)
		if err != nil {
			t.Fatalf("failed to convert price to numeric: %v", err)
		}
		fmt.Println("numericPrice", numericPrice)
	}
}
