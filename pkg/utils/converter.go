package utils

import (
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertNumericToDecimals(num *pgtype.Numeric, decimals uint8) (*big.Int, error) {
	if !num.Valid {
		return nil, fmt.Errorf("invalid numeric value")
	}
	if num.NaN {
		return nil, fmt.Errorf("NaN value not supported")
	}
	if num.InfinityModifier != 0 {
		return nil, fmt.Errorf("infinity not supported")
	}

	result := new(big.Int).Set(num.Int)

	expDiff := int64(num.Exp) + int64(decimals)

	if expDiff == 0 {
		return result, nil
	}

	if expDiff < 0 {
		divisor := new(big.Int).Exp(
			big.NewInt(10),
			big.NewInt(-expDiff),
			nil,
		)
		return result.Div(result, divisor), nil
	}

	// expDiff > 0
	multiplier := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(expDiff),
		nil,
	)
	result.Mul(result, multiplier)

	return result, nil
}

func ConvertDecimalsToWei(num *pgtype.Numeric) (*big.Int, error) {
	return ConvertNumericToDecimals(num, 18)
}

func ConvertFloat8ToDecimals(num pgtype.Float8, decimals uint64) (*big.Int, error) {
	if !num.Valid {
		return nil, fmt.Errorf("invalid float8 value")
	}

	bigFloat := new(big.Float).SetFloat64(num.Float64)

	multiplier := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(decimals)),
		nil,
	)

	bigFloat.Mul(bigFloat, new(big.Float).SetInt(multiplier))

	result := new(big.Int)
	bigFloat.Int(result)
	return result, nil
}
