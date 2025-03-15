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

// YexusOrder is an auto generated low-level Go binding around an user-defined struct.
type YexusOrder struct {
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

// YexusTwapOrder is an auto generated low-level Go binding around an user-defined struct.
type YexusTwapOrder struct {
	Account        common.Address
	Nonce          *big.Int
	Path           []byte
	Amount         *big.Int
	OrderSide      uint8
	Interval       *big.Int
	TotalOrders    *big.Int
	StartTimestamp *big.Int
	Signature      []byte
}

// YexusMetaData contains all meta data concerning the Yexus contract.
var YexusMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NAME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ONE_HUNDRED_PERCENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORDER_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TWAP_ORDER_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNISWAP_V3_FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNISWAP_V3_ROUTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETH_USDC_POOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structYexus.Order\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slippage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumYexus.OrderType\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"internalType\":\"enumYexus.OrderSide\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeTwapOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structYexus.TwapOrder\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"internalType\":\"enumYexus.OrderSide\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalOrders\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTokenPriceOnUsdc\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"twapCounts\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"twapCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderExecuted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"baseAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quoteAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slippage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumYexus.OrderType\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumYexus.OrderSide\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TwapOrderExecuted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"orderNth\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"baseAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quoteAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executedTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumYexus.OrderSide\"},{\"name\":\"totalBaseAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalOrders\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// YexusABI is the input ABI used to generate the binding from.
// Deprecated: Use YexusMetaData.ABI instead.
var YexusABI = YexusMetaData.ABI

// Yexus is an auto generated Go binding around an Ethereum contract.
type Yexus struct {
	YexusCaller     // Read-only binding to the contract
	YexusTransactor // Write-only binding to the contract
	YexusFilterer   // Log filterer for contract events
}

// YexusCaller is an auto generated read-only Go binding around an Ethereum contract.
type YexusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YexusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YexusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YexusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YexusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YexusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YexusSession struct {
	Contract     *Yexus            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YexusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YexusCallerSession struct {
	Contract *YexusCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YexusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YexusTransactorSession struct {
	Contract     *YexusTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YexusRaw is an auto generated low-level Go binding around an Ethereum contract.
type YexusRaw struct {
	Contract *Yexus // Generic contract binding to access the raw methods on
}

// YexusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YexusCallerRaw struct {
	Contract *YexusCaller // Generic read-only contract binding to access the raw methods on
}

// YexusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YexusTransactorRaw struct {
	Contract *YexusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYexus creates a new instance of Yexus, bound to a specific deployed contract.
func NewYexus(address common.Address, backend bind.ContractBackend) (*Yexus, error) {
	contract, err := bindYexus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yexus{YexusCaller: YexusCaller{contract: contract}, YexusTransactor: YexusTransactor{contract: contract}, YexusFilterer: YexusFilterer{contract: contract}}, nil
}

// NewYexusCaller creates a new read-only instance of Yexus, bound to a specific deployed contract.
func NewYexusCaller(address common.Address, caller bind.ContractCaller) (*YexusCaller, error) {
	contract, err := bindYexus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YexusCaller{contract: contract}, nil
}

// NewYexusTransactor creates a new write-only instance of Yexus, bound to a specific deployed contract.
func NewYexusTransactor(address common.Address, transactor bind.ContractTransactor) (*YexusTransactor, error) {
	contract, err := bindYexus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YexusTransactor{contract: contract}, nil
}

// NewYexusFilterer creates a new log filterer instance of Yexus, bound to a specific deployed contract.
func NewYexusFilterer(address common.Address, filterer bind.ContractFilterer) (*YexusFilterer, error) {
	contract, err := bindYexus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YexusFilterer{contract: contract}, nil
}

// bindYexus binds a generic wrapper to an already deployed contract.
func bindYexus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YexusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yexus *YexusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yexus.Contract.YexusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yexus *YexusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yexus.Contract.YexusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yexus *YexusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yexus.Contract.YexusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yexus *YexusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yexus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yexus *YexusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yexus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yexus *YexusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yexus.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Yexus *YexusCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Yexus *YexusSession) NAME() (string, error) {
	return _Yexus.Contract.NAME(&_Yexus.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Yexus *YexusCallerSession) NAME() (string, error) {
	return _Yexus.Contract.NAME(&_Yexus.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Yexus *YexusCaller) ONEHUNDREDPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "ONE_HUNDRED_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Yexus *YexusSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Yexus.Contract.ONEHUNDREDPERCENT(&_Yexus.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Yexus *YexusCallerSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Yexus.Contract.ONEHUNDREDPERCENT(&_Yexus.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Yexus.Contract.ORDERTYPEHASH(&_Yexus.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Yexus.Contract.ORDERTYPEHASH(&_Yexus.CallOpts)
}

// TWAPORDERTYPEHASH is a free data retrieval call binding the contract method 0x75225a34.
//
// Solidity: function TWAP_ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusCaller) TWAPORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "TWAP_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TWAPORDERTYPEHASH is a free data retrieval call binding the contract method 0x75225a34.
//
// Solidity: function TWAP_ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusSession) TWAPORDERTYPEHASH() ([32]byte, error) {
	return _Yexus.Contract.TWAPORDERTYPEHASH(&_Yexus.CallOpts)
}

// TWAPORDERTYPEHASH is a free data retrieval call binding the contract method 0x75225a34.
//
// Solidity: function TWAP_ORDER_TYPEHASH() view returns(bytes32)
func (_Yexus *YexusCallerSession) TWAPORDERTYPEHASH() ([32]byte, error) {
	return _Yexus.Contract.TWAPORDERTYPEHASH(&_Yexus.CallOpts)
}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Yexus *YexusCaller) UNISWAPV3FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "UNISWAP_V3_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Yexus *YexusSession) UNISWAPV3FACTORY() (common.Address, error) {
	return _Yexus.Contract.UNISWAPV3FACTORY(&_Yexus.CallOpts)
}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Yexus *YexusCallerSession) UNISWAPV3FACTORY() (common.Address, error) {
	return _Yexus.Contract.UNISWAPV3FACTORY(&_Yexus.CallOpts)
}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Yexus *YexusCaller) UNISWAPV3ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "UNISWAP_V3_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Yexus *YexusSession) UNISWAPV3ROUTER() (common.Address, error) {
	return _Yexus.Contract.UNISWAPV3ROUTER(&_Yexus.CallOpts)
}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Yexus *YexusCallerSession) UNISWAPV3ROUTER() (common.Address, error) {
	return _Yexus.Contract.UNISWAPV3ROUTER(&_Yexus.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Yexus *YexusCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Yexus *YexusSession) USDC() (common.Address, error) {
	return _Yexus.Contract.USDC(&_Yexus.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Yexus *YexusCallerSession) USDC() (common.Address, error) {
	return _Yexus.Contract.USDC(&_Yexus.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Yexus *YexusCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Yexus *YexusSession) VERSION() (string, error) {
	return _Yexus.Contract.VERSION(&_Yexus.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Yexus *YexusCallerSession) VERSION() (string, error) {
	return _Yexus.Contract.VERSION(&_Yexus.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Yexus *YexusCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Yexus *YexusSession) WETH() (common.Address, error) {
	return _Yexus.Contract.WETH(&_Yexus.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Yexus *YexusCallerSession) WETH() (common.Address, error) {
	return _Yexus.Contract.WETH(&_Yexus.CallOpts)
}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Yexus *YexusCaller) WETHUSDCPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "WETH_USDC_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Yexus *YexusSession) WETHUSDCPOOL() (common.Address, error) {
	return _Yexus.Contract.WETHUSDCPOOL(&_Yexus.CallOpts)
}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Yexus *YexusCallerSession) WETHUSDCPOOL() (common.Address, error) {
	return _Yexus.Contract.WETHUSDCPOOL(&_Yexus.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Yexus *YexusCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Yexus *YexusSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Yexus.Contract.Eip712Domain(&_Yexus.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Yexus *YexusCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Yexus.Contract.Eip712Domain(&_Yexus.CallOpts)
}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Yexus *YexusCaller) GetTokenPriceOnUsdc(opts *bind.CallOpts, path []byte) (*big.Int, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "getTokenPriceOnUsdc", path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Yexus *YexusSession) GetTokenPriceOnUsdc(path []byte) (*big.Int, error) {
	return _Yexus.Contract.GetTokenPriceOnUsdc(&_Yexus.CallOpts, path)
}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Yexus *YexusCallerSession) GetTokenPriceOnUsdc(path []byte) (*big.Int, error) {
	return _Yexus.Contract.GetTokenPriceOnUsdc(&_Yexus.CallOpts, path)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Yexus *YexusCaller) Nonces(opts *bind.CallOpts, account common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "nonces", account, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Yexus *YexusSession) Nonces(account common.Address, nonce *big.Int) (bool, error) {
	return _Yexus.Contract.Nonces(&_Yexus.CallOpts, account, nonce)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Yexus *YexusCallerSession) Nonces(account common.Address, nonce *big.Int) (bool, error) {
	return _Yexus.Contract.Nonces(&_Yexus.CallOpts, account, nonce)
}

// TwapCounts is a free data retrieval call binding the contract method 0x3fa9deda.
//
// Solidity: function twapCounts(address account, uint256 nonce) view returns(uint256 twapCount)
func (_Yexus *YexusCaller) TwapCounts(opts *bind.CallOpts, account common.Address, nonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yexus.contract.Call(opts, &out, "twapCounts", account, nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TwapCounts is a free data retrieval call binding the contract method 0x3fa9deda.
//
// Solidity: function twapCounts(address account, uint256 nonce) view returns(uint256 twapCount)
func (_Yexus *YexusSession) TwapCounts(account common.Address, nonce *big.Int) (*big.Int, error) {
	return _Yexus.Contract.TwapCounts(&_Yexus.CallOpts, account, nonce)
}

// TwapCounts is a free data retrieval call binding the contract method 0x3fa9deda.
//
// Solidity: function twapCounts(address account, uint256 nonce) view returns(uint256 twapCount)
func (_Yexus *YexusCallerSession) TwapCounts(account common.Address, nonce *big.Int) (*big.Int, error) {
	return _Yexus.Contract.TwapCounts(&_Yexus.CallOpts, account, nonce)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Yexus *YexusTransactor) ExecuteOrder(opts *bind.TransactOpts, order YexusOrder) (*types.Transaction, error) {
	return _Yexus.contract.Transact(opts, "executeOrder", order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Yexus *YexusSession) ExecuteOrder(order YexusOrder) (*types.Transaction, error) {
	return _Yexus.Contract.ExecuteOrder(&_Yexus.TransactOpts, order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Yexus *YexusTransactorSession) ExecuteOrder(order YexusOrder) (*types.Transaction, error) {
	return _Yexus.Contract.ExecuteOrder(&_Yexus.TransactOpts, order)
}

// ExecuteTwapOrder is a paid mutator transaction binding the contract method 0xdc3ff069.
//
// Solidity: function executeTwapOrder((address,uint256,bytes,uint256,uint8,uint256,uint256,uint256,bytes) order) returns()
func (_Yexus *YexusTransactor) ExecuteTwapOrder(opts *bind.TransactOpts, order YexusTwapOrder) (*types.Transaction, error) {
	return _Yexus.contract.Transact(opts, "executeTwapOrder", order)
}

// ExecuteTwapOrder is a paid mutator transaction binding the contract method 0xdc3ff069.
//
// Solidity: function executeTwapOrder((address,uint256,bytes,uint256,uint8,uint256,uint256,uint256,bytes) order) returns()
func (_Yexus *YexusSession) ExecuteTwapOrder(order YexusTwapOrder) (*types.Transaction, error) {
	return _Yexus.Contract.ExecuteTwapOrder(&_Yexus.TransactOpts, order)
}

// ExecuteTwapOrder is a paid mutator transaction binding the contract method 0xdc3ff069.
//
// Solidity: function executeTwapOrder((address,uint256,bytes,uint256,uint8,uint256,uint256,uint256,bytes) order) returns()
func (_Yexus *YexusTransactorSession) ExecuteTwapOrder(order YexusTwapOrder) (*types.Transaction, error) {
	return _Yexus.Contract.ExecuteTwapOrder(&_Yexus.TransactOpts, order)
}

// YexusEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Yexus contract.
type YexusEIP712DomainChangedIterator struct {
	Event *YexusEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *YexusEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YexusEIP712DomainChanged)
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
		it.Event = new(YexusEIP712DomainChanged)
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
func (it *YexusEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YexusEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YexusEIP712DomainChanged represents a EIP712DomainChanged event raised by the Yexus contract.
type YexusEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Yexus *YexusFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*YexusEIP712DomainChangedIterator, error) {

	logs, sub, err := _Yexus.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &YexusEIP712DomainChangedIterator{contract: _Yexus.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Yexus *YexusFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *YexusEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Yexus.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YexusEIP712DomainChanged)
				if err := _Yexus.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Yexus *YexusFilterer) ParseEIP712DomainChanged(log types.Log) (*YexusEIP712DomainChanged, error) {
	event := new(YexusEIP712DomainChanged)
	if err := _Yexus.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YexusOrderExecutedIterator is returned from FilterOrderExecuted and is used to iterate over the raw logs and unpacked data for OrderExecuted events raised by the Yexus contract.
type YexusOrderExecutedIterator struct {
	Event *YexusOrderExecuted // Event containing the contract specifics and raw log

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
func (it *YexusOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YexusOrderExecuted)
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
		it.Event = new(YexusOrderExecuted)
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
func (it *YexusOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YexusOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YexusOrderExecuted represents a OrderExecuted event raised by the Yexus contract.
type YexusOrderExecuted struct {
	Account      common.Address
	Nonce        *big.Int
	Path         []byte
	BaseAmount   *big.Int
	QuoteAmount  *big.Int
	TriggerPrice *big.Int
	Slippage     *big.Int
	OrderType    uint8
	OrderSide    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderExecuted is a free log retrieval operation binding the contract event 0xc93dd372329320fe5794f13c7039ac2ac5d719c59333fcce8ce1088bc6eae671.
//
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 baseAmount, uint256 quoteAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Yexus *YexusFilterer) FilterOrderExecuted(opts *bind.FilterOpts, account []common.Address, nonce []*big.Int) (*YexusOrderExecutedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Yexus.contract.FilterLogs(opts, "OrderExecuted", accountRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &YexusOrderExecutedIterator{contract: _Yexus.contract, event: "OrderExecuted", logs: logs, sub: sub}, nil
}

// WatchOrderExecuted is a free log subscription operation binding the contract event 0xc93dd372329320fe5794f13c7039ac2ac5d719c59333fcce8ce1088bc6eae671.
//
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 baseAmount, uint256 quoteAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Yexus *YexusFilterer) WatchOrderExecuted(opts *bind.WatchOpts, sink chan<- *YexusOrderExecuted, account []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Yexus.contract.WatchLogs(opts, "OrderExecuted", accountRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YexusOrderExecuted)
				if err := _Yexus.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
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

// ParseOrderExecuted is a log parse operation binding the contract event 0xc93dd372329320fe5794f13c7039ac2ac5d719c59333fcce8ce1088bc6eae671.
//
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 baseAmount, uint256 quoteAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Yexus *YexusFilterer) ParseOrderExecuted(log types.Log) (*YexusOrderExecuted, error) {
	event := new(YexusOrderExecuted)
	if err := _Yexus.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YexusTwapOrderExecutedIterator is returned from FilterTwapOrderExecuted and is used to iterate over the raw logs and unpacked data for TwapOrderExecuted events raised by the Yexus contract.
type YexusTwapOrderExecutedIterator struct {
	Event *YexusTwapOrderExecuted // Event containing the contract specifics and raw log

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
func (it *YexusTwapOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YexusTwapOrderExecuted)
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
		it.Event = new(YexusTwapOrderExecuted)
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
func (it *YexusTwapOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YexusTwapOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YexusTwapOrderExecuted represents a TwapOrderExecuted event raised by the Yexus contract.
type YexusTwapOrderExecuted struct {
	Account           common.Address
	Nonce             *big.Int
	OrderNth          *big.Int
	BaseAmount        *big.Int
	QuoteAmount       *big.Int
	ExecutedTimestamp *big.Int
	Path              []byte
	OrderSide         uint8
	TotalBaseAmount   *big.Int
	Interval          *big.Int
	TotalOrders       *big.Int
	StartTimestamp    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTwapOrderExecuted is a free log retrieval operation binding the contract event 0xbd1f49d33312b791c378c4e094f99abf2b236db830ec495e907218ee04604fd9.
//
// Solidity: event TwapOrderExecuted(address indexed account, uint256 indexed nonce, uint256 indexed orderNth, uint256 baseAmount, uint256 quoteAmount, uint256 executedTimestamp, bytes path, uint8 orderSide, uint256 totalBaseAmount, uint256 interval, uint256 totalOrders, uint256 startTimestamp)
func (_Yexus *YexusFilterer) FilterTwapOrderExecuted(opts *bind.FilterOpts, account []common.Address, nonce []*big.Int, orderNth []*big.Int) (*YexusTwapOrderExecutedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var orderNthRule []interface{}
	for _, orderNthItem := range orderNth {
		orderNthRule = append(orderNthRule, orderNthItem)
	}

	logs, sub, err := _Yexus.contract.FilterLogs(opts, "TwapOrderExecuted", accountRule, nonceRule, orderNthRule)
	if err != nil {
		return nil, err
	}
	return &YexusTwapOrderExecutedIterator{contract: _Yexus.contract, event: "TwapOrderExecuted", logs: logs, sub: sub}, nil
}

// WatchTwapOrderExecuted is a free log subscription operation binding the contract event 0xbd1f49d33312b791c378c4e094f99abf2b236db830ec495e907218ee04604fd9.
//
// Solidity: event TwapOrderExecuted(address indexed account, uint256 indexed nonce, uint256 indexed orderNth, uint256 baseAmount, uint256 quoteAmount, uint256 executedTimestamp, bytes path, uint8 orderSide, uint256 totalBaseAmount, uint256 interval, uint256 totalOrders, uint256 startTimestamp)
func (_Yexus *YexusFilterer) WatchTwapOrderExecuted(opts *bind.WatchOpts, sink chan<- *YexusTwapOrderExecuted, account []common.Address, nonce []*big.Int, orderNth []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var orderNthRule []interface{}
	for _, orderNthItem := range orderNth {
		orderNthRule = append(orderNthRule, orderNthItem)
	}

	logs, sub, err := _Yexus.contract.WatchLogs(opts, "TwapOrderExecuted", accountRule, nonceRule, orderNthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YexusTwapOrderExecuted)
				if err := _Yexus.contract.UnpackLog(event, "TwapOrderExecuted", log); err != nil {
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

// ParseTwapOrderExecuted is a log parse operation binding the contract event 0xbd1f49d33312b791c378c4e094f99abf2b236db830ec495e907218ee04604fd9.
//
// Solidity: event TwapOrderExecuted(address indexed account, uint256 indexed nonce, uint256 indexed orderNth, uint256 baseAmount, uint256 quoteAmount, uint256 executedTimestamp, bytes path, uint8 orderSide, uint256 totalBaseAmount, uint256 interval, uint256 totalOrders, uint256 startTimestamp)
func (_Yexus *YexusFilterer) ParseTwapOrderExecuted(log types.Log) (*YexusTwapOrderExecuted, error) {
	event := new(YexusTwapOrderExecuted)
	if err := _Yexus.contract.UnpackLog(event, "TwapOrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
