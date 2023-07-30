// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package syncswapvault

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashLoanFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashLoanFeePercentage\",\"type\":\"uint256\"}],\"name\":\"FlashLoanFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC3156_CALLBACK_SUCCESS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoanMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFlashLoanFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_flashLoanFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFlashLoanFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"withdrawAlternative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Storage *StorageCaller) ERC3156CALLBACKSUCCESS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "ERC3156_CALLBACK_SUCCESS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Storage *StorageSession) ERC3156CALLBACKSUCCESS() ([32]byte, error) {
	return _Storage.Contract.ERC3156CALLBACKSUCCESS(&_Storage.CallOpts)
}

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Storage *StorageCallerSession) ERC3156CALLBACKSUCCESS() ([32]byte, error) {
	return _Storage.Contract.ERC3156CALLBACKSUCCESS(&_Storage.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Storage *StorageCaller) BalanceOf(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "balanceOf", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Storage *StorageSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, token, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Storage *StorageCallerSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Storage.Contract.BalanceOf(&_Storage.CallOpts, token, account)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Storage *StorageCaller) FlashFee(opts *bind.CallOpts, arg0 common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "flashFee", arg0, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Storage *StorageSession) FlashFee(arg0 common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.FlashFee(&_Storage.CallOpts, arg0, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Storage *StorageCallerSession) FlashFee(arg0 common.Address, amount *big.Int) (*big.Int, error) {
	return _Storage.Contract.FlashFee(&_Storage.CallOpts, arg0, amount)
}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Storage *StorageCaller) FlashLoanFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "flashLoanFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Storage *StorageSession) FlashLoanFeePercentage() (*big.Int, error) {
	return _Storage.Contract.FlashLoanFeePercentage(&_Storage.CallOpts)
}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Storage *StorageCallerSession) FlashLoanFeePercentage() (*big.Int, error) {
	return _Storage.Contract.FlashLoanFeePercentage(&_Storage.CallOpts)
}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Storage *StorageCaller) FlashLoanFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "flashLoanFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Storage *StorageSession) FlashLoanFeeRecipient() (common.Address, error) {
	return _Storage.Contract.FlashLoanFeeRecipient(&_Storage.CallOpts)
}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Storage *StorageCallerSession) FlashLoanFeeRecipient() (common.Address, error) {
	return _Storage.Contract.FlashLoanFeeRecipient(&_Storage.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Storage *StorageCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Storage *StorageSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Storage.Contract.MaxFlashLoan(&_Storage.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Storage *StorageCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Storage.Contract.MaxFlashLoan(&_Storage.CallOpts, token)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageSession) Paused() (bool, error) {
	return _Storage.Contract.Paused(&_Storage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Storage *StorageCallerSession) Paused() (bool, error) {
	return _Storage.Contract.Paused(&_Storage.CallOpts)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Storage *StorageCaller) Reserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "reserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Storage *StorageSession) Reserves(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.Reserves(&_Storage.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Storage *StorageCallerSession) Reserves(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.Reserves(&_Storage.CallOpts, arg0)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Storage *StorageCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Storage *StorageSession) WETH() (common.Address, error) {
	return _Storage.Contract.WETH(&_Storage.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Storage *StorageCallerSession) WETH() (common.Address, error) {
	return _Storage.Contract.WETH(&_Storage.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Storage *StorageTransactor) Deposit(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deposit", token, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Storage *StorageSession) Deposit(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Deposit(&_Storage.TransactOpts, token, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Storage *StorageTransactorSession) Deposit(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Deposit(&_Storage.TransactOpts, token, to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Storage *StorageTransactor) DepositETH(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "depositETH", to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Storage *StorageSession) DepositETH(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DepositETH(&_Storage.TransactOpts, to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Storage *StorageTransactorSession) DepositETH(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.DepositETH(&_Storage.TransactOpts, to)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Storage *StorageTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "flashLoan", receiver, token, amount, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Storage *StorageSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.Contract.FlashLoan(&_Storage.TransactOpts, receiver, token, amount, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Storage *StorageTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.Contract.FlashLoan(&_Storage.TransactOpts, receiver, token, amount, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Storage *StorageTransactor) FlashLoanMultiple(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "flashLoanMultiple", recipient, tokens, amounts, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Storage *StorageSession) FlashLoanMultiple(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.Contract.FlashLoanMultiple(&_Storage.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Storage *StorageTransactorSession) FlashLoanMultiple(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Storage.Contract.FlashLoanMultiple(&_Storage.TransactOpts, recipient, tokens, amounts, userData)
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

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Storage *StorageTransactor) SetFlashLoanFeePercentage(opts *bind.TransactOpts, newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setFlashLoanFeePercentage", newFlashLoanFeePercentage)
}

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Storage *StorageSession) SetFlashLoanFeePercentage(newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetFlashLoanFeePercentage(&_Storage.TransactOpts, newFlashLoanFeePercentage)
}

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Storage *StorageTransactorSession) SetFlashLoanFeePercentage(newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetFlashLoanFeePercentage(&_Storage.TransactOpts, newFlashLoanFeePercentage)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Storage *StorageTransactor) SetFlashLoanFeeRecipient(opts *bind.TransactOpts, _flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setFlashLoanFeeRecipient", _flashLoanFeeRecipient)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Storage *StorageSession) SetFlashLoanFeeRecipient(_flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetFlashLoanFeeRecipient(&_Storage.TransactOpts, _flashLoanFeeRecipient)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Storage *StorageTransactorSession) SetFlashLoanFeeRecipient(_flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetFlashLoanFeeRecipient(&_Storage.TransactOpts, _flashLoanFeeRecipient)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Storage *StorageTransactor) SetPaused(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setPaused", _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Storage *StorageSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPaused(&_Storage.TransactOpts, _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Storage *StorageTransactorSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPaused(&_Storage.TransactOpts, _status)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transfer", token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Storage *StorageSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Transfer(&_Storage.TransactOpts, token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Storage *StorageTransactor) TransferAndDeposit(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferAndDeposit", token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Storage *StorageSession) TransferAndDeposit(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferAndDeposit(&_Storage.TransactOpts, token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Storage *StorageTransactorSession) TransferAndDeposit(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TransferAndDeposit(&_Storage.TransactOpts, token, to, amount)
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

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Withdraw(&_Storage.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Withdraw(&_Storage.TransactOpts, token, to, amount)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Storage *StorageTransactor) WithdrawAlternative(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawAlternative", token, to, amount, mode)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Storage *StorageSession) WithdrawAlternative(token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAlternative(&_Storage.TransactOpts, token, to, amount, mode)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Storage *StorageTransactorSession) WithdrawAlternative(token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawAlternative(&_Storage.TransactOpts, token, to, amount, mode)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Storage *StorageTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Storage *StorageSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawETH(&_Storage.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawETH(&_Storage.TransactOpts, to, amount)
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

// StorageFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Storage contract.
type StorageFlashLoanIterator struct {
	Event *StorageFlashLoan // Event containing the contract specifics and raw log

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
func (it *StorageFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFlashLoan)
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
		it.Event = new(StorageFlashLoan)
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
func (it *StorageFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFlashLoan represents a FlashLoan event raised by the Storage contract.
type StorageFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Storage *StorageFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*StorageFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StorageFlashLoanIterator{contract: _Storage.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Storage *StorageFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *StorageFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFlashLoan)
				if err := _Storage.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Storage *StorageFilterer) ParseFlashLoan(log types.Log) (*StorageFlashLoan, error) {
	event := new(StorageFlashLoan)
	if err := _Storage.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageFlashLoanFeePercentageChangedIterator is returned from FilterFlashLoanFeePercentageChanged and is used to iterate over the raw logs and unpacked data for FlashLoanFeePercentageChanged events raised by the Storage contract.
type StorageFlashLoanFeePercentageChangedIterator struct {
	Event *StorageFlashLoanFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *StorageFlashLoanFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFlashLoanFeePercentageChanged)
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
		it.Event = new(StorageFlashLoanFeePercentageChanged)
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
func (it *StorageFlashLoanFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFlashLoanFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFlashLoanFeePercentageChanged represents a FlashLoanFeePercentageChanged event raised by the Storage contract.
type StorageFlashLoanFeePercentageChanged struct {
	OldFlashLoanFeePercentage *big.Int
	NewFlashLoanFeePercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoanFeePercentageChanged is a free log retrieval operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Storage *StorageFilterer) FilterFlashLoanFeePercentageChanged(opts *bind.FilterOpts) (*StorageFlashLoanFeePercentageChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FlashLoanFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &StorageFlashLoanFeePercentageChangedIterator{contract: _Storage.contract, event: "FlashLoanFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchFlashLoanFeePercentageChanged is a free log subscription operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Storage *StorageFilterer) WatchFlashLoanFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *StorageFlashLoanFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FlashLoanFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFlashLoanFeePercentageChanged)
				if err := _Storage.contract.UnpackLog(event, "FlashLoanFeePercentageChanged", log); err != nil {
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

// ParseFlashLoanFeePercentageChanged is a log parse operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Storage *StorageFilterer) ParseFlashLoanFeePercentageChanged(log types.Log) (*StorageFlashLoanFeePercentageChanged, error) {
	event := new(StorageFlashLoanFeePercentageChanged)
	if err := _Storage.contract.UnpackLog(event, "FlashLoanFeePercentageChanged", log); err != nil {
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

// StoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Storage contract.
type StoragePausedIterator struct {
	Event *StoragePaused // Event containing the contract specifics and raw log

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
func (it *StoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePaused)
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
		it.Event = new(StoragePaused)
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
func (it *StoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePaused represents a Paused event raised by the Storage contract.
type StoragePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Storage *StorageFilterer) FilterPaused(opts *bind.FilterOpts) (*StoragePausedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StoragePausedIterator{contract: _Storage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Storage *StorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StoragePaused) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePaused)
				if err := _Storage.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Storage *StorageFilterer) ParsePaused(log types.Log) (*StoragePaused, error) {
	event := new(StoragePaused)
	if err := _Storage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Storage contract.
type StorageUnpausedIterator struct {
	Event *StorageUnpaused // Event containing the contract specifics and raw log

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
func (it *StorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUnpaused)
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
		it.Event = new(StorageUnpaused)
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
func (it *StorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUnpaused represents a Unpaused event raised by the Storage contract.
type StorageUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Storage *StorageFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StorageUnpausedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StorageUnpausedIterator{contract: _Storage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Storage *StorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StorageUnpaused) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUnpaused)
				if err := _Storage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Storage *StorageFilterer) ParseUnpaused(log types.Log) (*StorageUnpaused, error) {
	event := new(StorageUnpaused)
	if err := _Storage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
