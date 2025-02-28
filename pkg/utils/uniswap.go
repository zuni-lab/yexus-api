package utils

import (
	"math/big"
)

// Pre-calculate 2^96 as a big.Float for better precision
var POW96 = new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 96))

// CalculatePrice converts sqrtPriceX96 to a price with proper decimal adjustments
func CalculatePrice(sqrtPriceX96 *big.Int, token0Decimals, token1Decimals int32, token0IsUSD bool) *big.Float {
	// Set high precision to avoid rounding errors
	price := new(big.Float).SetPrec(256)

	// Convert sqrtPriceX96 to big.Float
	sqrtPrice := new(big.Float).SetPrec(256).SetInt(sqrtPriceX96)

	// Divide by 2^96
	sqrtPrice.Quo(sqrtPrice, POW96)

	// Square to get the price
	price.Mul(sqrtPrice, sqrtPrice)

	// Calculate decimal adjustment
	// Price = (token1/token0) * 10^(token0Decimals - token1Decimals)
	decimalAdjustment := token0Decimals - token1Decimals

	// If token0 is USD, invert the price
	if token0IsUSD {
		one := new(big.Float).SetPrec(256).SetInt64(1)
		price.Quo(one, price)
		// When inverting, we also need to invert the decimal adjustment
		decimalAdjustment = -decimalAdjustment
	}

	// Apply decimal adjustment using big.Float for precision
	if decimalAdjustment != 0 {
		// Create 10^|decimalAdjustment| using big.Int for exact representation
		powerTen := new(big.Int).Exp(
			big.NewInt(10),
			big.NewInt(int64(abs(decimalAdjustment))),
			nil,
		)

		// Convert to big.Float
		adjuster := new(big.Float).SetPrec(256).SetInt(powerTen)

		// Multiply or divide based on sign of decimalAdjustment
		if decimalAdjustment > 0 {
			price.Mul(price, adjuster)
		} else {
			price.Quo(price, adjuster)
		}
	}

	return price
}

// Helper function for absolute value of int32
func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}
