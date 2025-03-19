package utils_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zuni-lab/yexus-api/pkg/utils"
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

func TestNumericConversion(t *testing.T) {
	testCases := []struct {
		name     string
		input    pgtype.Numeric
		decimals uint8
		expected string
	}{
		{
			name: "USDC",
			input: pgtype.Numeric{
				Int: func() *big.Int {
					n, _ := new(big.Int).SetString("19005000000000000000000", 10)
					return n
				}(),
				Exp:   -19,
				Valid: true,
			},
			decimals: 18,
			expected: "1900500000000000000000", // 1900.5 * 10^18
		},
		{
			name: "WBTC",
			input: pgtype.Numeric{
				Int: func() *big.Int {
					n, _ := new(big.Int).SetString("19005000000000000000000", 10)
					return n
				}(),
				Exp:   -17,
				Valid: true,
			},
			decimals: 18,
			expected: "190050000000000000000000",
		},
		{
			name: "WSOL",
			input: pgtype.Numeric{
				Int: func() *big.Int {
					n, _ := new(big.Int).SetString("5", 10)
					return n
				}(),
				Exp:   5,
				Valid: true,
			},
			decimals: 8,
			expected: "50000000000000", // 5*10^13
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.ConvertNumericToDecimals(&tc.input, tc.decimals)
			if err != nil {
				t.Fatalf("ConvertNumericToDecimals failed: %v", err)
			}

			t.Logf("result: %+v\n", result)

			if result.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

func TestFloatToWei(t *testing.T) {
	float := pgtype.Float8{
		Float64: 0.1,
		Valid:   true,
	}

	result, err := utils.ConvertFloat8ToDecimals(float, 6)
	if err != nil {
		t.Fatalf("ConvertFloat8ToWei failed: %v", err)
	}

	t.Logf("result: %+v\n", result)
}
