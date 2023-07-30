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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stargateEthVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stargateRouter\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_poolId\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addLiquidityETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargateEthVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargateRouter\",\"outputs\":[{\"internalType\":\"contractIStargateRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"}],\"name\":\"swapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Storage *StorageCaller) PoolId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Storage *StorageSession) PoolId() (uint16, error) {
	return _Storage.Contract.PoolId(&_Storage.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint16)
func (_Storage *StorageCallerSession) PoolId() (uint16, error) {
	return _Storage.Contract.PoolId(&_Storage.CallOpts)
}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Storage *StorageCaller) StargateEthVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "stargateEthVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Storage *StorageSession) StargateEthVault() (common.Address, error) {
	return _Storage.Contract.StargateEthVault(&_Storage.CallOpts)
}

// StargateEthVault is a free data retrieval call binding the contract method 0x38e31d39.
//
// Solidity: function stargateEthVault() view returns(address)
func (_Storage *StorageCallerSession) StargateEthVault() (common.Address, error) {
	return _Storage.Contract.StargateEthVault(&_Storage.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Storage *StorageCaller) StargateRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "stargateRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Storage *StorageSession) StargateRouter() (common.Address, error) {
	return _Storage.Contract.StargateRouter(&_Storage.CallOpts)
}

// StargateRouter is a free data retrieval call binding the contract method 0xa9e56f3c.
//
// Solidity: function stargateRouter() view returns(address)
func (_Storage *StorageCallerSession) StargateRouter() (common.Address, error) {
	return _Storage.Contract.StargateRouter(&_Storage.CallOpts)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Storage *StorageTransactor) AddLiquidityETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidityETH")
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Storage *StorageSession) AddLiquidityETH() (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityETH(&_Storage.TransactOpts)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xed995307.
//
// Solidity: function addLiquidityETH() payable returns()
func (_Storage *StorageTransactorSession) AddLiquidityETH() (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityETH(&_Storage.TransactOpts)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Storage *StorageTransactor) SwapETH(opts *bind.TransactOpts, _dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapETH", _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Storage *StorageSession) SwapETH(_dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapETH(&_Storage.TransactOpts, _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// SwapETH is a paid mutator transaction binding the contract method 0x1114cd2a.
//
// Solidity: function swapETH(uint16 _dstChainId, address _refundAddress, bytes _toAddress, uint256 _amountLD, uint256 _minAmountLD) payable returns()
func (_Storage *StorageTransactorSession) SwapETH(_dstChainId uint16, _refundAddress common.Address, _toAddress []byte, _amountLD *big.Int, _minAmountLD *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapETH(&_Storage.TransactOpts, _dstChainId, _refundAddress, _toAddress, _amountLD, _minAmountLD)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactorSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}
