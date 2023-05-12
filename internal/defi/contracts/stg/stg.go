// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stg

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

// StgMetaData contains all meta data concerning the Stg contract.
var StgMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_mainEndpointId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupplyOnMainEndpoint\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"qty\",\"type\":\"uint256\"}],\"name\":\"ReceiveFromChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"qty\",\"type\":\"uint256\"}],\"name\":\"SendToChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"dstContractLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"txParameters\",\"type\":\"bytes\"}],\"name\":\"estimateSendTokensFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_fromAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_pause\",\"type\":\"bool\"}],\"name\":\"pauseSendTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_qty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParam\",\"type\":\"bytes\"}],\"name\":\"sendTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_destinationContractAddress\",\"type\":\"bytes\"}],\"name\":\"setDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StgABI is the input ABI used to generate the binding from.
// Deprecated: Use StgMetaData.ABI instead.
var StgABI = StgMetaData.ABI

// Stg is an auto generated Go binding around an Ethereum contract.
type Stg struct {
	StgCaller     // Read-only binding to the contract
	StgTransactor // Write-only binding to the contract
	StgFilterer   // Log filterer for contract events
}

// StgCaller is an auto generated read-only Go binding around an Ethereum contract.
type StgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StgSession struct {
	Contract     *Stg              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StgCallerSession struct {
	Contract *StgCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StgTransactorSession struct {
	Contract     *StgTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StgRaw is an auto generated low-level Go binding around an Ethereum contract.
type StgRaw struct {
	Contract *Stg // Generic contract binding to access the raw methods on
}

// StgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StgCallerRaw struct {
	Contract *StgCaller // Generic read-only contract binding to access the raw methods on
}

// StgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StgTransactorRaw struct {
	Contract *StgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStg creates a new instance of Stg, bound to a specific deployed contract.
func NewStg(address common.Address, backend bind.ContractBackend) (*Stg, error) {
	contract, err := bindStg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stg{StgCaller: StgCaller{contract: contract}, StgTransactor: StgTransactor{contract: contract}, StgFilterer: StgFilterer{contract: contract}}, nil
}

// NewStgCaller creates a new read-only instance of Stg, bound to a specific deployed contract.
func NewStgCaller(address common.Address, caller bind.ContractCaller) (*StgCaller, error) {
	contract, err := bindStg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StgCaller{contract: contract}, nil
}

// NewStgTransactor creates a new write-only instance of Stg, bound to a specific deployed contract.
func NewStgTransactor(address common.Address, transactor bind.ContractTransactor) (*StgTransactor, error) {
	contract, err := bindStg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StgTransactor{contract: contract}, nil
}

// NewStgFilterer creates a new log filterer instance of Stg, bound to a specific deployed contract.
func NewStgFilterer(address common.Address, filterer bind.ContractFilterer) (*StgFilterer, error) {
	contract, err := bindStg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StgFilterer{contract: contract}, nil
}

// bindStg binds a generic wrapper to an already deployed contract.
func bindStg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stg *StgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stg.Contract.StgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stg *StgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stg.Contract.StgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stg *StgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stg.Contract.StgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stg *StgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stg *StgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stg *StgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stg.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Stg *StgCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Stg *StgSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Stg.Contract.Allowance(&_Stg.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Stg *StgCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Stg.Contract.Allowance(&_Stg.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stg *StgCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stg *StgSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stg.Contract.BalanceOf(&_Stg.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stg *StgCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stg.Contract.BalanceOf(&_Stg.CallOpts, account)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Stg *StgCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Stg *StgSession) ChainId() (uint16, error) {
	return _Stg.Contract.ChainId(&_Stg.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Stg *StgCallerSession) ChainId() (uint16, error) {
	return _Stg.Contract.ChainId(&_Stg.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stg *StgCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stg *StgSession) Decimals() (uint8, error) {
	return _Stg.Contract.Decimals(&_Stg.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stg *StgCallerSession) Decimals() (uint8, error) {
	return _Stg.Contract.Decimals(&_Stg.CallOpts)
}

// DstContractLookup is a free data retrieval call binding the contract method 0x60f05c7a.
//
// Solidity: function dstContractLookup(uint16 ) view returns(bytes)
func (_Stg *StgCaller) DstContractLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "dstContractLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DstContractLookup is a free data retrieval call binding the contract method 0x60f05c7a.
//
// Solidity: function dstContractLookup(uint16 ) view returns(bytes)
func (_Stg *StgSession) DstContractLookup(arg0 uint16) ([]byte, error) {
	return _Stg.Contract.DstContractLookup(&_Stg.CallOpts, arg0)
}

// DstContractLookup is a free data retrieval call binding the contract method 0x60f05c7a.
//
// Solidity: function dstContractLookup(uint16 ) view returns(bytes)
func (_Stg *StgCallerSession) DstContractLookup(arg0 uint16) ([]byte, error) {
	return _Stg.Contract.DstContractLookup(&_Stg.CallOpts, arg0)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Stg *StgCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Stg *StgSession) Endpoint() (common.Address, error) {
	return _Stg.Contract.Endpoint(&_Stg.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Stg *StgCallerSession) Endpoint() (common.Address, error) {
	return _Stg.Contract.Endpoint(&_Stg.CallOpts)
}

// EstimateSendTokensFee is a free data retrieval call binding the contract method 0x73874336.
//
// Solidity: function estimateSendTokensFee(uint16 _dstChainId, bool _useZro, bytes txParameters) view returns(uint256 nativeFee, uint256 zroFee)
func (_Stg *StgCaller) EstimateSendTokensFee(opts *bind.CallOpts, _dstChainId uint16, _useZro bool, txParameters []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "estimateSendTokensFee", _dstChainId, _useZro, txParameters)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendTokensFee is a free data retrieval call binding the contract method 0x73874336.
//
// Solidity: function estimateSendTokensFee(uint16 _dstChainId, bool _useZro, bytes txParameters) view returns(uint256 nativeFee, uint256 zroFee)
func (_Stg *StgSession) EstimateSendTokensFee(_dstChainId uint16, _useZro bool, txParameters []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Stg.Contract.EstimateSendTokensFee(&_Stg.CallOpts, _dstChainId, _useZro, txParameters)
}

// EstimateSendTokensFee is a free data retrieval call binding the contract method 0x73874336.
//
// Solidity: function estimateSendTokensFee(uint16 _dstChainId, bool _useZro, bytes txParameters) view returns(uint256 nativeFee, uint256 zroFee)
func (_Stg *StgCallerSession) EstimateSendTokensFee(_dstChainId uint16, _useZro bool, txParameters []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Stg.Contract.EstimateSendTokensFee(&_Stg.CallOpts, _dstChainId, _useZro, txParameters)
}

// IsMain is a free data retrieval call binding the contract method 0x604269d1.
//
// Solidity: function isMain() view returns(bool)
func (_Stg *StgCaller) IsMain(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "isMain")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMain is a free data retrieval call binding the contract method 0x604269d1.
//
// Solidity: function isMain() view returns(bool)
func (_Stg *StgSession) IsMain() (bool, error) {
	return _Stg.Contract.IsMain(&_Stg.CallOpts)
}

// IsMain is a free data retrieval call binding the contract method 0x604269d1.
//
// Solidity: function isMain() view returns(bool)
func (_Stg *StgCallerSession) IsMain() (bool, error) {
	return _Stg.Contract.IsMain(&_Stg.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stg *StgCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stg *StgSession) Name() (string, error) {
	return _Stg.Contract.Name(&_Stg.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stg *StgCallerSession) Name() (string, error) {
	return _Stg.Contract.Name(&_Stg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stg *StgCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stg *StgSession) Owner() (common.Address, error) {
	return _Stg.Contract.Owner(&_Stg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stg *StgCallerSession) Owner() (common.Address, error) {
	return _Stg.Contract.Owner(&_Stg.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stg *StgCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stg *StgSession) Paused() (bool, error) {
	return _Stg.Contract.Paused(&_Stg.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stg *StgCallerSession) Paused() (bool, error) {
	return _Stg.Contract.Paused(&_Stg.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stg *StgCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stg *StgSession) Symbol() (string, error) {
	return _Stg.Contract.Symbol(&_Stg.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stg *StgCallerSession) Symbol() (string, error) {
	return _Stg.Contract.Symbol(&_Stg.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stg *StgCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stg.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stg *StgSession) TotalSupply() (*big.Int, error) {
	return _Stg.Contract.TotalSupply(&_Stg.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stg *StgCallerSession) TotalSupply() (*big.Int, error) {
	return _Stg.Contract.TotalSupply(&_Stg.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stg *StgTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stg *StgSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.Approve(&_Stg.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stg *StgTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.Approve(&_Stg.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Stg *StgTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Stg *StgSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.DecreaseAllowance(&_Stg.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Stg *StgTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.DecreaseAllowance(&_Stg.TransactOpts, spender, subtractedValue)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Stg *StgTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Stg *StgSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Stg.Contract.ForceResumeReceive(&_Stg.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Stg *StgTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Stg.Contract.ForceResumeReceive(&_Stg.TransactOpts, _srcChainId, _srcAddress)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Stg *StgTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Stg *StgSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.IncreaseAllowance(&_Stg.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Stg *StgTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.IncreaseAllowance(&_Stg.TransactOpts, spender, addedValue)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _fromAddress, uint64 nonce, bytes _payload) returns()
func (_Stg *StgTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _fromAddress []byte, nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "lzReceive", _srcChainId, _fromAddress, nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _fromAddress, uint64 nonce, bytes _payload) returns()
func (_Stg *StgSession) LzReceive(_srcChainId uint16, _fromAddress []byte, nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Stg.Contract.LzReceive(&_Stg.TransactOpts, _srcChainId, _fromAddress, nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _fromAddress, uint64 nonce, bytes _payload) returns()
func (_Stg *StgTransactorSession) LzReceive(_srcChainId uint16, _fromAddress []byte, nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Stg.Contract.LzReceive(&_Stg.TransactOpts, _srcChainId, _fromAddress, nonce, _payload)
}

// PauseSendTokens is a paid mutator transaction binding the contract method 0xf1878922.
//
// Solidity: function pauseSendTokens(bool _pause) returns()
func (_Stg *StgTransactor) PauseSendTokens(opts *bind.TransactOpts, _pause bool) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "pauseSendTokens", _pause)
}

// PauseSendTokens is a paid mutator transaction binding the contract method 0xf1878922.
//
// Solidity: function pauseSendTokens(bool _pause) returns()
func (_Stg *StgSession) PauseSendTokens(_pause bool) (*types.Transaction, error) {
	return _Stg.Contract.PauseSendTokens(&_Stg.TransactOpts, _pause)
}

// PauseSendTokens is a paid mutator transaction binding the contract method 0xf1878922.
//
// Solidity: function pauseSendTokens(bool _pause) returns()
func (_Stg *StgTransactorSession) PauseSendTokens(_pause bool) (*types.Transaction, error) {
	return _Stg.Contract.PauseSendTokens(&_Stg.TransactOpts, _pause)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stg *StgTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stg *StgSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stg.Contract.RenounceOwnership(&_Stg.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stg *StgTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stg.Contract.RenounceOwnership(&_Stg.TransactOpts)
}

// SendTokens is a paid mutator transaction binding the contract method 0x2e15238c.
//
// Solidity: function sendTokens(uint16 _dstChainId, bytes _to, uint256 _qty, address zroPaymentAddress, bytes adapterParam) payable returns()
func (_Stg *StgTransactor) SendTokens(opts *bind.TransactOpts, _dstChainId uint16, _to []byte, _qty *big.Int, zroPaymentAddress common.Address, adapterParam []byte) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "sendTokens", _dstChainId, _to, _qty, zroPaymentAddress, adapterParam)
}

// SendTokens is a paid mutator transaction binding the contract method 0x2e15238c.
//
// Solidity: function sendTokens(uint16 _dstChainId, bytes _to, uint256 _qty, address zroPaymentAddress, bytes adapterParam) payable returns()
func (_Stg *StgSession) SendTokens(_dstChainId uint16, _to []byte, _qty *big.Int, zroPaymentAddress common.Address, adapterParam []byte) (*types.Transaction, error) {
	return _Stg.Contract.SendTokens(&_Stg.TransactOpts, _dstChainId, _to, _qty, zroPaymentAddress, adapterParam)
}

// SendTokens is a paid mutator transaction binding the contract method 0x2e15238c.
//
// Solidity: function sendTokens(uint16 _dstChainId, bytes _to, uint256 _qty, address zroPaymentAddress, bytes adapterParam) payable returns()
func (_Stg *StgTransactorSession) SendTokens(_dstChainId uint16, _to []byte, _qty *big.Int, zroPaymentAddress common.Address, adapterParam []byte) (*types.Transaction, error) {
	return _Stg.Contract.SendTokens(&_Stg.TransactOpts, _dstChainId, _to, _qty, zroPaymentAddress, adapterParam)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Stg *StgTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Stg *StgSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Stg.Contract.SetConfig(&_Stg.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Stg *StgTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Stg.Contract.SetConfig(&_Stg.TransactOpts, _version, _chainId, _configType, _config)
}

// SetDestination is a paid mutator transaction binding the contract method 0x706d8fff.
//
// Solidity: function setDestination(uint16 _dstChainId, bytes _destinationContractAddress) returns()
func (_Stg *StgTransactor) SetDestination(opts *bind.TransactOpts, _dstChainId uint16, _destinationContractAddress []byte) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "setDestination", _dstChainId, _destinationContractAddress)
}

// SetDestination is a paid mutator transaction binding the contract method 0x706d8fff.
//
// Solidity: function setDestination(uint16 _dstChainId, bytes _destinationContractAddress) returns()
func (_Stg *StgSession) SetDestination(_dstChainId uint16, _destinationContractAddress []byte) (*types.Transaction, error) {
	return _Stg.Contract.SetDestination(&_Stg.TransactOpts, _dstChainId, _destinationContractAddress)
}

// SetDestination is a paid mutator transaction binding the contract method 0x706d8fff.
//
// Solidity: function setDestination(uint16 _dstChainId, bytes _destinationContractAddress) returns()
func (_Stg *StgTransactorSession) SetDestination(_dstChainId uint16, _destinationContractAddress []byte) (*types.Transaction, error) {
	return _Stg.Contract.SetDestination(&_Stg.TransactOpts, _dstChainId, _destinationContractAddress)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 version) returns()
func (_Stg *StgTransactor) SetReceiveVersion(opts *bind.TransactOpts, version uint16) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "setReceiveVersion", version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 version) returns()
func (_Stg *StgSession) SetReceiveVersion(version uint16) (*types.Transaction, error) {
	return _Stg.Contract.SetReceiveVersion(&_Stg.TransactOpts, version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 version) returns()
func (_Stg *StgTransactorSession) SetReceiveVersion(version uint16) (*types.Transaction, error) {
	return _Stg.Contract.SetReceiveVersion(&_Stg.TransactOpts, version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 version) returns()
func (_Stg *StgTransactor) SetSendVersion(opts *bind.TransactOpts, version uint16) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "setSendVersion", version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 version) returns()
func (_Stg *StgSession) SetSendVersion(version uint16) (*types.Transaction, error) {
	return _Stg.Contract.SetSendVersion(&_Stg.TransactOpts, version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 version) returns()
func (_Stg *StgTransactorSession) SetSendVersion(version uint16) (*types.Transaction, error) {
	return _Stg.Contract.SetSendVersion(&_Stg.TransactOpts, version)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Stg *StgTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Stg *StgSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.Transfer(&_Stg.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Stg *StgTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.Transfer(&_Stg.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Stg *StgTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Stg *StgSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.TransferFrom(&_Stg.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Stg *StgTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stg.Contract.TransferFrom(&_Stg.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stg *StgTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stg.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stg *StgSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stg.Contract.TransferOwnership(&_Stg.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stg *StgTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stg.Contract.TransferOwnership(&_Stg.TransactOpts, newOwner)
}

// StgApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Stg contract.
type StgApprovalIterator struct {
	Event *StgApproval // Event containing the contract specifics and raw log

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
func (it *StgApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgApproval)
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
		it.Event = new(StgApproval)
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
func (it *StgApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgApproval represents a Approval event raised by the Stg contract.
type StgApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stg *StgFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StgApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Stg.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StgApprovalIterator{contract: _Stg.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stg *StgFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StgApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Stg.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgApproval)
				if err := _Stg.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stg *StgFilterer) ParseApproval(log types.Log) (*StgApproval, error) {
	event := new(StgApproval)
	if err := _Stg.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StgOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stg contract.
type StgOwnershipTransferredIterator struct {
	Event *StgOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StgOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgOwnershipTransferred)
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
		it.Event = new(StgOwnershipTransferred)
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
func (it *StgOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgOwnershipTransferred represents a OwnershipTransferred event raised by the Stg contract.
type StgOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stg *StgFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StgOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stg.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StgOwnershipTransferredIterator{contract: _Stg.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stg *StgFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StgOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stg.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgOwnershipTransferred)
				if err := _Stg.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stg *StgFilterer) ParseOwnershipTransferred(log types.Log) (*StgOwnershipTransferred, error) {
	event := new(StgOwnershipTransferred)
	if err := _Stg.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StgPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Stg contract.
type StgPausedIterator struct {
	Event *StgPaused // Event containing the contract specifics and raw log

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
func (it *StgPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgPaused)
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
		it.Event = new(StgPaused)
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
func (it *StgPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgPaused represents a Paused event raised by the Stg contract.
type StgPaused struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool isPaused)
func (_Stg *StgFilterer) FilterPaused(opts *bind.FilterOpts) (*StgPausedIterator, error) {

	logs, sub, err := _Stg.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StgPausedIterator{contract: _Stg.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool isPaused)
func (_Stg *StgFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StgPaused) (event.Subscription, error) {

	logs, sub, err := _Stg.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgPaused)
				if err := _Stg.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool isPaused)
func (_Stg *StgFilterer) ParsePaused(log types.Log) (*StgPaused, error) {
	event := new(StgPaused)
	if err := _Stg.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StgReceiveFromChainIterator is returned from FilterReceiveFromChain and is used to iterate over the raw logs and unpacked data for ReceiveFromChain events raised by the Stg contract.
type StgReceiveFromChainIterator struct {
	Event *StgReceiveFromChain // Event containing the contract specifics and raw log

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
func (it *StgReceiveFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgReceiveFromChain)
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
		it.Event = new(StgReceiveFromChain)
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
func (it *StgReceiveFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgReceiveFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgReceiveFromChain represents a ReceiveFromChain event raised by the Stg contract.
type StgReceiveFromChain struct {
	SrcChainId uint16
	Nonce      uint64
	Qty        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveFromChain is a free log retrieval operation binding the contract event 0x831bc68226f8d1f734ffcca73602efc4eca13711402ba1d2cc05ee17bb54f631.
//
// Solidity: event ReceiveFromChain(uint16 srcChainId, uint64 nonce, uint256 qty)
func (_Stg *StgFilterer) FilterReceiveFromChain(opts *bind.FilterOpts) (*StgReceiveFromChainIterator, error) {

	logs, sub, err := _Stg.contract.FilterLogs(opts, "ReceiveFromChain")
	if err != nil {
		return nil, err
	}
	return &StgReceiveFromChainIterator{contract: _Stg.contract, event: "ReceiveFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveFromChain is a free log subscription operation binding the contract event 0x831bc68226f8d1f734ffcca73602efc4eca13711402ba1d2cc05ee17bb54f631.
//
// Solidity: event ReceiveFromChain(uint16 srcChainId, uint64 nonce, uint256 qty)
func (_Stg *StgFilterer) WatchReceiveFromChain(opts *bind.WatchOpts, sink chan<- *StgReceiveFromChain) (event.Subscription, error) {

	logs, sub, err := _Stg.contract.WatchLogs(opts, "ReceiveFromChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgReceiveFromChain)
				if err := _Stg.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
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

// ParseReceiveFromChain is a log parse operation binding the contract event 0x831bc68226f8d1f734ffcca73602efc4eca13711402ba1d2cc05ee17bb54f631.
//
// Solidity: event ReceiveFromChain(uint16 srcChainId, uint64 nonce, uint256 qty)
func (_Stg *StgFilterer) ParseReceiveFromChain(log types.Log) (*StgReceiveFromChain, error) {
	event := new(StgReceiveFromChain)
	if err := _Stg.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StgSendToChainIterator is returned from FilterSendToChain and is used to iterate over the raw logs and unpacked data for SendToChain events raised by the Stg contract.
type StgSendToChainIterator struct {
	Event *StgSendToChain // Event containing the contract specifics and raw log

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
func (it *StgSendToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgSendToChain)
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
		it.Event = new(StgSendToChain)
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
func (it *StgSendToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgSendToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgSendToChain represents a SendToChain event raised by the Stg contract.
type StgSendToChain struct {
	DstChainId uint16
	To         []byte
	Qty        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendToChain is a free log retrieval operation binding the contract event 0x664e26797cde1146ddfcb9a5d3f4de61179f9c11b2698599bb09e686f442172b.
//
// Solidity: event SendToChain(uint16 dstChainId, bytes to, uint256 qty)
func (_Stg *StgFilterer) FilterSendToChain(opts *bind.FilterOpts) (*StgSendToChainIterator, error) {

	logs, sub, err := _Stg.contract.FilterLogs(opts, "SendToChain")
	if err != nil {
		return nil, err
	}
	return &StgSendToChainIterator{contract: _Stg.contract, event: "SendToChain", logs: logs, sub: sub}, nil
}

// WatchSendToChain is a free log subscription operation binding the contract event 0x664e26797cde1146ddfcb9a5d3f4de61179f9c11b2698599bb09e686f442172b.
//
// Solidity: event SendToChain(uint16 dstChainId, bytes to, uint256 qty)
func (_Stg *StgFilterer) WatchSendToChain(opts *bind.WatchOpts, sink chan<- *StgSendToChain) (event.Subscription, error) {

	logs, sub, err := _Stg.contract.WatchLogs(opts, "SendToChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgSendToChain)
				if err := _Stg.contract.UnpackLog(event, "SendToChain", log); err != nil {
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

// ParseSendToChain is a log parse operation binding the contract event 0x664e26797cde1146ddfcb9a5d3f4de61179f9c11b2698599bb09e686f442172b.
//
// Solidity: event SendToChain(uint16 dstChainId, bytes to, uint256 qty)
func (_Stg *StgFilterer) ParseSendToChain(log types.Log) (*StgSendToChain, error) {
	event := new(StgSendToChain)
	if err := _Stg.contract.UnpackLog(event, "SendToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StgTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Stg contract.
type StgTransferIterator struct {
	Event *StgTransfer // Event containing the contract specifics and raw log

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
func (it *StgTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StgTransfer)
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
		it.Event = new(StgTransfer)
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
func (it *StgTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StgTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StgTransfer represents a Transfer event raised by the Stg contract.
type StgTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Stg *StgFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StgTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stg.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StgTransferIterator{contract: _Stg.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Stg *StgFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StgTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Stg.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StgTransfer)
				if err := _Stg.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Stg *StgFilterer) ParseTransfer(log types.Log) (*StgTransfer, error) {
	event := new(StgTransfer)
	if err := _Stg.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
