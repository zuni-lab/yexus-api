// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// UniswapV3MetaData contains all meta data concerning the UniswapV3 contract.
var UniswapV3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// UniswapV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3MetaData.ABI instead.
var UniswapV3ABI = UniswapV3MetaData.ABI

// UniswapV3 is an auto generated Go binding around an Ethereum contract.
type UniswapV3 struct {
	UniswapV3Caller     // Read-only binding to the contract
	UniswapV3Transactor // Write-only binding to the contract
	UniswapV3Filterer   // Log filterer for contract events
}

// UniswapV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3Session struct {
	Contract     *UniswapV3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3CallerSession struct {
	Contract *UniswapV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UniswapV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3TransactorSession struct {
	Contract     *UniswapV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UniswapV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3Raw struct {
	Contract *UniswapV3 // Generic contract binding to access the raw methods on
}

// UniswapV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3CallerRaw struct {
	Contract *UniswapV3Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3TransactorRaw struct {
	Contract *UniswapV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3 creates a new instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3(address common.Address, backend bind.ContractBackend) (*UniswapV3, error) {
	contract, err := bindUniswapV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3{UniswapV3Caller: UniswapV3Caller{contract: contract}, UniswapV3Transactor: UniswapV3Transactor{contract: contract}, UniswapV3Filterer: UniswapV3Filterer{contract: contract}}, nil
}

// NewUniswapV3Caller creates a new read-only instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Caller(address common.Address, caller bind.ContractCaller) (*UniswapV3Caller, error) {
	contract, err := bindUniswapV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Caller{contract: contract}, nil
}

// NewUniswapV3Transactor creates a new write-only instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3Transactor, error) {
	contract, err := bindUniswapV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Transactor{contract: contract}, nil
}

// NewUniswapV3Filterer creates a new log filterer instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3Filterer, error) {
	contract, err := bindUniswapV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Filterer{contract: contract}, nil
}

// bindUniswapV3 binds a generic wrapper to an already deployed contract.
func bindUniswapV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3 *UniswapV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3.Contract.UniswapV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3 *UniswapV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3.Contract.UniswapV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3 *UniswapV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3.Contract.UniswapV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3 *UniswapV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3 *UniswapV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3 *UniswapV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3.Contract.contract.Transact(opts, method, params...)
}

// UniswapV3SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the UniswapV3 contract.
type UniswapV3SwapIterator struct {
	Event *UniswapV3Swap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapV3SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV3Swap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapV3Swap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapV3SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV3SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV3Swap represents a Swap event raised by the UniswapV3 contract.
type UniswapV3Swap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_UniswapV3 *UniswapV3Filterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*UniswapV3SwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &UniswapV3SwapIterator{contract: _UniswapV3.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_UniswapV3 *UniswapV3Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *UniswapV3Swap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _UniswapV3.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV3Swap)
				if err := _UniswapV3.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_UniswapV3 *UniswapV3Filterer) ParseSwap(log types.Log) (*UniswapV3Swap, error) {
	event := new(UniswapV3Swap)
	if err := _UniswapV3.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
