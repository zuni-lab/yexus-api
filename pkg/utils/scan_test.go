package utils_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/zuni-lab/dexon-service/pkg/utils"
)

func TestScan(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"12345678901234567890.4444", "12345678901234567890.4444"},
	}

	for i, tc := range testCases {
		price := big.NewFloat(0.0)
		price.SetPrec(256)
		price.SetString(tc.input)

		numericPrice, err := utils.BigFloatToNumeric(price)
		if err != nil {
			t.Fatalf("Test case %d failed to convert: %v", i, err)
		}

		// For debugging
		fmt.Printf("Case %d:\n", i)
		fmt.Printf("  Input: %s\n", tc.input)
		fmt.Printf("  Numeric: Int=%v, Exp=%v\n", numericPrice.Int, numericPrice.Exp)

	}
}

func TestFloatTextLarge(t *testing.T) {
	price := big.NewInt(0)
	price.SetString("123456789012345678904444", 10)

	priceFloat := big.NewFloat(0.0)
	priceFloat.SetPrec(256)
	priceFloat.SetInt(price)
	priceFloat.Quo(priceFloat, big.NewFloat(math.Pow10(4)))

	fmt.Println(priceFloat.Text('f', -1))
}
