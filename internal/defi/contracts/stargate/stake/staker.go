// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// StakerMetaData contains all meta data concerning the Staker contract.
var StakerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"min_time\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVotingEscrow.DepositType\",\"name\":\"deposit_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAXTIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"add_to_whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_t\",\"type\":\"uint256\"}],\"name\":\"balanceOfAtT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newController\",\"type\":\"address\"}],\"name\":\"changeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contracts_whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"name\":\"create_lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit_for\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"get_last_user_slope\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"increase_amount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"name\":\"increase_amount_and_time\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"name\":\"increase_unlock_time\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"amount\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"locked__end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"point_history\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"bias\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"slope\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blk\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"remove_from_whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slope_changes\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAtT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfersEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"user_point_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_point_history\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"bias\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"slope\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blk\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"user_point_history__ts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"name\":\"withdraw_and_create_lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerMetaData.ABI instead.
var StakerABI = StakerMetaData.ABI

// Staker is an auto generated Go binding around an Ethereum contract.
type Staker struct {
	StakerCaller     // Read-only binding to the contract
	StakerTransactor // Write-only binding to the contract
	StakerFilterer   // Log filterer for contract events
}

// StakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerSession struct {
	Contract     *Staker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerCallerSession struct {
	Contract *StakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerTransactorSession struct {
	Contract     *StakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakerRaw struct {
	Contract *Staker // Generic contract binding to access the raw methods on
}

// StakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerCallerRaw struct {
	Contract *StakerCaller // Generic read-only contract binding to access the raw methods on
}

// StakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakerTransactorRaw struct {
	Contract *StakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaker creates a new instance of Staker, bound to a specific deployed contract.
func NewStaker(address common.Address, backend bind.ContractBackend) (*Staker, error) {
	contract, err := bindStaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staker{StakerCaller: StakerCaller{contract: contract}, StakerTransactor: StakerTransactor{contract: contract}, StakerFilterer: StakerFilterer{contract: contract}}, nil
}

// NewStakerCaller creates a new read-only instance of Staker, bound to a specific deployed contract.
func NewStakerCaller(address common.Address, caller bind.ContractCaller) (*StakerCaller, error) {
	contract, err := bindStaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerCaller{contract: contract}, nil
}

// NewStakerTransactor creates a new write-only instance of Staker, bound to a specific deployed contract.
func NewStakerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakerTransactor, error) {
	contract, err := bindStaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerTransactor{contract: contract}, nil
}

// NewStakerFilterer creates a new log filterer instance of Staker, bound to a specific deployed contract.
func NewStakerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakerFilterer, error) {
	contract, err := bindStaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerFilterer{contract: contract}, nil
}

// bindStaker binds a generic wrapper to an already deployed contract.
func bindStaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staker *StakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staker.Contract.StakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staker *StakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.Contract.StakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staker *StakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staker.Contract.StakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staker *StakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staker *StakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staker *StakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staker.Contract.contract.Transact(opts, method, params...)
}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Staker *StakerCaller) MAXTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "MAXTIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Staker *StakerSession) MAXTIME() (*big.Int, error) {
	return _Staker.Contract.MAXTIME(&_Staker.CallOpts)
}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Staker *StakerCallerSession) MAXTIME() (*big.Int, error) {
	return _Staker.Contract.MAXTIME(&_Staker.CallOpts)
}

// MINTIME is a free data retrieval call binding the contract method 0x5b51c308.
//
// Solidity: function MINTIME() view returns(uint256)
func (_Staker *StakerCaller) MINTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "MINTIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTIME is a free data retrieval call binding the contract method 0x5b51c308.
//
// Solidity: function MINTIME() view returns(uint256)
func (_Staker *StakerSession) MINTIME() (*big.Int, error) {
	return _Staker.Contract.MINTIME(&_Staker.CallOpts)
}

// MINTIME is a free data retrieval call binding the contract method 0x5b51c308.
//
// Solidity: function MINTIME() view returns(uint256)
func (_Staker *StakerCallerSession) MINTIME() (*big.Int, error) {
	return _Staker.Contract.MINTIME(&_Staker.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Staker *StakerCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.BalanceOf(&_Staker.CallOpts, addr)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Staker *StakerCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Staker *StakerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Staker.Contract.BalanceOfAt(&_Staker.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Staker *StakerCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Staker.Contract.BalanceOfAt(&_Staker.CallOpts, addr, _block)
}

// BalanceOfAtT is a free data retrieval call binding the contract method 0xd07b705f.
//
// Solidity: function balanceOfAtT(address addr, uint256 _t) view returns(uint256)
func (_Staker *StakerCaller) BalanceOfAtT(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "balanceOfAtT", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtT is a free data retrieval call binding the contract method 0xd07b705f.
//
// Solidity: function balanceOfAtT(address addr, uint256 _t) view returns(uint256)
func (_Staker *StakerSession) BalanceOfAtT(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Staker.Contract.BalanceOfAtT(&_Staker.CallOpts, addr, _t)
}

// BalanceOfAtT is a free data retrieval call binding the contract method 0xd07b705f.
//
// Solidity: function balanceOfAtT(address addr, uint256 _t) view returns(uint256)
func (_Staker *StakerCallerSession) BalanceOfAtT(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Staker.Contract.BalanceOfAtT(&_Staker.CallOpts, addr, _t)
}

// ContractsWhitelist is a free data retrieval call binding the contract method 0x3617a204.
//
// Solidity: function contracts_whitelist(address ) view returns(bool)
func (_Staker *StakerCaller) ContractsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "contracts_whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractsWhitelist is a free data retrieval call binding the contract method 0x3617a204.
//
// Solidity: function contracts_whitelist(address ) view returns(bool)
func (_Staker *StakerSession) ContractsWhitelist(arg0 common.Address) (bool, error) {
	return _Staker.Contract.ContractsWhitelist(&_Staker.CallOpts, arg0)
}

// ContractsWhitelist is a free data retrieval call binding the contract method 0x3617a204.
//
// Solidity: function contracts_whitelist(address ) view returns(bool)
func (_Staker *StakerCallerSession) ContractsWhitelist(arg0 common.Address) (bool, error) {
	return _Staker.Contract.ContractsWhitelist(&_Staker.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staker *StakerCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staker *StakerSession) Controller() (common.Address, error) {
	return _Staker.Contract.Controller(&_Staker.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Staker *StakerCallerSession) Controller() (common.Address, error) {
	return _Staker.Contract.Controller(&_Staker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staker *StakerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staker *StakerSession) Decimals() (uint8, error) {
	return _Staker.Contract.Decimals(&_Staker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staker *StakerCallerSession) Decimals() (uint8, error) {
	return _Staker.Contract.Decimals(&_Staker.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerSession) Epoch() (*big.Int, error) {
	return _Staker.Contract.Epoch(&_Staker.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Staker *StakerCallerSession) Epoch() (*big.Int, error) {
	return _Staker.Contract.Epoch(&_Staker.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Staker *StakerCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Staker *StakerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.GetLastUserSlope(&_Staker.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Staker *StakerCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Staker.Contract.GetLastUserSlope(&_Staker.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(int128 amount, uint256 end)
func (_Staker *StakerCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "locked", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(int128 amount, uint256 end)
func (_Staker *StakerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Staker.Contract.Locked(&_Staker.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(int128 amount, uint256 end)
func (_Staker *StakerCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Staker.Contract.Locked(&_Staker.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Staker *StakerCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Staker *StakerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Staker.Contract.LockedEnd(&_Staker.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Staker *StakerCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Staker.Contract.LockedEnd(&_Staker.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staker *StakerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staker *StakerSession) Name() (string, error) {
	return _Staker.Contract.Name(&_Staker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staker *StakerCallerSession) Name() (string, error) {
	return _Staker.Contract.Name(&_Staker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staker *StakerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staker *StakerSession) Owner() (common.Address, error) {
	return _Staker.Contract.Owner(&_Staker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staker *StakerCallerSession) Owner() (common.Address, error) {
	return _Staker.Contract.Owner(&_Staker.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Staker.Contract.PointHistory(&_Staker.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Staker.Contract.PointHistory(&_Staker.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Staker *StakerCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Staker *StakerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Staker.Contract.SlopeChanges(&_Staker.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 ) view returns(int128)
func (_Staker *StakerCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Staker.Contract.SlopeChanges(&_Staker.CallOpts, arg0)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Staker *StakerCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Staker *StakerSession) Supply() (*big.Int, error) {
	return _Staker.Contract.Supply(&_Staker.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Staker *StakerCallerSession) Supply() (*big.Int, error) {
	return _Staker.Contract.Supply(&_Staker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staker *StakerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staker *StakerSession) Symbol() (string, error) {
	return _Staker.Contract.Symbol(&_Staker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staker *StakerCallerSession) Symbol() (string, error) {
	return _Staker.Contract.Symbol(&_Staker.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Staker *StakerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Staker *StakerSession) Token() (common.Address, error) {
	return _Staker.Contract.Token(&_Staker.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Staker *StakerCallerSession) Token() (common.Address, error) {
	return _Staker.Contract.Token(&_Staker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staker *StakerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staker *StakerSession) TotalSupply() (*big.Int, error) {
	return _Staker.Contract.TotalSupply(&_Staker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staker *StakerCallerSession) TotalSupply() (*big.Int, error) {
	return _Staker.Contract.TotalSupply(&_Staker.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Staker *StakerCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Staker *StakerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Staker.Contract.TotalSupplyAt(&_Staker.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Staker *StakerCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Staker.Contract.TotalSupplyAt(&_Staker.CallOpts, _block)
}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Staker *StakerCaller) TotalSupplyAtT(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "totalSupplyAtT", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Staker *StakerSession) TotalSupplyAtT(t *big.Int) (*big.Int, error) {
	return _Staker.Contract.TotalSupplyAtT(&_Staker.CallOpts, t)
}

// TotalSupplyAtT is a free data retrieval call binding the contract method 0x7116c60c.
//
// Solidity: function totalSupplyAtT(uint256 t) view returns(uint256)
func (_Staker *StakerCallerSession) TotalSupplyAtT(t *big.Int) (*big.Int, error) {
	return _Staker.Contract.TotalSupplyAtT(&_Staker.CallOpts, t)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Staker *StakerCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "transfersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Staker *StakerSession) TransfersEnabled() (bool, error) {
	return _Staker.Contract.TransfersEnabled(&_Staker.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Staker *StakerCallerSession) TransfersEnabled() (bool, error) {
	return _Staker.Contract.TransfersEnabled(&_Staker.CallOpts)
}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Staker *StakerCaller) Unlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "unlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Staker *StakerSession) Unlocked() (bool, error) {
	return _Staker.Contract.Unlocked(&_Staker.CallOpts)
}

// Unlocked is a free data retrieval call binding the contract method 0x6a5e2650.
//
// Solidity: function unlocked() view returns(bool)
func (_Staker *StakerCallerSession) Unlocked() (bool, error) {
	return _Staker.Contract.Unlocked(&_Staker.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address ) view returns(uint256)
func (_Staker *StakerCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address ) view returns(uint256)
func (_Staker *StakerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Staker.Contract.UserPointEpoch(&_Staker.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address ) view returns(uint256)
func (_Staker *StakerCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Staker.Contract.UserPointEpoch(&_Staker.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Staker.Contract.UserPointHistory(&_Staker.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address , uint256 ) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Staker *StakerCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Staker.Contract.UserPointHistory(&_Staker.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Staker *StakerCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Staker *StakerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Staker.Contract.UserPointHistoryTs(&_Staker.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Staker *StakerCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Staker.Contract.UserPointHistoryTs(&_Staker.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Staker *StakerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staker.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Staker *StakerSession) Version() (string, error) {
	return _Staker.Contract.Version(&_Staker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Staker *StakerCallerSession) Version() (string, error) {
	return _Staker.Contract.Version(&_Staker.CallOpts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xac25f266.
//
// Solidity: function add_to_whitelist(address addr) returns()
func (_Staker *StakerTransactor) AddToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "add_to_whitelist", addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xac25f266.
//
// Solidity: function add_to_whitelist(address addr) returns()
func (_Staker *StakerSession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Staker.Contract.AddToWhitelist(&_Staker.TransactOpts, addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xac25f266.
//
// Solidity: function add_to_whitelist(address addr) returns()
func (_Staker *StakerTransactorSession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Staker.Contract.AddToWhitelist(&_Staker.TransactOpts, addr)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Staker *StakerTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Staker *StakerSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Staker.Contract.ChangeController(&_Staker.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Staker *StakerTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Staker.Contract.ChangeController(&_Staker.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Staker *StakerTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Staker *StakerSession) Checkpoint() (*types.Transaction, error) {
	return _Staker.Contract.Checkpoint(&_Staker.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Staker *StakerTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Staker.Contract.Checkpoint(&_Staker.TransactOpts)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.CreateLock(&_Staker.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.CreateLock(&_Staker.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Staker *StakerTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Staker *StakerSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.DepositFor(&_Staker.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Staker *StakerTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.DepositFor(&_Staker.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Staker *StakerTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Staker *StakerSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseAmount(&_Staker.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Staker *StakerTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseAmount(&_Staker.TransactOpts, _value)
}

// IncreaseAmountAndTime is a paid mutator transaction binding the contract method 0x7142a6a6.
//
// Solidity: function increase_amount_and_time(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactor) IncreaseAmountAndTime(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "increase_amount_and_time", _value, _unlock_time)
}

// IncreaseAmountAndTime is a paid mutator transaction binding the contract method 0x7142a6a6.
//
// Solidity: function increase_amount_and_time(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerSession) IncreaseAmountAndTime(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseAmountAndTime(&_Staker.TransactOpts, _value, _unlock_time)
}

// IncreaseAmountAndTime is a paid mutator transaction binding the contract method 0x7142a6a6.
//
// Solidity: function increase_amount_and_time(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactorSession) IncreaseAmountAndTime(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseAmountAndTime(&_Staker.TransactOpts, _value, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Staker *StakerTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Staker *StakerSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseUnlockTime(&_Staker.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Staker *StakerTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.IncreaseUnlockTime(&_Staker.TransactOpts, _unlock_time)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x90fad1e6.
//
// Solidity: function remove_from_whitelist(address addr) returns()
func (_Staker *StakerTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "remove_from_whitelist", addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x90fad1e6.
//
// Solidity: function remove_from_whitelist(address addr) returns()
func (_Staker *StakerSession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Staker.Contract.RemoveFromWhitelist(&_Staker.TransactOpts, addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x90fad1e6.
//
// Solidity: function remove_from_whitelist(address addr) returns()
func (_Staker *StakerTransactorSession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Staker.Contract.RemoveFromWhitelist(&_Staker.TransactOpts, addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staker *StakerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staker *StakerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staker.Contract.RenounceOwnership(&_Staker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staker *StakerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staker.Contract.RenounceOwnership(&_Staker.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staker *StakerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staker *StakerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staker.Contract.TransferOwnership(&_Staker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staker *StakerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staker.Contract.TransferOwnership(&_Staker.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Staker *StakerTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Staker *StakerSession) Unlock() (*types.Transaction, error) {
	return _Staker.Contract.Unlock(&_Staker.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Staker *StakerTransactorSession) Unlock() (*types.Transaction, error) {
	return _Staker.Contract.Unlock(&_Staker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerSession) Withdraw() (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *StakerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Staker.Contract.Withdraw(&_Staker.TransactOpts)
}

// WithdrawAndCreateLock is a paid mutator transaction binding the contract method 0x2371eb23.
//
// Solidity: function withdraw_and_create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactor) WithdrawAndCreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.contract.Transact(opts, "withdraw_and_create_lock", _value, _unlock_time)
}

// WithdrawAndCreateLock is a paid mutator transaction binding the contract method 0x2371eb23.
//
// Solidity: function withdraw_and_create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerSession) WithdrawAndCreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.WithdrawAndCreateLock(&_Staker.TransactOpts, _value, _unlock_time)
}

// WithdrawAndCreateLock is a paid mutator transaction binding the contract method 0x2371eb23.
//
// Solidity: function withdraw_and_create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Staker *StakerTransactorSession) WithdrawAndCreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Staker.Contract.WithdrawAndCreateLock(&_Staker.TransactOpts, _value, _unlock_time)
}

// StakerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Staker contract.
type StakerDepositIterator struct {
	Event *StakerDeposit // Event containing the contract specifics and raw log

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
func (it *StakerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerDeposit)
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
		it.Event = new(StakerDeposit)
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
func (it *StakerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerDeposit represents a Deposit event raised by the Staker contract.
type StakerDeposit struct {
	Provider    common.Address
	Value       *big.Int
	Locktime    *big.Int
	DepositType uint8
	Ts          *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xbe9cf0e939c614fad640a623a53ba0a807c8cb503c4c4c8dacabe27b86ff2dd5.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Staker *StakerFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*StakerDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &StakerDepositIterator{contract: _Staker.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xbe9cf0e939c614fad640a623a53ba0a807c8cb503c4c4c8dacabe27b86ff2dd5.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Staker *StakerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakerDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerDeposit)
				if err := _Staker.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xbe9cf0e939c614fad640a623a53ba0a807c8cb503c4c4c8dacabe27b86ff2dd5.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, uint8 deposit_type, uint256 ts)
func (_Staker *StakerFilterer) ParseDeposit(log types.Log) (*StakerDeposit, error) {
	event := new(StakerDeposit)
	if err := _Staker.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staker contract.
type StakerOwnershipTransferredIterator struct {
	Event *StakerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerOwnershipTransferred)
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
		it.Event = new(StakerOwnershipTransferred)
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
func (it *StakerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerOwnershipTransferred represents a OwnershipTransferred event raised by the Staker contract.
type StakerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staker *StakerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakerOwnershipTransferredIterator{contract: _Staker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staker *StakerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerOwnershipTransferred)
				if err := _Staker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staker *StakerFilterer) ParseOwnershipTransferred(log types.Log) (*StakerOwnershipTransferred, error) {
	event := new(StakerOwnershipTransferred)
	if err := _Staker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Staker contract.
type StakerSupplyIterator struct {
	Event *StakerSupply // Event containing the contract specifics and raw log

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
func (it *StakerSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerSupply)
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
		it.Event = new(StakerSupply)
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
func (it *StakerSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerSupply represents a Supply event raised by the Staker contract.
type StakerSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Staker *StakerFilterer) FilterSupply(opts *bind.FilterOpts) (*StakerSupplyIterator, error) {

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &StakerSupplyIterator{contract: _Staker.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Staker *StakerFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *StakerSupply) (event.Subscription, error) {

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerSupply)
				if err := _Staker.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Staker *StakerFilterer) ParseSupply(log types.Log) (*StakerSupply, error) {
	event := new(StakerSupply)
	if err := _Staker.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Staker contract.
type StakerWithdrawIterator struct {
	Event *StakerWithdraw // Event containing the contract specifics and raw log

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
func (it *StakerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerWithdraw)
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
		it.Event = new(StakerWithdraw)
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
func (it *StakerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerWithdraw represents a Withdraw event raised by the Staker contract.
type StakerWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Staker *StakerFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*StakerWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Staker.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &StakerWithdrawIterator{contract: _Staker.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Staker *StakerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakerWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Staker.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerWithdraw)
				if err := _Staker.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Staker *StakerFilterer) ParseWithdraw(log types.Log) (*StakerWithdraw, error) {
	event := new(StakerWithdraw)
	if err := _Staker.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
