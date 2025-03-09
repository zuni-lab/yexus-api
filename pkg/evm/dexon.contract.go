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

// DexonOrder is an auto generated low-level Go binding around an user-defined struct.
type DexonOrder struct {
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

// DexonMetaData contains all meta data concerning the Dexon contract.
var DexonMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NAME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ONE_HUNDRED_PERCENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORDER_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNISWAP_V3_FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNISWAP_V3_ROUTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETH_USDC_POOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structDexon.Order\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slippage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumDexon.OrderType\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"internalType\":\"enumDexon.OrderSide\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTokenPriceOnUsdc\",\"inputs\":[{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderExecuted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actualSwapAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slippage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"orderType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumDexon.OrderType\"},{\"name\":\"orderSide\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumDexon.OrderSide\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// DexonABI is the input ABI used to generate the binding from.
// Deprecated: Use DexonMetaData.ABI instead.
var DexonABI = DexonMetaData.ABI

// Dexon is an auto generated Go binding around an Ethereum contract.
type Dexon struct {
	DexonCaller     // Read-only binding to the contract
	DexonTransactor // Write-only binding to the contract
	DexonFilterer   // Log filterer for contract events
}

// DexonCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexonSession struct {
	Contract     *Dexon            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexonCallerSession struct {
	Contract *DexonCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexonTransactorSession struct {
	Contract     *DexonTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexonRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexonRaw struct {
	Contract *Dexon // Generic contract binding to access the raw methods on
}

// DexonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexonCallerRaw struct {
	Contract *DexonCaller // Generic read-only contract binding to access the raw methods on
}

// DexonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexonTransactorRaw struct {
	Contract *DexonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexon creates a new instance of Dexon, bound to a specific deployed contract.
func NewDexon(address common.Address, backend bind.ContractBackend) (*Dexon, error) {
	contract, err := bindDexon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dexon{DexonCaller: DexonCaller{contract: contract}, DexonTransactor: DexonTransactor{contract: contract}, DexonFilterer: DexonFilterer{contract: contract}}, nil
}

// NewDexonCaller creates a new read-only instance of Dexon, bound to a specific deployed contract.
func NewDexonCaller(address common.Address, caller bind.ContractCaller) (*DexonCaller, error) {
	contract, err := bindDexon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexonCaller{contract: contract}, nil
}

// NewDexonTransactor creates a new write-only instance of Dexon, bound to a specific deployed contract.
func NewDexonTransactor(address common.Address, transactor bind.ContractTransactor) (*DexonTransactor, error) {
	contract, err := bindDexon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexonTransactor{contract: contract}, nil
}

// NewDexonFilterer creates a new log filterer instance of Dexon, bound to a specific deployed contract.
func NewDexonFilterer(address common.Address, filterer bind.ContractFilterer) (*DexonFilterer, error) {
	contract, err := bindDexon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexonFilterer{contract: contract}, nil
}

// bindDexon binds a generic wrapper to an already deployed contract.
func bindDexon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DexonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexon *DexonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dexon.Contract.DexonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexon *DexonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexon.Contract.DexonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexon *DexonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexon.Contract.DexonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dexon *DexonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dexon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dexon *DexonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dexon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dexon *DexonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dexon.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Dexon *DexonCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Dexon *DexonSession) NAME() (string, error) {
	return _Dexon.Contract.NAME(&_Dexon.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Dexon *DexonCallerSession) NAME() (string, error) {
	return _Dexon.Contract.NAME(&_Dexon.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Dexon *DexonCaller) ONEHUNDREDPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "ONE_HUNDRED_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Dexon *DexonSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Dexon.Contract.ONEHUNDREDPERCENT(&_Dexon.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Dexon *DexonCallerSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Dexon.Contract.ONEHUNDREDPERCENT(&_Dexon.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Dexon *DexonCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Dexon *DexonSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Dexon.Contract.ORDERTYPEHASH(&_Dexon.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Dexon *DexonCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Dexon.Contract.ORDERTYPEHASH(&_Dexon.CallOpts)
}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Dexon *DexonCaller) UNISWAPV3FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "UNISWAP_V3_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Dexon *DexonSession) UNISWAPV3FACTORY() (common.Address, error) {
	return _Dexon.Contract.UNISWAPV3FACTORY(&_Dexon.CallOpts)
}

// UNISWAPV3FACTORY is a free data retrieval call binding the contract method 0xf73e5aab.
//
// Solidity: function UNISWAP_V3_FACTORY() view returns(address)
func (_Dexon *DexonCallerSession) UNISWAPV3FACTORY() (common.Address, error) {
	return _Dexon.Contract.UNISWAPV3FACTORY(&_Dexon.CallOpts)
}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Dexon *DexonCaller) UNISWAPV3ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "UNISWAP_V3_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Dexon *DexonSession) UNISWAPV3ROUTER() (common.Address, error) {
	return _Dexon.Contract.UNISWAPV3ROUTER(&_Dexon.CallOpts)
}

// UNISWAPV3ROUTER is a free data retrieval call binding the contract method 0x41c64a2f.
//
// Solidity: function UNISWAP_V3_ROUTER() view returns(address)
func (_Dexon *DexonCallerSession) UNISWAPV3ROUTER() (common.Address, error) {
	return _Dexon.Contract.UNISWAPV3ROUTER(&_Dexon.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Dexon *DexonCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Dexon *DexonSession) USDC() (common.Address, error) {
	return _Dexon.Contract.USDC(&_Dexon.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Dexon *DexonCallerSession) USDC() (common.Address, error) {
	return _Dexon.Contract.USDC(&_Dexon.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Dexon *DexonCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Dexon *DexonSession) VERSION() (string, error) {
	return _Dexon.Contract.VERSION(&_Dexon.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Dexon *DexonCallerSession) VERSION() (string, error) {
	return _Dexon.Contract.VERSION(&_Dexon.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Dexon *DexonCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Dexon *DexonSession) WETH() (common.Address, error) {
	return _Dexon.Contract.WETH(&_Dexon.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Dexon *DexonCallerSession) WETH() (common.Address, error) {
	return _Dexon.Contract.WETH(&_Dexon.CallOpts)
}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Dexon *DexonCaller) WETHUSDCPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "WETH_USDC_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Dexon *DexonSession) WETHUSDCPOOL() (common.Address, error) {
	return _Dexon.Contract.WETHUSDCPOOL(&_Dexon.CallOpts)
}

// WETHUSDCPOOL is a free data retrieval call binding the contract method 0xe492288f.
//
// Solidity: function WETH_USDC_POOL() view returns(address)
func (_Dexon *DexonCallerSession) WETHUSDCPOOL() (common.Address, error) {
	return _Dexon.Contract.WETHUSDCPOOL(&_Dexon.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Dexon *DexonCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "eip712Domain")

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
func (_Dexon *DexonSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Dexon.Contract.Eip712Domain(&_Dexon.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Dexon *DexonCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Dexon.Contract.Eip712Domain(&_Dexon.CallOpts)
}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Dexon *DexonCaller) GetTokenPriceOnUsdc(opts *bind.CallOpts, path []byte) (*big.Int, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "getTokenPriceOnUsdc", path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Dexon *DexonSession) GetTokenPriceOnUsdc(path []byte) (*big.Int, error) {
	return _Dexon.Contract.GetTokenPriceOnUsdc(&_Dexon.CallOpts, path)
}

// GetTokenPriceOnUsdc is a free data retrieval call binding the contract method 0x74ddafad.
//
// Solidity: function getTokenPriceOnUsdc(bytes path) view returns(uint256)
func (_Dexon *DexonCallerSession) GetTokenPriceOnUsdc(path []byte) (*big.Int, error) {
	return _Dexon.Contract.GetTokenPriceOnUsdc(&_Dexon.CallOpts, path)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Dexon *DexonCaller) Nonces(opts *bind.CallOpts, account common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Dexon.contract.Call(opts, &out, "nonces", account, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Dexon *DexonSession) Nonces(account common.Address, nonce *big.Int) (bool, error) {
	return _Dexon.Contract.Nonces(&_Dexon.CallOpts, account, nonce)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address account, uint256 nonce) view returns(bool used)
func (_Dexon *DexonCallerSession) Nonces(account common.Address, nonce *big.Int) (bool, error) {
	return _Dexon.Contract.Nonces(&_Dexon.CallOpts, account, nonce)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Dexon *DexonTransactor) ExecuteOrder(opts *bind.TransactOpts, order DexonOrder) (*types.Transaction, error) {
	return _Dexon.contract.Transact(opts, "executeOrder", order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Dexon *DexonSession) ExecuteOrder(order DexonOrder) (*types.Transaction, error) {
	return _Dexon.Contract.ExecuteOrder(&_Dexon.TransactOpts, order)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6bc180f8.
//
// Solidity: function executeOrder((address,uint256,bytes,uint256,uint256,uint256,uint8,uint8,uint256,bytes) order) returns()
func (_Dexon *DexonTransactorSession) ExecuteOrder(order DexonOrder) (*types.Transaction, error) {
	return _Dexon.Contract.ExecuteOrder(&_Dexon.TransactOpts, order)
}

// DexonEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Dexon contract.
type DexonEIP712DomainChangedIterator struct {
	Event *DexonEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DexonEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexonEIP712DomainChanged)
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
		it.Event = new(DexonEIP712DomainChanged)
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
func (it *DexonEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexonEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexonEIP712DomainChanged represents a EIP712DomainChanged event raised by the Dexon contract.
type DexonEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Dexon *DexonFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DexonEIP712DomainChangedIterator, error) {

	logs, sub, err := _Dexon.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DexonEIP712DomainChangedIterator{contract: _Dexon.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Dexon *DexonFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DexonEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Dexon.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexonEIP712DomainChanged)
				if err := _Dexon.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Dexon *DexonFilterer) ParseEIP712DomainChanged(log types.Log) (*DexonEIP712DomainChanged, error) {
	event := new(DexonEIP712DomainChanged)
	if err := _Dexon.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexonOrderExecutedIterator is returned from FilterOrderExecuted and is used to iterate over the raw logs and unpacked data for OrderExecuted events raised by the Dexon contract.
type DexonOrderExecutedIterator struct {
	Event *DexonOrderExecuted // Event containing the contract specifics and raw log

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
func (it *DexonOrderExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexonOrderExecuted)
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
		it.Event = new(DexonOrderExecuted)
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
func (it *DexonOrderExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexonOrderExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexonOrderExecuted represents a OrderExecuted event raised by the Dexon contract.
type DexonOrderExecuted struct {
	Account          common.Address
	Nonce            *big.Int
	Path             []byte
	Amount           *big.Int
	ActualSwapAmount *big.Int
	TriggerPrice     *big.Int
	Slippage         *big.Int
	OrderType        uint8
	OrderSide        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderExecuted is a free log retrieval operation binding the contract event 0xc93dd372329320fe5794f13c7039ac2ac5d719c59333fcce8ce1088bc6eae671.
//
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 amount, uint256 actualSwapAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Dexon *DexonFilterer) FilterOrderExecuted(opts *bind.FilterOpts, account []common.Address, nonce []*big.Int) (*DexonOrderExecutedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Dexon.contract.FilterLogs(opts, "OrderExecuted", accountRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &DexonOrderExecutedIterator{contract: _Dexon.contract, event: "OrderExecuted", logs: logs, sub: sub}, nil
}

// WatchOrderExecuted is a free log subscription operation binding the contract event 0xc93dd372329320fe5794f13c7039ac2ac5d719c59333fcce8ce1088bc6eae671.
//
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 amount, uint256 actualSwapAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Dexon *DexonFilterer) WatchOrderExecuted(opts *bind.WatchOpts, sink chan<- *DexonOrderExecuted, account []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Dexon.contract.WatchLogs(opts, "OrderExecuted", accountRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexonOrderExecuted)
				if err := _Dexon.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
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
// Solidity: event OrderExecuted(address indexed account, uint256 indexed nonce, bytes path, uint256 amount, uint256 actualSwapAmount, uint256 triggerPrice, uint256 slippage, uint8 orderType, uint8 orderSide)
func (_Dexon *DexonFilterer) ParseOrderExecuted(log types.Log) (*DexonOrderExecuted, error) {
	event := new(DexonOrderExecuted)
	if err := _Dexon.contract.UnpackLog(event, "OrderExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
