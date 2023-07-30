// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package muteiorouter

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOutExpanded\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Storage *StorageCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Storage *StorageSession) WETH() (common.Address, error) {
	return _Storage.Contract.WETH(&_Storage.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Storage *StorageCallerSession) WETH() (common.Address, error) {
	return _Storage.Contract.WETH(&_Storage.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageSession) Factory() (common.Address, error) {
	return _Storage.Contract.Factory(&_Storage.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Storage *StorageCallerSession) Factory() (common.Address, error) {
	return _Storage.Contract.Factory(&_Storage.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_Storage *StorageCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

	outstruct := new(struct {
		AmountOut *big.Int
		Stable    bool
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_Storage *StorageSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _Storage.Contract.GetAmountOut(&_Storage.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_Storage *StorageCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _Storage.Contract.GetAmountOut(&_Storage.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_Storage *StorageCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAmountsOut", amountIn, path, stable)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_Storage *StorageSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	return _Storage.Contract.GetAmountsOut(&_Storage.CallOpts, amountIn, path, stable)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_Storage *StorageCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	return _Storage.Contract.GetAmountsOut(&_Storage.CallOpts, amountIn, path, stable)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_Storage *StorageCaller) GetAmountsOutExpanded(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAmountsOutExpanded", amountIn, path)

	outstruct := new(struct {
		Amounts []*big.Int
		Stable  []bool
		Fees    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.Fees = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_Storage *StorageSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _Storage.Contract.GetAmountsOutExpanded(&_Storage.CallOpts, amountIn, path)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_Storage *StorageCallerSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _Storage.Contract.GetAmountsOutExpanded(&_Storage.CallOpts, amountIn, path)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_Storage *StorageCaller) GetPairInfo(opts *bind.CallOpts, path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPairInfo", path, stable)

	outstruct := new(struct {
		TokenA   common.Address
		TokenB   common.Address
		Pair     common.Address
		ReserveA *big.Int
		ReserveB *big.Int
		Fee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Pair = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReserveA = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_Storage *StorageSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _Storage.Contract.GetPairInfo(&_Storage.CallOpts, path, stable)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_Storage *StorageCallerSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _Storage.Contract.GetPairInfo(&_Storage.CallOpts, path, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Storage *StorageCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Storage *StorageSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Storage.Contract.GetReserves(&_Storage.CallOpts, tokenA, tokenB, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_Storage *StorageCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Storage.Contract.GetReserves(&_Storage.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Storage *StorageCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Storage *StorageSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Storage.Contract.PairFor(&_Storage.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_Storage *StorageCallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _Storage.Contract.PairFor(&_Storage.CallOpts, tokenA, tokenB, stable)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Storage *StorageCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Storage *StorageSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Storage.Contract.Quote(&_Storage.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Storage *StorageCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Storage.Contract.Quote(&_Storage.CallOpts, amountA, reserveA, reserveB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Storage *StorageCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Storage *StorageSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Storage.Contract.SortTokens(&_Storage.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_Storage *StorageCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _Storage.Contract.SortTokens(&_Storage.CallOpts, tokenA, tokenB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Storage *StorageTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Storage *StorageSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidity(&_Storage.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Storage *StorageTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidity(&_Storage.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Storage *StorageTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Storage *StorageSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityETH(&_Storage.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Storage *StorageTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityETH(&_Storage.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_Storage *StorageTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_Storage *StorageSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_Storage *StorageTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_Storage *StorageTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_Storage *StorageSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityETH(&_Storage.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_Storage *StorageTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityETH(&_Storage.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_Storage *StorageTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_Storage *StorageSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Storage.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_Storage *StorageTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Storage.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_Storage *StorageTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_Storage *StorageSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactETHForTokens(&_Storage.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_Storage *StorageTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactETHForTokens(&_Storage.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_Storage *StorageTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_Storage *StorageSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_Storage *StorageTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForETH(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForETH(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_Storage *StorageTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_Storage *StorageTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _Storage.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Storage.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}
