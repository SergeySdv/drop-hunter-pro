// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package syncswappoolmaster

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forwarderRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotWhitelistedFactory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"poolType\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RegisterPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"SetFactoryWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"UpdateFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newForwarderRegistry\",\"type\":\"address\"}],\"name\":\"UpdateForwarderRegistry\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forwarderRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isFactoryWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"poolType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"registerPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setFactoryWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForwarderRegistry\",\"type\":\"address\"}],\"name\":\"setForwarderRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Storage *StorageCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Storage *StorageSession) FeeManager() (common.Address, error) {
	return _Storage.Contract.FeeManager(&_Storage.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_Storage *StorageCallerSession) FeeManager() (common.Address, error) {
	return _Storage.Contract.FeeManager(&_Storage.CallOpts)
}

// ForwarderRegistry is a free data retrieval call binding the contract method 0x2b4c9f16.
//
// Solidity: function forwarderRegistry() view returns(address)
func (_Storage *StorageCaller) ForwarderRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "forwarderRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForwarderRegistry is a free data retrieval call binding the contract method 0x2b4c9f16.
//
// Solidity: function forwarderRegistry() view returns(address)
func (_Storage *StorageSession) ForwarderRegistry() (common.Address, error) {
	return _Storage.Contract.ForwarderRegistry(&_Storage.CallOpts)
}

// ForwarderRegistry is a free data retrieval call binding the contract method 0x2b4c9f16.
//
// Solidity: function forwarderRegistry() view returns(address)
func (_Storage *StorageCallerSession) ForwarderRegistry() (common.Address, error) {
	return _Storage.Contract.ForwarderRegistry(&_Storage.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address recipient)
func (_Storage *StorageCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address recipient)
func (_Storage *StorageSession) GetFeeRecipient() (common.Address, error) {
	return _Storage.Contract.GetFeeRecipient(&_Storage.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address recipient)
func (_Storage *StorageCallerSession) GetFeeRecipient() (common.Address, error) {
	return _Storage.Contract.GetFeeRecipient(&_Storage.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 ) view returns(address)
func (_Storage *StorageCaller) GetPool(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 ) view returns(address)
func (_Storage *StorageSession) GetPool(arg0 [32]byte) (common.Address, error) {
	return _Storage.Contract.GetPool(&_Storage.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 ) view returns(address)
func (_Storage *StorageCallerSession) GetPool(arg0 [32]byte) (common.Address, error) {
	return _Storage.Contract.GetPool(&_Storage.CallOpts, arg0)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address pool) view returns(uint24 fee)
func (_Storage *StorageCaller) GetProtocolFee(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getProtocolFee", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address pool) view returns(uint24 fee)
func (_Storage *StorageSession) GetProtocolFee(pool common.Address) (*big.Int, error) {
	return _Storage.Contract.GetProtocolFee(&_Storage.CallOpts, pool)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x0a992e0c.
//
// Solidity: function getProtocolFee(address pool) view returns(uint24 fee)
func (_Storage *StorageCallerSession) GetProtocolFee(pool common.Address) (*big.Int, error) {
	return _Storage.Contract.GetProtocolFee(&_Storage.CallOpts, pool)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 fee)
func (_Storage *StorageCaller) GetSwapFee(opts *bind.CallOpts, pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getSwapFee", pool, sender, tokenIn, tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 fee)
func (_Storage *StorageSession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _Storage.Contract.GetSwapFee(&_Storage.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 fee)
func (_Storage *StorageCallerSession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _Storage.Contract.GetSwapFee(&_Storage.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// IsFactoryWhitelisted is a free data retrieval call binding the contract method 0x43a0fcc4.
//
// Solidity: function isFactoryWhitelisted(address ) view returns(bool)
func (_Storage *StorageCaller) IsFactoryWhitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isFactoryWhitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFactoryWhitelisted is a free data retrieval call binding the contract method 0x43a0fcc4.
//
// Solidity: function isFactoryWhitelisted(address ) view returns(bool)
func (_Storage *StorageSession) IsFactoryWhitelisted(arg0 common.Address) (bool, error) {
	return _Storage.Contract.IsFactoryWhitelisted(&_Storage.CallOpts, arg0)
}

// IsFactoryWhitelisted is a free data retrieval call binding the contract method 0x43a0fcc4.
//
// Solidity: function isFactoryWhitelisted(address ) view returns(bool)
func (_Storage *StorageCallerSession) IsFactoryWhitelisted(arg0 common.Address) (bool, error) {
	return _Storage.Contract.IsFactoryWhitelisted(&_Storage.CallOpts, arg0)
}

// IsForwarder is a free data retrieval call binding the contract method 0xabcef554.
//
// Solidity: function isForwarder(address forwarder) view returns(bool)
func (_Storage *StorageCaller) IsForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForwarder is a free data retrieval call binding the contract method 0xabcef554.
//
// Solidity: function isForwarder(address forwarder) view returns(bool)
func (_Storage *StorageSession) IsForwarder(forwarder common.Address) (bool, error) {
	return _Storage.Contract.IsForwarder(&_Storage.CallOpts, forwarder)
}

// IsForwarder is a free data retrieval call binding the contract method 0xabcef554.
//
// Solidity: function isForwarder(address forwarder) view returns(bool)
func (_Storage *StorageCallerSession) IsForwarder(forwarder common.Address) (bool, error) {
	return _Storage.Contract.IsForwarder(&_Storage.CallOpts, forwarder)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_Storage *StorageCaller) IsPool(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isPool", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_Storage *StorageSession) IsPool(arg0 common.Address) (bool, error) {
	return _Storage.Contract.IsPool(&_Storage.CallOpts, arg0)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_Storage *StorageCallerSession) IsPool(arg0 common.Address) (bool, error) {
	return _Storage.Contract.IsPool(&_Storage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Storage *StorageCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Storage *StorageSession) PendingOwner() (common.Address, error) {
	return _Storage.Contract.PendingOwner(&_Storage.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Storage *StorageCallerSession) PendingOwner() (common.Address, error) {
	return _Storage.Contract.PendingOwner(&_Storage.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Storage *StorageCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Storage *StorageSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.Pools(&_Storage.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_Storage *StorageCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.Pools(&_Storage.CallOpts, arg0)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_Storage *StorageCaller) PoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "poolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_Storage *StorageSession) PoolsLength() (*big.Int, error) {
	return _Storage.Contract.PoolsLength(&_Storage.CallOpts)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_Storage *StorageCallerSession) PoolsLength() (*big.Int, error) {
	return _Storage.Contract.PoolsLength(&_Storage.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Storage *StorageCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Storage *StorageSession) Vault() (common.Address, error) {
	return _Storage.Contract.Vault(&_Storage.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Storage *StorageCallerSession) Vault() (common.Address, error) {
	return _Storage.Contract.Vault(&_Storage.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Storage *StorageTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Storage *StorageSession) AcceptOwnership() (*types.Transaction, error) {
	return _Storage.Contract.AcceptOwnership(&_Storage.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Storage *StorageTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Storage.Contract.AcceptOwnership(&_Storage.TransactOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) returns(address pool)
func (_Storage *StorageTransactor) CreatePool(opts *bind.TransactOpts, factory common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createPool", factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) returns(address pool)
func (_Storage *StorageSession) CreatePool(factory common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.CreatePool(&_Storage.TransactOpts, factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address factory, bytes data) returns(address pool)
func (_Storage *StorageTransactorSession) CreatePool(factory common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.CreatePool(&_Storage.TransactOpts, factory, data)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x784198d9.
//
// Solidity: function registerPool(address pool, uint16 poolType, bytes data) returns()
func (_Storage *StorageTransactor) RegisterPool(opts *bind.TransactOpts, pool common.Address, poolType uint16, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "registerPool", pool, poolType, data)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x784198d9.
//
// Solidity: function registerPool(address pool, uint16 poolType, bytes data) returns()
func (_Storage *StorageSession) RegisterPool(pool common.Address, poolType uint16, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.RegisterPool(&_Storage.TransactOpts, pool, poolType, data)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x784198d9.
//
// Solidity: function registerPool(address pool, uint16 poolType, bytes data) returns()
func (_Storage *StorageTransactorSession) RegisterPool(pool common.Address, poolType uint16, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.RegisterPool(&_Storage.TransactOpts, pool, poolType, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SetFactoryWhitelisted is a paid mutator transaction binding the contract method 0x0e16943b.
//
// Solidity: function setFactoryWhitelisted(address factory, bool whitelisted) returns()
func (_Storage *StorageTransactor) SetFactoryWhitelisted(opts *bind.TransactOpts, factory common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setFactoryWhitelisted", factory, whitelisted)
}

// SetFactoryWhitelisted is a paid mutator transaction binding the contract method 0x0e16943b.
//
// Solidity: function setFactoryWhitelisted(address factory, bool whitelisted) returns()
func (_Storage *StorageSession) SetFactoryWhitelisted(factory common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Storage.Contract.SetFactoryWhitelisted(&_Storage.TransactOpts, factory, whitelisted)
}

// SetFactoryWhitelisted is a paid mutator transaction binding the contract method 0x0e16943b.
//
// Solidity: function setFactoryWhitelisted(address factory, bool whitelisted) returns()
func (_Storage *StorageTransactorSession) SetFactoryWhitelisted(factory common.Address, whitelisted bool) (*types.Transaction, error) {
	return _Storage.Contract.SetFactoryWhitelisted(&_Storage.TransactOpts, factory, whitelisted)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_Storage *StorageTransactor) SetFeeManager(opts *bind.TransactOpts, newFeeManager common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setFeeManager", newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_Storage *StorageSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetFeeManager(&_Storage.TransactOpts, newFeeManager)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address newFeeManager) returns()
func (_Storage *StorageTransactorSession) SetFeeManager(newFeeManager common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetFeeManager(&_Storage.TransactOpts, newFeeManager)
}

// SetForwarderRegistry is a paid mutator transaction binding the contract method 0xfaff4f08.
//
// Solidity: function setForwarderRegistry(address newForwarderRegistry) returns()
func (_Storage *StorageTransactor) SetForwarderRegistry(opts *bind.TransactOpts, newForwarderRegistry common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setForwarderRegistry", newForwarderRegistry)
}

// SetForwarderRegistry is a paid mutator transaction binding the contract method 0xfaff4f08.
//
// Solidity: function setForwarderRegistry(address newForwarderRegistry) returns()
func (_Storage *StorageSession) SetForwarderRegistry(newForwarderRegistry common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetForwarderRegistry(&_Storage.TransactOpts, newForwarderRegistry)
}

// SetForwarderRegistry is a paid mutator transaction binding the contract method 0xfaff4f08.
//
// Solidity: function setForwarderRegistry(address newForwarderRegistry) returns()
func (_Storage *StorageTransactorSession) SetForwarderRegistry(newForwarderRegistry common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetForwarderRegistry(&_Storage.TransactOpts, newForwarderRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// StorageOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Storage contract.
type StorageOwnershipTransferStartedIterator struct {
	Event *StorageOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferStarted)
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
		it.Event = new(StorageOwnershipTransferStarted)
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
func (it *StorageOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Storage contract.
type StorageOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferStartedIterator{contract: _Storage.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferStarted)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) ParseOwnershipTransferStarted(log types.Log) (*StorageOwnershipTransferStarted, error) {
	event := new(StorageOwnershipTransferStarted)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRegisterPoolIterator is returned from FilterRegisterPool and is used to iterate over the raw logs and unpacked data for RegisterPool events raised by the Storage contract.
type StorageRegisterPoolIterator struct {
	Event *StorageRegisterPool // Event containing the contract specifics and raw log

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
func (it *StorageRegisterPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRegisterPool)
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
		it.Event = new(StorageRegisterPool)
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
func (it *StorageRegisterPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRegisterPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRegisterPool represents a RegisterPool event raised by the Storage contract.
type StorageRegisterPool struct {
	Factory  common.Address
	Pool     common.Address
	PoolType uint16
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegisterPool is a free log retrieval operation binding the contract event 0x4318beca5ca4f759b99c1f5f581fa8255b077a82e4c07f17213c471af5a0f56a.
//
// Solidity: event RegisterPool(address indexed factory, address indexed pool, uint16 indexed poolType, bytes data)
func (_Storage *StorageFilterer) FilterRegisterPool(opts *bind.FilterOpts, factory []common.Address, pool []common.Address, poolType []uint16) (*StorageRegisterPoolIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var poolTypeRule []interface{}
	for _, poolTypeItem := range poolType {
		poolTypeRule = append(poolTypeRule, poolTypeItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RegisterPool", factoryRule, poolRule, poolTypeRule)
	if err != nil {
		return nil, err
	}
	return &StorageRegisterPoolIterator{contract: _Storage.contract, event: "RegisterPool", logs: logs, sub: sub}, nil
}

// WatchRegisterPool is a free log subscription operation binding the contract event 0x4318beca5ca4f759b99c1f5f581fa8255b077a82e4c07f17213c471af5a0f56a.
//
// Solidity: event RegisterPool(address indexed factory, address indexed pool, uint16 indexed poolType, bytes data)
func (_Storage *StorageFilterer) WatchRegisterPool(opts *bind.WatchOpts, sink chan<- *StorageRegisterPool, factory []common.Address, pool []common.Address, poolType []uint16) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var poolTypeRule []interface{}
	for _, poolTypeItem := range poolType {
		poolTypeRule = append(poolTypeRule, poolTypeItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RegisterPool", factoryRule, poolRule, poolTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRegisterPool)
				if err := _Storage.contract.UnpackLog(event, "RegisterPool", log); err != nil {
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

// ParseRegisterPool is a log parse operation binding the contract event 0x4318beca5ca4f759b99c1f5f581fa8255b077a82e4c07f17213c471af5a0f56a.
//
// Solidity: event RegisterPool(address indexed factory, address indexed pool, uint16 indexed poolType, bytes data)
func (_Storage *StorageFilterer) ParseRegisterPool(log types.Log) (*StorageRegisterPool, error) {
	event := new(StorageRegisterPool)
	if err := _Storage.contract.UnpackLog(event, "RegisterPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetFactoryWhitelistedIterator is returned from FilterSetFactoryWhitelisted and is used to iterate over the raw logs and unpacked data for SetFactoryWhitelisted events raised by the Storage contract.
type StorageSetFactoryWhitelistedIterator struct {
	Event *StorageSetFactoryWhitelisted // Event containing the contract specifics and raw log

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
func (it *StorageSetFactoryWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetFactoryWhitelisted)
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
		it.Event = new(StorageSetFactoryWhitelisted)
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
func (it *StorageSetFactoryWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetFactoryWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetFactoryWhitelisted represents a SetFactoryWhitelisted event raised by the Storage contract.
type StorageSetFactoryWhitelisted struct {
	Factory     common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetFactoryWhitelisted is a free log retrieval operation binding the contract event 0x2c54fd9c1d7578da6d66f809aa31e327bc1a6cad6dc1d439f4b21adce023a156.
//
// Solidity: event SetFactoryWhitelisted(address indexed factory, bool whitelisted)
func (_Storage *StorageFilterer) FilterSetFactoryWhitelisted(opts *bind.FilterOpts, factory []common.Address) (*StorageSetFactoryWhitelistedIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetFactoryWhitelisted", factoryRule)
	if err != nil {
		return nil, err
	}
	return &StorageSetFactoryWhitelistedIterator{contract: _Storage.contract, event: "SetFactoryWhitelisted", logs: logs, sub: sub}, nil
}

// WatchSetFactoryWhitelisted is a free log subscription operation binding the contract event 0x2c54fd9c1d7578da6d66f809aa31e327bc1a6cad6dc1d439f4b21adce023a156.
//
// Solidity: event SetFactoryWhitelisted(address indexed factory, bool whitelisted)
func (_Storage *StorageFilterer) WatchSetFactoryWhitelisted(opts *bind.WatchOpts, sink chan<- *StorageSetFactoryWhitelisted, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetFactoryWhitelisted", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetFactoryWhitelisted)
				if err := _Storage.contract.UnpackLog(event, "SetFactoryWhitelisted", log); err != nil {
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

// ParseSetFactoryWhitelisted is a log parse operation binding the contract event 0x2c54fd9c1d7578da6d66f809aa31e327bc1a6cad6dc1d439f4b21adce023a156.
//
// Solidity: event SetFactoryWhitelisted(address indexed factory, bool whitelisted)
func (_Storage *StorageFilterer) ParseSetFactoryWhitelisted(log types.Log) (*StorageSetFactoryWhitelisted, error) {
	event := new(StorageSetFactoryWhitelisted)
	if err := _Storage.contract.UnpackLog(event, "SetFactoryWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpdateFeeManagerIterator is returned from FilterUpdateFeeManager and is used to iterate over the raw logs and unpacked data for UpdateFeeManager events raised by the Storage contract.
type StorageUpdateFeeManagerIterator struct {
	Event *StorageUpdateFeeManager // Event containing the contract specifics and raw log

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
func (it *StorageUpdateFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpdateFeeManager)
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
		it.Event = new(StorageUpdateFeeManager)
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
func (it *StorageUpdateFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpdateFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpdateFeeManager represents a UpdateFeeManager event raised by the Storage contract.
type StorageUpdateFeeManager struct {
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeManager is a free log retrieval operation binding the contract event 0x98a0dc993512fd2ddd1a4ee28a53d1275ec3c174565e996b03d4718909237bf8.
//
// Solidity: event UpdateFeeManager(address indexed newFeeManager)
func (_Storage *StorageFilterer) FilterUpdateFeeManager(opts *bind.FilterOpts, newFeeManager []common.Address) (*StorageUpdateFeeManagerIterator, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UpdateFeeManager", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpdateFeeManagerIterator{contract: _Storage.contract, event: "UpdateFeeManager", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeManager is a free log subscription operation binding the contract event 0x98a0dc993512fd2ddd1a4ee28a53d1275ec3c174565e996b03d4718909237bf8.
//
// Solidity: event UpdateFeeManager(address indexed newFeeManager)
func (_Storage *StorageFilterer) WatchUpdateFeeManager(opts *bind.WatchOpts, sink chan<- *StorageUpdateFeeManager, newFeeManager []common.Address) (event.Subscription, error) {

	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UpdateFeeManager", newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpdateFeeManager)
				if err := _Storage.contract.UnpackLog(event, "UpdateFeeManager", log); err != nil {
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

// ParseUpdateFeeManager is a log parse operation binding the contract event 0x98a0dc993512fd2ddd1a4ee28a53d1275ec3c174565e996b03d4718909237bf8.
//
// Solidity: event UpdateFeeManager(address indexed newFeeManager)
func (_Storage *StorageFilterer) ParseUpdateFeeManager(log types.Log) (*StorageUpdateFeeManager, error) {
	event := new(StorageUpdateFeeManager)
	if err := _Storage.contract.UnpackLog(event, "UpdateFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpdateForwarderRegistryIterator is returned from FilterUpdateForwarderRegistry and is used to iterate over the raw logs and unpacked data for UpdateForwarderRegistry events raised by the Storage contract.
type StorageUpdateForwarderRegistryIterator struct {
	Event *StorageUpdateForwarderRegistry // Event containing the contract specifics and raw log

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
func (it *StorageUpdateForwarderRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpdateForwarderRegistry)
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
		it.Event = new(StorageUpdateForwarderRegistry)
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
func (it *StorageUpdateForwarderRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpdateForwarderRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpdateForwarderRegistry represents a UpdateForwarderRegistry event raised by the Storage contract.
type StorageUpdateForwarderRegistry struct {
	NewForwarderRegistry common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateForwarderRegistry is a free log retrieval operation binding the contract event 0x41e309325204d4979853ee58efce248d83b12200f22fe333be8e33bbf748a71b.
//
// Solidity: event UpdateForwarderRegistry(address indexed newForwarderRegistry)
func (_Storage *StorageFilterer) FilterUpdateForwarderRegistry(opts *bind.FilterOpts, newForwarderRegistry []common.Address) (*StorageUpdateForwarderRegistryIterator, error) {

	var newForwarderRegistryRule []interface{}
	for _, newForwarderRegistryItem := range newForwarderRegistry {
		newForwarderRegistryRule = append(newForwarderRegistryRule, newForwarderRegistryItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "UpdateForwarderRegistry", newForwarderRegistryRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpdateForwarderRegistryIterator{contract: _Storage.contract, event: "UpdateForwarderRegistry", logs: logs, sub: sub}, nil
}

// WatchUpdateForwarderRegistry is a free log subscription operation binding the contract event 0x41e309325204d4979853ee58efce248d83b12200f22fe333be8e33bbf748a71b.
//
// Solidity: event UpdateForwarderRegistry(address indexed newForwarderRegistry)
func (_Storage *StorageFilterer) WatchUpdateForwarderRegistry(opts *bind.WatchOpts, sink chan<- *StorageUpdateForwarderRegistry, newForwarderRegistry []common.Address) (event.Subscription, error) {

	var newForwarderRegistryRule []interface{}
	for _, newForwarderRegistryItem := range newForwarderRegistry {
		newForwarderRegistryRule = append(newForwarderRegistryRule, newForwarderRegistryItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "UpdateForwarderRegistry", newForwarderRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpdateForwarderRegistry)
				if err := _Storage.contract.UnpackLog(event, "UpdateForwarderRegistry", log); err != nil {
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

// ParseUpdateForwarderRegistry is a log parse operation binding the contract event 0x41e309325204d4979853ee58efce248d83b12200f22fe333be8e33bbf748a71b.
//
// Solidity: event UpdateForwarderRegistry(address indexed newForwarderRegistry)
func (_Storage *StorageFilterer) ParseUpdateForwarderRegistry(log types.Log) (*StorageUpdateForwarderRegistry, error) {
	event := new(StorageUpdateForwarderRegistry)
	if err := _Storage.contract.UnpackLog(event, "UpdateForwarderRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
