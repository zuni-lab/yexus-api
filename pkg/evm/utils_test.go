package evm_test

import (
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestAbi(t *testing.T) {
	abiJson, err := os.ReadFile("../../assets/DexonABI.json")
	if err != nil {
		t.Errorf("Error reading abi.json: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		t.Errorf("Error parsing abi.json: %v", err)
	}

	_ = contractAbi
}
