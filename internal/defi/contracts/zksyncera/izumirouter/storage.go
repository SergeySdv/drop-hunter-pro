// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izumirouter

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

// SwapSwapAmountParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapAmountParams struct {
	Path        []byte
	Recipient   common.Address
	Amount      *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// SwapSwapDesireParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapDesireParams struct {
	Path      []byte
	Recipient common.Address
	Desire    *big.Int
	MaxPayed  *big.Int
	Deadline  *big.Int
}

// SwapSwapParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapParams struct {
	TokenX      common.Address
	TokenY      common.Address
	Fee         *big.Int
	BoundaryPt  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	MaxPayed    *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapAmountParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desire\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapDesireParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapDesire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2Y\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2YCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2X\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2XCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Storage *StorageCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Storage *StorageSession) WETH9() (common.Address, error) {
	return _Storage.Contract.WETH9(&_Storage.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Storage *StorageCallerSession) WETH9() (common.Address, error) {
	return _Storage.Contract.WETH9(&_Storage.CallOpts)
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

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Storage *StorageCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Storage *StorageSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Storage.Contract.Pool(&_Storage.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Storage *StorageCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Storage.Contract.Pool(&_Storage.CallOpts, tokenX, tokenY, fee)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Storage *StorageTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Storage *StorageSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Storage.Contract.Multicall(&_Storage.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Storage *StorageTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Storage.Contract.Multicall(&_Storage.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Storage *StorageTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Storage *StorageSession) RefundETH() (*types.Transaction, error) {
	return _Storage.Contract.RefundETH(&_Storage.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Storage *StorageTransactorSession) RefundETH() (*types.Transaction, error) {
	return _Storage.Contract.RefundETH(&_Storage.TransactOpts)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageTransactor) SwapAmount(opts *bind.TransactOpts, params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapAmount", params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageSession) SwapAmount(params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapAmount(&_Storage.TransactOpts, params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageTransactorSession) SwapAmount(params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapAmount(&_Storage.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageTransactor) SwapDesire(opts *bind.TransactOpts, params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapDesire", params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageSession) SwapDesire(params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapDesire(&_Storage.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Storage *StorageTransactorSession) SwapDesire(params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapDesire(&_Storage.TransactOpts, params)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactor) SwapX2Y(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2Y", swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageSession) SwapX2Y(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactorSession) SwapX2Y(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, swapParams)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageTransactor) SwapX2YCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2YCallback", x, y, data)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageSession) SwapX2YCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YCallback(&_Storage.TransactOpts, x, y, data)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageTransactorSession) SwapX2YCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YCallback(&_Storage.TransactOpts, x, y, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactor) SwapX2YDesireY(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2YDesireY", swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageSession) SwapX2YDesireY(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactorSession) SwapX2YDesireY(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactor) SwapY2X(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2X", swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageSession) SwapY2X(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactorSession) SwapY2X(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, swapParams)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageTransactor) SwapY2XCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2XCallback", x, y, data)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageSession) SwapY2XCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XCallback(&_Storage.TransactOpts, x, y, data)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Storage *StorageTransactorSession) SwapY2XCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XCallback(&_Storage.TransactOpts, x, y, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactor) SwapY2XDesireX(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2XDesireX", swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageSession) SwapY2XDesireX(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Storage *StorageTransactorSession) SwapY2XDesireX(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, swapParams)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sweepToken", token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SweepToken(&_Storage.TransactOpts, token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageTransactorSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SweepToken(&_Storage.TransactOpts, token, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageTransactor) UnwrapWETH9(opts *bind.TransactOpts, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "unwrapWETH9", minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UnwrapWETH9(&_Storage.TransactOpts, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Storage *StorageTransactorSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UnwrapWETH9(&_Storage.TransactOpts, minAmount, recipient)
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
