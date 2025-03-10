package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zuni-lab/dexon-service/config"
)

func CreateTxData(contract *DexonTransactor, method string, args ...interface{}) ([]byte, error) {
	parsed, err := DexonMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %w", err)
	}

	data, err := parsed.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data: %w", err)
	}

	return data, nil
}

type Order struct {
	Account      common.Address
	Nonce        *big.Int
	Path         []byte
	Amount       *big.Int
	TriggerPrice *big.Int
	Slippage     *big.Int
	OrderType    uint8
	OrderSide    uint8
	Deadline     *big.Int
	Signature    []byte
}

func ExecuteOrderData(contract *DexonTransactor, order *Order) ([]byte, error) {
	return CreateTxData(contract, "executeOrder", order)
}

type TwapOrder struct {
	Account        common.Address
	Nonce          *big.Int
	Path           []byte
	Amount         *big.Int
	OrderSide      uint8
	Signature      []byte
	Interval       *big.Int
	TotalOrders    *big.Int
	StartTimestamp *big.Int
}

func ExecuteTwapOrderData(contract *DexonTransactor, order *TwapOrder) ([]byte, error) {
	return CreateTxData(contract, "executeTwapOrder", order)
}

func ParseOrderExecutedEvent(filterer *DexonFilterer, receipt *types.Receipt) (*DexonOrderExecuted, error) {
	for _, log := range receipt.Logs {
		if log.Address == config.Env.DexonContractAddress {
			event, err := filterer.ParseOrderExecuted(*log)
			if err != nil {
				continue
			}
			return event, nil
		}
	}
	return nil, fmt.Errorf("failed to parse OrderExecuted event")
}

func ParseTwapOrderExecutedEvent(filterer *DexonFilterer, receipt *types.Receipt) (*DexonTwapOrderExecuted, error) {
	for _, log := range receipt.Logs {
		if log.Address == config.Env.DexonContractAddress {
			event, err := filterer.ParseTwapOrderExecuted(*log)
			if err != nil {
				continue
			}
			return event, nil
		}
	}
	return nil, fmt.Errorf("failed to parse TwapOrderExecuted event")
}
