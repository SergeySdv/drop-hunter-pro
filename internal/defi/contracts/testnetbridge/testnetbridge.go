// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testnetbridge

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

// TestnetbridgeMetaData contains all meta data concerning the Testnetbridge contract.
var TestnetbridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"oft\",\"outputs\":[{\"internalType\":\"contractIOFTCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"name\":\"swapAndBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TestnetbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use TestnetbridgeMetaData.ABI instead.
var TestnetbridgeABI = TestnetbridgeMetaData.ABI

// Testnetbridge is an auto generated Go binding around an Ethereum contract.
type Testnetbridge struct {
	TestnetbridgeCaller     // Read-only binding to the contract
	TestnetbridgeTransactor // Write-only binding to the contract
	TestnetbridgeFilterer   // Log filterer for contract events
}

// TestnetbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestnetbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestnetbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestnetbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestnetbridgeSession struct {
	Contract     *Testnetbridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestnetbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestnetbridgeCallerSession struct {
	Contract *TestnetbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestnetbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestnetbridgeTransactorSession struct {
	Contract     *TestnetbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestnetbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestnetbridgeRaw struct {
	Contract *Testnetbridge // Generic contract binding to access the raw methods on
}

// TestnetbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestnetbridgeCallerRaw struct {
	Contract *TestnetbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// TestnetbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestnetbridgeTransactorRaw struct {
	Contract *TestnetbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestnetbridge creates a new instance of Testnetbridge, bound to a specific deployed contract.
func NewTestnetbridge(address common.Address, backend bind.ContractBackend) (*Testnetbridge, error) {
	contract, err := bindTestnetbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testnetbridge{TestnetbridgeCaller: TestnetbridgeCaller{contract: contract}, TestnetbridgeTransactor: TestnetbridgeTransactor{contract: contract}, TestnetbridgeFilterer: TestnetbridgeFilterer{contract: contract}}, nil
}

// NewTestnetbridgeCaller creates a new read-only instance of Testnetbridge, bound to a specific deployed contract.
func NewTestnetbridgeCaller(address common.Address, caller bind.ContractCaller) (*TestnetbridgeCaller, error) {
	contract, err := bindTestnetbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestnetbridgeCaller{contract: contract}, nil
}

// NewTestnetbridgeTransactor creates a new write-only instance of Testnetbridge, bound to a specific deployed contract.
func NewTestnetbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*TestnetbridgeTransactor, error) {
	contract, err := bindTestnetbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestnetbridgeTransactor{contract: contract}, nil
}

// NewTestnetbridgeFilterer creates a new log filterer instance of Testnetbridge, bound to a specific deployed contract.
func NewTestnetbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*TestnetbridgeFilterer, error) {
	contract, err := bindTestnetbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestnetbridgeFilterer{contract: contract}, nil
}

// bindTestnetbridge binds a generic wrapper to an already deployed contract.
func bindTestnetbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestnetbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testnetbridge *TestnetbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testnetbridge.Contract.TestnetbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testnetbridge *TestnetbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnetbridge.Contract.TestnetbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testnetbridge *TestnetbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testnetbridge.Contract.TestnetbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testnetbridge *TestnetbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testnetbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testnetbridge *TestnetbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnetbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testnetbridge *TestnetbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testnetbridge.Contract.contract.Transact(opts, method, params...)
}

// Oft is a free data retrieval call binding the contract method 0x9b5215f6.
//
// Solidity: function oft() view returns(address)
func (_Testnetbridge *TestnetbridgeCaller) Oft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Testnetbridge.contract.Call(opts, &out, "oft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oft is a free data retrieval call binding the contract method 0x9b5215f6.
//
// Solidity: function oft() view returns(address)
func (_Testnetbridge *TestnetbridgeSession) Oft() (common.Address, error) {
	return _Testnetbridge.Contract.Oft(&_Testnetbridge.CallOpts)
}

// Oft is a free data retrieval call binding the contract method 0x9b5215f6.
//
// Solidity: function oft() view returns(address)
func (_Testnetbridge *TestnetbridgeCallerSession) Oft() (common.Address, error) {
	return _Testnetbridge.Contract.Oft(&_Testnetbridge.CallOpts)
}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_Testnetbridge *TestnetbridgeCaller) PoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Testnetbridge.contract.Call(opts, &out, "poolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_Testnetbridge *TestnetbridgeSession) PoolFee() (*big.Int, error) {
	return _Testnetbridge.Contract.PoolFee(&_Testnetbridge.CallOpts)
}

// PoolFee is a free data retrieval call binding the contract method 0x089fe6aa.
//
// Solidity: function poolFee() view returns(uint24)
func (_Testnetbridge *TestnetbridgeCallerSession) PoolFee() (*big.Int, error) {
	return _Testnetbridge.Contract.PoolFee(&_Testnetbridge.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Testnetbridge *TestnetbridgeCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Testnetbridge.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Testnetbridge *TestnetbridgeSession) SwapRouter() (common.Address, error) {
	return _Testnetbridge.Contract.SwapRouter(&_Testnetbridge.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Testnetbridge *TestnetbridgeCallerSession) SwapRouter() (common.Address, error) {
	return _Testnetbridge.Contract.SwapRouter(&_Testnetbridge.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Testnetbridge *TestnetbridgeCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Testnetbridge.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Testnetbridge *TestnetbridgeSession) Weth() (common.Address, error) {
	return _Testnetbridge.Contract.Weth(&_Testnetbridge.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Testnetbridge *TestnetbridgeCallerSession) Weth() (common.Address, error) {
	return _Testnetbridge.Contract.Weth(&_Testnetbridge.CallOpts)
}

// SwapAndBridge is a paid mutator transaction binding the contract method 0xae30f6ee.
//
// Solidity: function swapAndBridge(uint256 amountIn, uint256 amountOutMin, uint16 dstChainId, address to, address refundAddress, address zroPaymentAddress, bytes adapterParams) payable returns()
func (_Testnetbridge *TestnetbridgeTransactor) SwapAndBridge(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, dstChainId uint16, to common.Address, refundAddress common.Address, zroPaymentAddress common.Address, adapterParams []byte) (*types.Transaction, error) {
	return _Testnetbridge.contract.Transact(opts, "swapAndBridge", amountIn, amountOutMin, dstChainId, to, refundAddress, zroPaymentAddress, adapterParams)
}

// SwapAndBridge is a paid mutator transaction binding the contract method 0xae30f6ee.
//
// Solidity: function swapAndBridge(uint256 amountIn, uint256 amountOutMin, uint16 dstChainId, address to, address refundAddress, address zroPaymentAddress, bytes adapterParams) payable returns()
func (_Testnetbridge *TestnetbridgeSession) SwapAndBridge(amountIn *big.Int, amountOutMin *big.Int, dstChainId uint16, to common.Address, refundAddress common.Address, zroPaymentAddress common.Address, adapterParams []byte) (*types.Transaction, error) {
	return _Testnetbridge.Contract.SwapAndBridge(&_Testnetbridge.TransactOpts, amountIn, amountOutMin, dstChainId, to, refundAddress, zroPaymentAddress, adapterParams)
}

// SwapAndBridge is a paid mutator transaction binding the contract method 0xae30f6ee.
//
// Solidity: function swapAndBridge(uint256 amountIn, uint256 amountOutMin, uint16 dstChainId, address to, address refundAddress, address zroPaymentAddress, bytes adapterParams) payable returns()
func (_Testnetbridge *TestnetbridgeTransactorSession) SwapAndBridge(amountIn *big.Int, amountOutMin *big.Int, dstChainId uint16, to common.Address, refundAddress common.Address, zroPaymentAddress common.Address, adapterParams []byte) (*types.Transaction, error) {
	return _Testnetbridge.Contract.SwapAndBridge(&_Testnetbridge.TransactOpts, amountIn, amountOutMin, dstChainId, to, refundAddress, zroPaymentAddress, adapterParams)
}
