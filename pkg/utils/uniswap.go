package utils

import (
	"math"
	"math/big"
)

var POW96 = new(big.Float).SetFloat64(math.Pow(2, 96))

func CalculatePrice(sqrtPriceX96 *big.Int, token0Decimals, token1Decimals uint8, token0IsUSD bool) *big.Float {
	// Convert sqrtPriceX96 to price
	// Formula: price = (sqrtPriceX96/2^96)^2
	// Create big.Float for precise decimal calculations
	price := new(big.Float).SetPrec(256)

	// Convert sqrtPriceX96 to big.Float
	sqrtPrice := new(big.Float).SetInt(sqrtPriceX96)

	// Divide sqrtPrice by 2^96
	sqrtPrice.Quo(sqrtPrice, POW96)

	// Square the result to get the actual price
	price.Mul(sqrtPrice, sqrtPrice)

	// Handle decimal adjustments
	decimals := token1Decimals - token0Decimals
	if token0IsUSD {
		// If token0 is USD, we need to invert the price
		one := new(big.Float).SetFloat64(1.0)
		price.Quo(one, price)
		decimals = -decimals
	}

	// Adjust for decimal differences
	if decimals != 0 {
		adjuster := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))
		price.Mul(price, adjuster)
	}

	return price
}
