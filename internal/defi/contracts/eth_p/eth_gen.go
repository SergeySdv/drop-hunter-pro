// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_p

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
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_numOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetSpentToday\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_h\",\"type\":\"bytes32\"}],\"name\":\"confirm\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newLimit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"name\":\"_r\",\"type\":\"bytes32\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operation\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRequired\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operation\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"hasConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"kill\",\"outputs\":[],\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"m_dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_daylimit\",\"type\":\"uint256\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"OwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newRequirement\",\"type\":\"uint256\"}],\"name\":\"RequirementChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SingleTransact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MultiTransact\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operation\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ConfirmationNeeded\",\"type\":\"event\"}]",
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

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) returns(bool)
func (_Storage *StorageCaller) HasConfirmed(opts *bind.CallOpts, _operation [32]byte, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasConfirmed", _operation, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) returns(bool)
func (_Storage *StorageSession) HasConfirmed(_operation [32]byte, _owner common.Address) (bool, error) {
	return _Storage.Contract.HasConfirmed(&_Storage.CallOpts, _operation, _owner)
}

// HasConfirmed is a free data retrieval call binding the contract method 0xc2cf7326.
//
// Solidity: function hasConfirmed(bytes32 _operation, address _owner) returns(bool)
func (_Storage *StorageCallerSession) HasConfirmed(_operation [32]byte, _owner common.Address) (bool, error) {
	return _Storage.Contract.HasConfirmed(&_Storage.CallOpts, _operation, _owner)
}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() returns(uint256)
func (_Storage *StorageCaller) MDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "m_dailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() returns(uint256)
func (_Storage *StorageSession) MDailyLimit() (*big.Int, error) {
	return _Storage.Contract.MDailyLimit(&_Storage.CallOpts)
}

// MDailyLimit is a free data retrieval call binding the contract method 0xf1736d86.
//
// Solidity: function m_dailyLimit() returns(uint256)
func (_Storage *StorageCallerSession) MDailyLimit() (*big.Int, error) {
	return _Storage.Contract.MDailyLimit(&_Storage.CallOpts)
}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() returns(uint256)
func (_Storage *StorageCaller) MNumOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "m_numOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() returns(uint256)
func (_Storage *StorageSession) MNumOwners() (*big.Int, error) {
	return _Storage.Contract.MNumOwners(&_Storage.CallOpts)
}

// MNumOwners is a free data retrieval call binding the contract method 0x4123cb6b.
//
// Solidity: function m_numOwners() returns(uint256)
func (_Storage *StorageCallerSession) MNumOwners() (*big.Int, error) {
	return _Storage.Contract.MNumOwners(&_Storage.CallOpts)
}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() returns(uint256)
func (_Storage *StorageCaller) MRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "m_required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() returns(uint256)
func (_Storage *StorageSession) MRequired() (*big.Int, error) {
	return _Storage.Contract.MRequired(&_Storage.CallOpts)
}

// MRequired is a free data retrieval call binding the contract method 0x746c9171.
//
// Solidity: function m_required() returns(uint256)
func (_Storage *StorageCallerSession) MRequired() (*big.Int, error) {
	return _Storage.Contract.MRequired(&_Storage.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_Storage *StorageTransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addOwner", _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_Storage *StorageSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddOwner(&_Storage.TransactOpts, _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address _owner) returns()
func (_Storage *StorageTransactorSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddOwner(&_Storage.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_Storage *StorageTransactor) ChangeOwner(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeOwner", _from, _to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_Storage *StorageSession) ChangeOwner(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ChangeOwner(&_Storage.TransactOpts, _from, _to)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _from, address _to) returns()
func (_Storage *StorageTransactorSession) ChangeOwner(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.ChangeOwner(&_Storage.TransactOpts, _from, _to)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_Storage *StorageTransactor) ChangeRequirement(opts *bind.TransactOpts, _newRequired *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "changeRequirement", _newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_Storage *StorageSession) ChangeRequirement(_newRequired *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeRequirement(&_Storage.TransactOpts, _newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _newRequired) returns()
func (_Storage *StorageTransactorSession) ChangeRequirement(_newRequired *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.ChangeRequirement(&_Storage.TransactOpts, _newRequired)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_Storage *StorageTransactor) Confirm(opts *bind.TransactOpts, _h [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "confirm", _h)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_Storage *StorageSession) Confirm(_h [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.Confirm(&_Storage.TransactOpts, _h)
}

// Confirm is a paid mutator transaction binding the contract method 0x797af627.
//
// Solidity: function confirm(bytes32 _h) returns(bool)
func (_Storage *StorageTransactorSession) Confirm(_h [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.Confirm(&_Storage.TransactOpts, _h)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_Storage *StorageTransactor) Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "execute", _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_Storage *StorageSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Execute(&_Storage.TransactOpts, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bytes32 _r)
func (_Storage *StorageTransactorSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Execute(&_Storage.TransactOpts, _to, _value, _data)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_Storage *StorageTransactor) IsOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "isOwner", _addr)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_Storage *StorageSession) IsOwner(_addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.IsOwner(&_Storage.TransactOpts, _addr)
}

// IsOwner is a paid mutator transaction binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address _addr) returns(bool)
func (_Storage *StorageTransactorSession) IsOwner(_addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.IsOwner(&_Storage.TransactOpts, _addr)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_Storage *StorageTransactor) Kill(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "kill", _to)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_Storage *StorageSession) Kill(_to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts, _to)
}

// Kill is a paid mutator transaction binding the contract method 0xcbf0b0c0.
//
// Solidity: function kill(address _to) returns()
func (_Storage *StorageTransactorSession) Kill(_to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts, _to)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_Storage *StorageTransactor) RemoveOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeOwner", _owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_Storage *StorageSession) RemoveOwner(_owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RemoveOwner(&_Storage.TransactOpts, _owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address _owner) returns()
func (_Storage *StorageTransactorSession) RemoveOwner(_owner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RemoveOwner(&_Storage.TransactOpts, _owner)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_Storage *StorageTransactor) ResetSpentToday(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "resetSpentToday")
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_Storage *StorageSession) ResetSpentToday() (*types.Transaction, error) {
	return _Storage.Contract.ResetSpentToday(&_Storage.TransactOpts)
}

// ResetSpentToday is a paid mutator transaction binding the contract method 0x5c52c2f5.
//
// Solidity: function resetSpentToday() returns()
func (_Storage *StorageTransactorSession) ResetSpentToday() (*types.Transaction, error) {
	return _Storage.Contract.ResetSpentToday(&_Storage.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_Storage *StorageTransactor) Revoke(opts *bind.TransactOpts, _operation [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "revoke", _operation)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_Storage *StorageSession) Revoke(_operation [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.Revoke(&_Storage.TransactOpts, _operation)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 _operation) returns()
func (_Storage *StorageTransactorSession) Revoke(_operation [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.Revoke(&_Storage.TransactOpts, _operation)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_Storage *StorageTransactor) SetDailyLimit(opts *bind.TransactOpts, _newLimit *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDailyLimit", _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_Storage *StorageSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDailyLimit(&_Storage.TransactOpts, _newLimit)
}

// SetDailyLimit is a paid mutator transaction binding the contract method 0xb20d30a9.
//
// Solidity: function setDailyLimit(uint256 _newLimit) returns()
func (_Storage *StorageTransactorSession) SetDailyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDailyLimit(&_Storage.TransactOpts, _newLimit)
}

// StorageConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the Storage contract.
type StorageConfirmationIterator struct {
	Event *StorageConfirmation // Event containing the contract specifics and raw log

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
func (it *StorageConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageConfirmation)
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
		it.Event = new(StorageConfirmation)
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
func (it *StorageConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageConfirmation represents a Confirmation event raised by the Storage contract.
type StorageConfirmation struct {
	Owner     common.Address
	Operation [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0xe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda.
//
// Solidity: event Confirmation(address owner, bytes32 operation)
func (_Storage *StorageFilterer) FilterConfirmation(opts *bind.FilterOpts) (*StorageConfirmationIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return &StorageConfirmationIterator{contract: _Storage.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0xe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda.
//
// Solidity: event Confirmation(address owner, bytes32 operation)
func (_Storage *StorageFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *StorageConfirmation) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageConfirmation)
				if err := _Storage.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// ParseConfirmation is a log parse operation binding the contract event 0xe1c52dc63b719ade82e8bea94cc41a0d5d28e4aaf536adb5e9cccc9ff8c1aeda.
//
// Solidity: event Confirmation(address owner, bytes32 operation)
func (_Storage *StorageFilterer) ParseConfirmation(log types.Log) (*StorageConfirmation, error) {
	event := new(StorageConfirmation)
	if err := _Storage.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageConfirmationNeededIterator is returned from FilterConfirmationNeeded and is used to iterate over the raw logs and unpacked data for ConfirmationNeeded events raised by the Storage contract.
type StorageConfirmationNeededIterator struct {
	Event *StorageConfirmationNeeded // Event containing the contract specifics and raw log

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
func (it *StorageConfirmationNeededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageConfirmationNeeded)
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
		it.Event = new(StorageConfirmationNeeded)
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
func (it *StorageConfirmationNeededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageConfirmationNeededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageConfirmationNeeded represents a ConfirmationNeeded event raised by the Storage contract.
type StorageConfirmationNeeded struct {
	Operation [32]byte
	Initiator common.Address
	Value     *big.Int
	To        common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationNeeded is a free log retrieval operation binding the contract event 0x1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf32.
//
// Solidity: event ConfirmationNeeded(bytes32 operation, address initiator, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) FilterConfirmationNeeded(opts *bind.FilterOpts) (*StorageConfirmationNeededIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ConfirmationNeeded")
	if err != nil {
		return nil, err
	}
	return &StorageConfirmationNeededIterator{contract: _Storage.contract, event: "ConfirmationNeeded", logs: logs, sub: sub}, nil
}

// WatchConfirmationNeeded is a free log subscription operation binding the contract event 0x1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf32.
//
// Solidity: event ConfirmationNeeded(bytes32 operation, address initiator, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) WatchConfirmationNeeded(opts *bind.WatchOpts, sink chan<- *StorageConfirmationNeeded) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ConfirmationNeeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageConfirmationNeeded)
				if err := _Storage.contract.UnpackLog(event, "ConfirmationNeeded", log); err != nil {
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

// ParseConfirmationNeeded is a log parse operation binding the contract event 0x1733cbb53659d713b79580f79f3f9ff215f78a7c7aa45890f3b89fc5cddfbf32.
//
// Solidity: event ConfirmationNeeded(bytes32 operation, address initiator, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) ParseConfirmationNeeded(log types.Log) (*StorageConfirmationNeeded, error) {
	event := new(StorageConfirmationNeeded)
	if err := _Storage.contract.UnpackLog(event, "ConfirmationNeeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Storage contract.
type StorageDepositIterator struct {
	Event *StorageDeposit // Event containing the contract specifics and raw log

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
func (it *StorageDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDeposit)
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
		it.Event = new(StorageDeposit)
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
func (it *StorageDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDeposit represents a Deposit event raised by the Storage contract.
type StorageDeposit struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _from, uint256 value)
func (_Storage *StorageFilterer) FilterDeposit(opts *bind.FilterOpts) (*StorageDepositIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StorageDepositIterator{contract: _Storage.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _from, uint256 value)
func (_Storage *StorageFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StorageDeposit) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDeposit)
				if err := _Storage.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _from, uint256 value)
func (_Storage *StorageFilterer) ParseDeposit(log types.Log) (*StorageDeposit, error) {
	event := new(StorageDeposit)
	if err := _Storage.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMultiTransactIterator is returned from FilterMultiTransact and is used to iterate over the raw logs and unpacked data for MultiTransact events raised by the Storage contract.
type StorageMultiTransactIterator struct {
	Event *StorageMultiTransact // Event containing the contract specifics and raw log

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
func (it *StorageMultiTransactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMultiTransact)
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
		it.Event = new(StorageMultiTransact)
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
func (it *StorageMultiTransactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMultiTransactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMultiTransact represents a MultiTransact event raised by the Storage contract.
type StorageMultiTransact struct {
	Owner     common.Address
	Operation [32]byte
	Value     *big.Int
	To        common.Address
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMultiTransact is a free log retrieval operation binding the contract event 0xe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a.
//
// Solidity: event MultiTransact(address owner, bytes32 operation, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) FilterMultiTransact(opts *bind.FilterOpts) (*StorageMultiTransactIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MultiTransact")
	if err != nil {
		return nil, err
	}
	return &StorageMultiTransactIterator{contract: _Storage.contract, event: "MultiTransact", logs: logs, sub: sub}, nil
}

// WatchMultiTransact is a free log subscription operation binding the contract event 0xe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a.
//
// Solidity: event MultiTransact(address owner, bytes32 operation, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) WatchMultiTransact(opts *bind.WatchOpts, sink chan<- *StorageMultiTransact) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MultiTransact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMultiTransact)
				if err := _Storage.contract.UnpackLog(event, "MultiTransact", log); err != nil {
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

// ParseMultiTransact is a log parse operation binding the contract event 0xe7c957c06e9a662c1a6c77366179f5b702b97651dc28eee7d5bf1dff6e40bb4a.
//
// Solidity: event MultiTransact(address owner, bytes32 operation, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) ParseMultiTransact(log types.Log) (*StorageMultiTransact, error) {
	event := new(StorageMultiTransact)
	if err := _Storage.contract.UnpackLog(event, "MultiTransact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnerAddedIterator is returned from FilterOwnerAdded and is used to iterate over the raw logs and unpacked data for OwnerAdded events raised by the Storage contract.
type StorageOwnerAddedIterator struct {
	Event *StorageOwnerAdded // Event containing the contract specifics and raw log

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
func (it *StorageOwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnerAdded)
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
		it.Event = new(StorageOwnerAdded)
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
func (it *StorageOwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnerAdded represents a OwnerAdded event raised by the Storage contract.
type StorageOwnerAdded struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerAdded is a free log retrieval operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address newOwner)
func (_Storage *StorageFilterer) FilterOwnerAdded(opts *bind.FilterOpts) (*StorageOwnerAddedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnerAdded")
	if err != nil {
		return nil, err
	}
	return &StorageOwnerAddedIterator{contract: _Storage.contract, event: "OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchOwnerAdded is a free log subscription operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address newOwner)
func (_Storage *StorageFilterer) WatchOwnerAdded(opts *bind.WatchOpts, sink chan<- *StorageOwnerAdded) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnerAdded)
				if err := _Storage.contract.UnpackLog(event, "OwnerAdded", log); err != nil {
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

// ParseOwnerAdded is a log parse operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address newOwner)
func (_Storage *StorageFilterer) ParseOwnerAdded(log types.Log) (*StorageOwnerAdded, error) {
	event := new(StorageOwnerAdded)
	if err := _Storage.contract.UnpackLog(event, "OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Storage contract.
type StorageOwnerChangedIterator struct {
	Event *StorageOwnerChanged // Event containing the contract specifics and raw log

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
func (it *StorageOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnerChanged)
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
		it.Event = new(StorageOwnerChanged)
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
func (it *StorageOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnerChanged represents a OwnerChanged event raised by the Storage contract.
type StorageOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Storage *StorageFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*StorageOwnerChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &StorageOwnerChangedIterator{contract: _Storage.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Storage *StorageFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *StorageOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnerChanged)
				if err := _Storage.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Storage *StorageFilterer) ParseOwnerChanged(log types.Log) (*StorageOwnerChanged, error) {
	event := new(StorageOwnerChanged)
	if err := _Storage.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnerRemovedIterator is returned from FilterOwnerRemoved and is used to iterate over the raw logs and unpacked data for OwnerRemoved events raised by the Storage contract.
type StorageOwnerRemovedIterator struct {
	Event *StorageOwnerRemoved // Event containing the contract specifics and raw log

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
func (it *StorageOwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnerRemoved)
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
		it.Event = new(StorageOwnerRemoved)
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
func (it *StorageOwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnerRemoved represents a OwnerRemoved event raised by the Storage contract.
type StorageOwnerRemoved struct {
	OldOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoved is a free log retrieval operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address oldOwner)
func (_Storage *StorageFilterer) FilterOwnerRemoved(opts *bind.FilterOpts) (*StorageOwnerRemovedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnerRemoved")
	if err != nil {
		return nil, err
	}
	return &StorageOwnerRemovedIterator{contract: _Storage.contract, event: "OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoved is a free log subscription operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address oldOwner)
func (_Storage *StorageFilterer) WatchOwnerRemoved(opts *bind.WatchOpts, sink chan<- *StorageOwnerRemoved) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnerRemoved)
				if err := _Storage.contract.UnpackLog(event, "OwnerRemoved", log); err != nil {
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

// ParseOwnerRemoved is a log parse operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address oldOwner)
func (_Storage *StorageFilterer) ParseOwnerRemoved(log types.Log) (*StorageOwnerRemoved, error) {
	event := new(StorageOwnerRemoved)
	if err := _Storage.contract.UnpackLog(event, "OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRequirementChangedIterator is returned from FilterRequirementChanged and is used to iterate over the raw logs and unpacked data for RequirementChanged events raised by the Storage contract.
type StorageRequirementChangedIterator struct {
	Event *StorageRequirementChanged // Event containing the contract specifics and raw log

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
func (it *StorageRequirementChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRequirementChanged)
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
		it.Event = new(StorageRequirementChanged)
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
func (it *StorageRequirementChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRequirementChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRequirementChanged represents a RequirementChanged event raised by the Storage contract.
type StorageRequirementChanged struct {
	NewRequirement *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequirementChanged is a free log retrieval operation binding the contract event 0xacbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da.
//
// Solidity: event RequirementChanged(uint256 newRequirement)
func (_Storage *StorageFilterer) FilterRequirementChanged(opts *bind.FilterOpts) (*StorageRequirementChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RequirementChanged")
	if err != nil {
		return nil, err
	}
	return &StorageRequirementChangedIterator{contract: _Storage.contract, event: "RequirementChanged", logs: logs, sub: sub}, nil
}

// WatchRequirementChanged is a free log subscription operation binding the contract event 0xacbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da.
//
// Solidity: event RequirementChanged(uint256 newRequirement)
func (_Storage *StorageFilterer) WatchRequirementChanged(opts *bind.WatchOpts, sink chan<- *StorageRequirementChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RequirementChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRequirementChanged)
				if err := _Storage.contract.UnpackLog(event, "RequirementChanged", log); err != nil {
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

// ParseRequirementChanged is a log parse operation binding the contract event 0xacbdb084c721332ac59f9b8e392196c9eb0e4932862da8eb9beaf0dad4f550da.
//
// Solidity: event RequirementChanged(uint256 newRequirement)
func (_Storage *StorageFilterer) ParseRequirementChanged(log types.Log) (*StorageRequirementChanged, error) {
	event := new(StorageRequirementChanged)
	if err := _Storage.contract.UnpackLog(event, "RequirementChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the Storage contract.
type StorageRevokeIterator struct {
	Event *StorageRevoke // Event containing the contract specifics and raw log

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
func (it *StorageRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRevoke)
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
		it.Event = new(StorageRevoke)
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
func (it *StorageRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRevoke represents a Revoke event raised by the Storage contract.
type StorageRevoke struct {
	Owner     common.Address
	Operation [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address owner, bytes32 operation)
func (_Storage *StorageFilterer) FilterRevoke(opts *bind.FilterOpts) (*StorageRevokeIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return &StorageRevokeIterator{contract: _Storage.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address owner, bytes32 operation)
func (_Storage *StorageFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *StorageRevoke) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRevoke)
				if err := _Storage.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address owner, bytes32 operation)
func (_Storage *StorageFilterer) ParseRevoke(log types.Log) (*StorageRevoke, error) {
	event := new(StorageRevoke)
	if err := _Storage.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSingleTransactIterator is returned from FilterSingleTransact and is used to iterate over the raw logs and unpacked data for SingleTransact events raised by the Storage contract.
type StorageSingleTransactIterator struct {
	Event *StorageSingleTransact // Event containing the contract specifics and raw log

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
func (it *StorageSingleTransactIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSingleTransact)
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
		it.Event = new(StorageSingleTransact)
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
func (it *StorageSingleTransactIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSingleTransactIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSingleTransact represents a SingleTransact event raised by the Storage contract.
type StorageSingleTransact struct {
	Owner common.Address
	Value *big.Int
	To    common.Address
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSingleTransact is a free log retrieval operation binding the contract event 0x92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd004.
//
// Solidity: event SingleTransact(address owner, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) FilterSingleTransact(opts *bind.FilterOpts) (*StorageSingleTransactIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SingleTransact")
	if err != nil {
		return nil, err
	}
	return &StorageSingleTransactIterator{contract: _Storage.contract, event: "SingleTransact", logs: logs, sub: sub}, nil
}

// WatchSingleTransact is a free log subscription operation binding the contract event 0x92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd004.
//
// Solidity: event SingleTransact(address owner, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) WatchSingleTransact(opts *bind.WatchOpts, sink chan<- *StorageSingleTransact) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SingleTransact")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSingleTransact)
				if err := _Storage.contract.UnpackLog(event, "SingleTransact", log); err != nil {
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

// ParseSingleTransact is a log parse operation binding the contract event 0x92ca3a80853e6663fa31fa10b99225f18d4902939b4c53a9caae9043f6efd004.
//
// Solidity: event SingleTransact(address owner, uint256 value, address to, bytes data)
func (_Storage *StorageFilterer) ParseSingleTransact(log types.Log) (*StorageSingleTransact, error) {
	event := new(StorageSingleTransact)
	if err := _Storage.contract.UnpackLog(event, "SingleTransact", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
