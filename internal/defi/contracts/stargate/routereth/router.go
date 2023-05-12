// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routereth

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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stargateEthVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stargateRouter\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_poolId\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addLiquidityETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargateEthVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargateRouter\",\"outputs\":[{\"internalType\":\"contractIStargateRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"}],\"name\":\"swapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Router *RouterCaller) PoolId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Router *RouterSession) PoolId() (uint16, error) {
	return _Router.Contract.PoolId(&_Router.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Router *RouterCallerSession) PoolId() (uint16, error) {
	return _Router.Contract.PoolId(&_Router.CallOpts)
}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Router *RouterCaller) StargateEthVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "stargateEthVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Router *RouterSession) StargateEthVault() (common.Address, error) {
	return _Router.Contract.StargateEthVault(&_Router.CallOpts)
}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Router *RouterCallerSession) StargateEthVault() (common.Address, error) {
	return _Router.Contract.StargateEthVault(&_Router.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Router *RouterCaller) StargateRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "stargateRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Router *RouterSession) StargateRouter() (common.Address, error) {
	return _Router.Contract.StargateRouter(&_Router.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Router *RouterCallerSession) StargateRouter() (common.Address, error) {
	return _Router.Contract.StargateRouter(&_Router.CallOpts)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Router *RouterTransactor) AddLiquidityETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityETH")
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Router *RouterSession) AddLiquidityETH() (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Router *RouterTransactorSession) AddLiquidityETH() (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityETH(&_Router.TransactOpts)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Router *RouterTransactor) SwapETH(opts *bind.TransactOpts, _dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapETH", _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Router *RouterSession) SwapETH(_dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapETH(&_Router.TransactOpts, _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Router *RouterTransactorSession) SwapETH(_dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapETH(&_Router.TransactOpts, _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}
