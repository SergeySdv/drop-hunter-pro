// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package startgaterouter

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

// ICommonOFTLzCallParams is an auto generated low-level Go binding around an user-defined struct.
type ICommonOFTLzCallParams struct {
	RefundAddress     common.Address
	ZroPaymentAddress common.Address
	AdapterParams     []byte
}

// IOFTWrapperFeeObj is an auto generated low-level Go binding around an user-defined struct.
type IOFTWrapperFeeObj struct {
	CallerBps *big.Int
	Caller    common.Address
	PartnerId [2]byte
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultBps\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WrapperFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrapperFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callerFee\",\"type\":\"uint256\"}],\"name\":\"WrapperFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"estimateSendFeeV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callerBps\",\"type\":\"uint256\"}],\"name\":\"getAmountAndFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrapperFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oftBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFTFeeV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendOFTV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFTFeeV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyOft\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_toAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"internalType\":\"structICommonOFT.LzCallParams\",\"name\":\"_callParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"callerBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"partnerId\",\"type\":\"bytes2\"}],\"internalType\":\"structIOFTWrapper.FeeObj\",\"name\":\"_feeObj\",\"type\":\"tuple\"}],\"name\":\"sendProxyOFTV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setOFTBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Storage *StorageCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Storage *StorageSession) BPSDENOMINATOR() (*big.Int, error) {
	return _Storage.Contract.BPSDENOMINATOR(&_Storage.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_Storage *StorageCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _Storage.Contract.BPSDENOMINATOR(&_Storage.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Storage *StorageCaller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Storage *StorageSession) MAXUINT() (*big.Int, error) {
	return _Storage.Contract.MAXUINT(&_Storage.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Storage *StorageCallerSession) MAXUINT() (*big.Int, error) {
	return _Storage.Contract.MAXUINT(&_Storage.CallOpts)
}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Storage *StorageCaller) DefaultBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "defaultBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Storage *StorageSession) DefaultBps() (*big.Int, error) {
	return _Storage.Contract.DefaultBps(&_Storage.CallOpts)
}

// DefaultBps is a free data retrieval call binding the contract method 0xd1b308dc.
//
// Solidity: function defaultBps() view returns(uint256)
func (_Storage *StorageCallerSession) DefaultBps() (*big.Int, error) {
	return _Storage.Contract.DefaultBps(&_Storage.CallOpts)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageCaller) EstimateSendFee(opts *bind.CallOpts, _oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "estimateSendFee", _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)

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

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageSession) EstimateSendFee(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Storage.Contract.EstimateSendFee(&_Storage.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0xe1bafc80.
//
// Solidity: function estimateSendFee(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageCallerSession) EstimateSendFee(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Storage.Contract.EstimateSendFee(&_Storage.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageCaller) EstimateSendFeeV2(opts *bind.CallOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "estimateSendFeeV2", _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)

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

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageSession) EstimateSendFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Storage.Contract.EstimateSendFeeV2(&_Storage.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// EstimateSendFeeV2 is a free data retrieval call binding the contract method 0x8d8c915c.
//
// Solidity: function estimateSendFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, bool _useZro, bytes _adapterParams, (uint256,address,bytes2) _feeObj) view returns(uint256 nativeFee, uint256 zroFee)
func (_Storage *StorageCallerSession) EstimateSendFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _useZro bool, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Storage.Contract.EstimateSendFeeV2(&_Storage.CallOpts, _oft, _dstChainId, _toAddress, _amount, _useZro, _adapterParams, _feeObj)
}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageCaller) GetAmountAndFees(opts *bind.CallOpts, _token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAmountAndFees", _token, _amount, _callerBps)

	outstruct := new(struct {
		Amount     *big.Int
		WrapperFee *big.Int
		CallerFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WrapperFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CallerFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageSession) GetAmountAndFees(_token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	return _Storage.Contract.GetAmountAndFees(&_Storage.CallOpts, _token, _amount, _callerBps)
}

// GetAmountAndFees is a free data retrieval call binding the contract method 0x17696f64.
//
// Solidity: function getAmountAndFees(address _token, uint256 _amount, uint256 _callerBps) view returns(uint256 amount, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageCallerSession) GetAmountAndFees(_token common.Address, _amount *big.Int, _callerBps *big.Int) (struct {
	Amount     *big.Int
	WrapperFee *big.Int
	CallerFee  *big.Int
}, error) {
	return _Storage.Contract.GetAmountAndFees(&_Storage.CallOpts, _token, _amount, _callerBps)
}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Storage *StorageCaller) OftBps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "oftBps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Storage *StorageSession) OftBps(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.OftBps(&_Storage.CallOpts, arg0)
}

// OftBps is a free data retrieval call binding the contract method 0x7a751182.
//
// Solidity: function oftBps(address ) view returns(uint256)
func (_Storage *StorageCallerSession) OftBps(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.OftBps(&_Storage.CallOpts, arg0)
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

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendOFT(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendOFT", _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendOFT(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFT(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFT is a paid mutator transaction binding the contract method 0x498eff64.
//
// Solidity: function sendOFT(address _oft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendOFT(_oft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFT(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendOFTFeeV2(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendOFTFeeV2", _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendOFTFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFTFeeV2(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTFeeV2 is a paid mutator transaction binding the contract method 0x85154849.
//
// Solidity: function sendOFTFeeV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendOFTFeeV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFTFeeV2(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendOFTV2(opts *bind.TransactOpts, _oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendOFTV2", _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendOFTV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFTV2(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendOFTV2 is a paid mutator transaction binding the contract method 0xa8198c00.
//
// Solidity: function sendOFTV2(address _oft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendOFTV2(_oft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendOFTV2(&_Storage.TransactOpts, _oft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendProxyOFT(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendProxyOFT", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendProxyOFT(_proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFT(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFT is a paid mutator transaction binding the contract method 0xc3c8032a.
//
// Solidity: function sendProxyOFT(address _proxyOft, uint16 _dstChainId, bytes _toAddress, uint256 _amount, uint256 _minAmount, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendProxyOFT(_proxyOft common.Address, _dstChainId uint16, _toAddress []byte, _amount *big.Int, _minAmount *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFT(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _refundAddress, _zroPaymentAddress, _adapterParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendProxyOFTFeeV2(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendProxyOFTFeeV2", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendProxyOFTFeeV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFTFeeV2(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTFeeV2 is a paid mutator transaction binding the contract method 0x8bcb586c.
//
// Solidity: function sendProxyOFTFeeV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendProxyOFTFeeV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFTFeeV2(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactor) SendProxyOFTV2(opts *bind.TransactOpts, _proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sendProxyOFTV2", _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageSession) SendProxyOFTV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFTV2(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SendProxyOFTV2 is a paid mutator transaction binding the contract method 0xdda16a10.
//
// Solidity: function sendProxyOFTV2(address _proxyOft, uint16 _dstChainId, bytes32 _toAddress, uint256 _amount, uint256 _minAmount, (address,address,bytes) _callParams, (uint256,address,bytes2) _feeObj) payable returns()
func (_Storage *StorageTransactorSession) SendProxyOFTV2(_proxyOft common.Address, _dstChainId uint16, _toAddress [32]byte, _amount *big.Int, _minAmount *big.Int, _callParams ICommonOFTLzCallParams, _feeObj IOFTWrapperFeeObj) (*types.Transaction, error) {
	return _Storage.Contract.SendProxyOFTV2(&_Storage.TransactOpts, _proxyOft, _dstChainId, _toAddress, _amount, _minAmount, _callParams, _feeObj)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Storage *StorageTransactor) SetDefaultBps(opts *bind.TransactOpts, _defaultBps *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDefaultBps", _defaultBps)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Storage *StorageSession) SetDefaultBps(_defaultBps *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDefaultBps(&_Storage.TransactOpts, _defaultBps)
}

// SetDefaultBps is a paid mutator transaction binding the contract method 0xa46d74bc.
//
// Solidity: function setDefaultBps(uint256 _defaultBps) returns()
func (_Storage *StorageTransactorSession) SetDefaultBps(_defaultBps *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetDefaultBps(&_Storage.TransactOpts, _defaultBps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Storage *StorageTransactor) SetOFTBps(opts *bind.TransactOpts, _token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setOFTBps", _token, _bps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Storage *StorageSession) SetOFTBps(_token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetOFTBps(&_Storage.TransactOpts, _token, _bps)
}

// SetOFTBps is a paid mutator transaction binding the contract method 0x0c3d2756.
//
// Solidity: function setOFTBps(address _token, uint256 _bps) returns()
func (_Storage *StorageTransactorSession) SetOFTBps(_token common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetOFTBps(&_Storage.TransactOpts, _token, _bps)
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

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Storage *StorageTransactor) WithdrawFees(opts *bind.TransactOpts, _oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawFees", _oft, _to, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Storage *StorageSession) WithdrawFees(_oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawFees(&_Storage.TransactOpts, _oft, _to, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xe55dc4e6.
//
// Solidity: function withdrawFees(address _oft, address _to, uint256 _amount) returns()
func (_Storage *StorageTransactorSession) WithdrawFees(_oft common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawFees(&_Storage.TransactOpts, _oft, _to, _amount)
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

// StorageWrapperFeeWithdrawnIterator is returned from FilterWrapperFeeWithdrawn and is used to iterate over the raw logs and unpacked data for WrapperFeeWithdrawn events raised by the Storage contract.
type StorageWrapperFeeWithdrawnIterator struct {
	Event *StorageWrapperFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StorageWrapperFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageWrapperFeeWithdrawn)
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
		it.Event = new(StorageWrapperFeeWithdrawn)
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
func (it *StorageWrapperFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageWrapperFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageWrapperFeeWithdrawn represents a WrapperFeeWithdrawn event raised by the Storage contract.
type StorageWrapperFeeWithdrawn struct {
	Oft    common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWrapperFeeWithdrawn is a free log retrieval operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Storage *StorageFilterer) FilterWrapperFeeWithdrawn(opts *bind.FilterOpts, oft []common.Address) (*StorageWrapperFeeWithdrawnIterator, error) {

	var oftRule []interface{}
	for _, oftItem := range oft {
		oftRule = append(oftRule, oftItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "WrapperFeeWithdrawn", oftRule)
	if err != nil {
		return nil, err
	}
	return &StorageWrapperFeeWithdrawnIterator{contract: _Storage.contract, event: "WrapperFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchWrapperFeeWithdrawn is a free log subscription operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Storage *StorageFilterer) WatchWrapperFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *StorageWrapperFeeWithdrawn, oft []common.Address) (event.Subscription, error) {

	var oftRule []interface{}
	for _, oftItem := range oft {
		oftRule = append(oftRule, oftItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "WrapperFeeWithdrawn", oftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageWrapperFeeWithdrawn)
				if err := _Storage.contract.UnpackLog(event, "WrapperFeeWithdrawn", log); err != nil {
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

// ParseWrapperFeeWithdrawn is a log parse operation binding the contract event 0xf6514f9f283faac4cf3f3a6a702c116227ad0f2c727fb336e4c10b418bc6d613.
//
// Solidity: event WrapperFeeWithdrawn(address indexed oft, address to, uint256 amount)
func (_Storage *StorageFilterer) ParseWrapperFeeWithdrawn(log types.Log) (*StorageWrapperFeeWithdrawn, error) {
	event := new(StorageWrapperFeeWithdrawn)
	if err := _Storage.contract.UnpackLog(event, "WrapperFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageWrapperFeesIterator is returned from FilterWrapperFees and is used to iterate over the raw logs and unpacked data for WrapperFees events raised by the Storage contract.
type StorageWrapperFeesIterator struct {
	Event *StorageWrapperFees // Event containing the contract specifics and raw log

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
func (it *StorageWrapperFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageWrapperFees)
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
		it.Event = new(StorageWrapperFees)
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
func (it *StorageWrapperFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageWrapperFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageWrapperFees represents a WrapperFees event raised by the Storage contract.
type StorageWrapperFees struct {
	PartnerId  [2]byte
	Token      common.Address
	WrapperFee *big.Int
	CallerFee  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWrapperFees is a free log retrieval operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageFilterer) FilterWrapperFees(opts *bind.FilterOpts, partnerId [][2]byte) (*StorageWrapperFeesIterator, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "WrapperFees", partnerIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageWrapperFeesIterator{contract: _Storage.contract, event: "WrapperFees", logs: logs, sub: sub}, nil
}

// WatchWrapperFees is a free log subscription operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageFilterer) WatchWrapperFees(opts *bind.WatchOpts, sink chan<- *StorageWrapperFees, partnerId [][2]byte) (event.Subscription, error) {

	var partnerIdRule []interface{}
	for _, partnerIdItem := range partnerId {
		partnerIdRule = append(partnerIdRule, partnerIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "WrapperFees", partnerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageWrapperFees)
				if err := _Storage.contract.UnpackLog(event, "WrapperFees", log); err != nil {
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

// ParseWrapperFees is a log parse operation binding the contract event 0x97bcdc1dd7ab82ef93280983f23d391afea463d0333fddd1a4617693b9ccfeea.
//
// Solidity: event WrapperFees(bytes2 indexed partnerId, address token, uint256 wrapperFee, uint256 callerFee)
func (_Storage *StorageFilterer) ParseWrapperFees(log types.Log) (*StorageWrapperFees, error) {
	event := new(StorageWrapperFees)
	if err := _Storage.contract.UnpackLog(event, "WrapperFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
