package evm

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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
