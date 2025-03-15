package utils

import (
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

// BigFloatToNumeric converts a big.Float to pgtype.Numeric
// preserving full precision including decimals
func BigFloatToNumeric(f *big.Float) (pgtype.Numeric, error) {
	str := f.Text('f', -1)
	var numeric pgtype.Numeric
	err := numeric.Scan(str)
	if err != nil {
		return pgtype.Numeric{}, err
	}
	return numeric, nil
}

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

// ScanValue is a generic function that converts a string value to a pgtype type
func ScanNumericValue(value string) (*pgtype.Numeric, error) {
	var numeric pgtype.Numeric
	if err := numeric.Scan(value); err != nil {
		return nil, fmt.Errorf("failed to scan value %q: %w", value, err)
	}

	return &numeric, nil
}

func ScanStringValue(value string) (*pgtype.Text, error) {
	var text pgtype.Text
	if err := text.Scan(value); err != nil {
		return nil, fmt.Errorf("failed to scan value %q: %w", value, err)
	}

	return &text, nil
}

func ScanBoolValue(value string) (*pgtype.Bool, error) {
	var bool pgtype.Bool
	if err := bool.Scan(value); err != nil {
		return nil, fmt.Errorf("failed to scan value %q: %w", value, err)
	}

	return &bool, nil
}
