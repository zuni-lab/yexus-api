package evm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5/pgtype"
)

func NormalizeAddress(address string) (common.Address, error) {
	address = strings.TrimPrefix(address, "0x")
	if len(address) != 40 {
		return common.Address{}, errors.New("invalid address length")
	}
	if !IsHexAddress(address) {
		return common.Address{}, errors.New("not a valid hex address")
	}

	return common.HexToAddress(address), nil
}

func NormalizeHex(h string) ([]byte, error) {
	h = strings.TrimPrefix(h, "0x")
	_h, err := hex.DecodeString(h)
	if err != nil {
		return nil, err
	}
	return _h, nil
}

func IsHex(hex string) bool {
	for _, c := range hex {
		if !isHexCharacter(byte(c)) {
			return false
		}
	}
	return true
}

func IsHexAddress(address string) bool {
	for _, c := range address {
		if !isHexCharacter(byte(c)) {
			return false
		}
	}
	return true
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
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

func ConvertFloat8ToWei(num pgtype.Float8) (*big.Int, error) {
	if !num.Valid {
		return nil, fmt.Errorf("invalid float8 value")
	}

	bigFloat := new(big.Float).SetFloat64(num.Float64)
	bigFloat.Mul(bigFloat, big.NewFloat(1e18))

	result := new(big.Int)
	bigFloat.Int(result)
	return result, nil
}
