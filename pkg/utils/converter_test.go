package utils_test

import (
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zuni-lab/yexus-api/pkg/utils"
)

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
