// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izumiquoter

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"},{\"internalType\":\"int24[]\",\"name\":\"pointAfterList\",\"type\":\"int24[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"desire\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapDesire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"int24[]\",\"name\":\"pointAfterList\",\"type\":\"int24[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"}],\"name\":\"swapX2Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapX2YCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"desireY\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"}],\"name\":\"swapY2X\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapY2XCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"desireX\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// SwapAmount is a free data retrieval call binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) view returns(uint256 acquire, int24[] pointAfterList)
func (_Storage *StorageCaller) SwapAmount(opts *bind.CallOpts, amount *big.Int, path []byte) (struct {
	Acquire        *big.Int
	PointAfterList []*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "swapAmount", amount, path)

	outstruct := new(struct {
		Acquire        *big.Int
		PointAfterList []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Acquire = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PointAfterList = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// SwapAmount is a free data retrieval call binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) view returns(uint256 acquire, int24[] pointAfterList)
func (_Storage *StorageSession) SwapAmount(amount *big.Int, path []byte) (struct {
	Acquire        *big.Int
	PointAfterList []*big.Int
}, error) {
	return _Storage.Contract.SwapAmount(&_Storage.CallOpts, amount, path)
}

// SwapAmount is a free data retrieval call binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) view returns(uint256 acquire, int24[] pointAfterList)
func (_Storage *StorageCallerSession) SwapAmount(amount *big.Int, path []byte) (struct {
	Acquire        *big.Int
	PointAfterList []*big.Int
}, error) {
	return _Storage.Contract.SwapAmount(&_Storage.CallOpts, amount, path)
}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageCaller) SwapX2YCallback(opts *bind.CallOpts, x *big.Int, y *big.Int, path []byte) error {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "swapX2YCallback", x, y, path)

	if err != nil {
		return err
	}

	return err

}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageSession) SwapX2YCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Storage.Contract.SwapX2YCallback(&_Storage.CallOpts, x, y, path)
}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageCallerSession) SwapX2YCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Storage.Contract.SwapX2YCallback(&_Storage.CallOpts, x, y, path)
}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageCaller) SwapY2XCallback(opts *bind.CallOpts, x *big.Int, y *big.Int, path []byte) error {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "swapY2XCallback", x, y, path)

	if err != nil {
		return err
	}

	return err

}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageSession) SwapY2XCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Storage.Contract.SwapY2XCallback(&_Storage.CallOpts, x, y, path)
}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Storage *StorageCallerSession) SwapY2XCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Storage.Contract.SwapY2XCallback(&_Storage.CallOpts, x, y, path)
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

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Storage *StorageTransactor) SwapDesire(opts *bind.TransactOpts, desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapDesire", desire, path)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Storage *StorageSession) SwapDesire(desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapDesire(&_Storage.TransactOpts, desire, path)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Storage *StorageTransactorSession) SwapDesire(desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapDesire(&_Storage.TransactOpts, desire, path)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageTransactor) SwapX2Y(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2Y", tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageSession) SwapX2Y(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageTransactorSession) SwapX2Y(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageTransactor) SwapX2YDesireY(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2YDesireY", tokenX, tokenY, fee, desireY, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageSession) SwapX2YDesireY(tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, tokenX, tokenY, fee, desireY, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageTransactorSession) SwapX2YDesireY(tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, tokenX, tokenY, fee, desireY, lowPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageTransactor) SwapY2X(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2X", tokenX, tokenY, fee, amount, highPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageSession) SwapY2X(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, tokenX, tokenY, fee, amount, highPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Storage *StorageTransactorSession) SwapY2X(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, tokenX, tokenY, fee, amount, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageTransactor) SwapY2XDesireX(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2XDesireX", tokenX, tokenY, fee, desireX, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageSession) SwapY2XDesireX(tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, tokenX, tokenY, fee, desireX, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Storage *StorageTransactorSession) SwapY2XDesireX(tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, tokenX, tokenY, fee, desireX, highPt)
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
