// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maverickrouter

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

// IPoolAddLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolAddLiquidityParams struct {
	Kind    uint8
	Pos     int32
	IsDelta bool
	DeltaA  *big.Int
	DeltaB  *big.Int
}

// IPoolBinDelta is an auto generated low-level Go binding around an user-defined struct.
type IPoolBinDelta struct {
	DeltaA         *big.Int
	DeltaB         *big.Int
	DeltaLpBalance *big.Int
	BinId          *big.Int
	Kind           uint8
	LowerTick      int32
	IsActive       bool
}

// IPoolRemoveLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRemoveLiquidityParams struct {
	BinId  *big.Int
	Amount *big.Int
}

// IRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IRouterPoolParams is an auto generated low-level Go binding around an user-defined struct.
type IRouterPoolParams struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Lookback    *big.Int
	ActiveTick  int32
	TokenA      common.Address
	TokenB      common.Address
}

// ISlimRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISlimRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Pool              common.Address
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitD18 *big.Int
}

// ISlimRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISlimRouterExactOutputSingleParams struct {
	TokenIn         common.Address
	TokenOut        common.Address
	Pool            common.Address
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"contractIWETH9\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"addLiquidityCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"pos\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isDelta\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.AddLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokenBAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityToPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivingTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"pos\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isDelta\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.AddLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokenBAmount\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"minActiveTick\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"maxActiveTick\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityWTickLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivingTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sqrtPriceLimitD18\",\"type\":\"uint256\"}],\"internalType\":\"structISlimRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISlimRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tickSpacing\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"lookback\",\"type\":\"int256\"},{\"internalType\":\"int32\",\"name\":\"activeTick\",\"type\":\"int32\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"}],\"internalType\":\"structIRouter.PoolParams\",\"name\":\"poolParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"pos\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isDelta\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.AddLiquidityParams[]\",\"name\":\"addParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokenBAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"getOrCreatePoolAndAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivingTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint128[]\",\"name\":\"binIds\",\"type\":\"uint128[]\"},{\"internalType\":\"uint32\",\"name\":\"maxRecursion\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"migrateBinsUpStack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position\",\"outputs\":[{\"internalType\":\"contractIPosition\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.RemoveLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"minTokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokenBAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToPay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"swapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Storage *StorageCaller) Position(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "position")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Storage *StorageSession) Position() (common.Address, error) {
	return _Storage.Contract.Position(&_Storage.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x09218e91.
//
// Solidity: function position() view returns(address)
func (_Storage *StorageCallerSession) Position() (common.Address, error) {
	return _Storage.Contract.Position(&_Storage.CallOpts)
}

// AddLiquidityCallback is a paid mutator transaction binding the contract method 0xdc8fd182.
//
// Solidity: function addLiquidityCallback(uint256 amountA, uint256 amountB, bytes _data) returns()
func (_Storage *StorageTransactor) AddLiquidityCallback(opts *bind.TransactOpts, amountA *big.Int, amountB *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidityCallback", amountA, amountB, _data)
}

// AddLiquidityCallback is a paid mutator transaction binding the contract method 0xdc8fd182.
//
// Solidity: function addLiquidityCallback(uint256 amountA, uint256 amountB, bytes _data) returns()
func (_Storage *StorageSession) AddLiquidityCallback(amountA *big.Int, amountB *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityCallback(&_Storage.TransactOpts, amountA, amountB, _data)
}

// AddLiquidityCallback is a paid mutator transaction binding the contract method 0xdc8fd182.
//
// Solidity: function addLiquidityCallback(uint256 amountA, uint256 amountB, bytes _data) returns()
func (_Storage *StorageTransactorSession) AddLiquidityCallback(amountA *big.Int, amountB *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityCallback(&_Storage.TransactOpts, amountA, amountB, _data)
}

// AddLiquidityToPool is a paid mutator transaction binding the contract method 0x79b28ef3.
//
// Solidity: function addLiquidityToPool(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactor) AddLiquidityToPool(opts *bind.TransactOpts, pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidityToPool", pool, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// AddLiquidityToPool is a paid mutator transaction binding the contract method 0x79b28ef3.
//
// Solidity: function addLiquidityToPool(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageSession) AddLiquidityToPool(pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityToPool(&_Storage.TransactOpts, pool, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// AddLiquidityToPool is a paid mutator transaction binding the contract method 0x79b28ef3.
//
// Solidity: function addLiquidityToPool(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactorSession) AddLiquidityToPool(pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityToPool(&_Storage.TransactOpts, pool, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// AddLiquidityWTickLimits is a paid mutator transaction binding the contract method 0x26b3e1cc.
//
// Solidity: function addLiquidityWTickLimits(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, int32 minActiveTick, int32 maxActiveTick, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactor) AddLiquidityWTickLimits(opts *bind.TransactOpts, pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, minActiveTick int32, maxActiveTick int32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLiquidityWTickLimits", pool, tokenId, params, minTokenAAmount, minTokenBAmount, minActiveTick, maxActiveTick, deadline)
}

// AddLiquidityWTickLimits is a paid mutator transaction binding the contract method 0x26b3e1cc.
//
// Solidity: function addLiquidityWTickLimits(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, int32 minActiveTick, int32 maxActiveTick, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageSession) AddLiquidityWTickLimits(pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, minActiveTick int32, maxActiveTick int32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityWTickLimits(&_Storage.TransactOpts, pool, tokenId, params, minTokenAAmount, minTokenBAmount, minActiveTick, maxActiveTick, deadline)
}

// AddLiquidityWTickLimits is a paid mutator transaction binding the contract method 0x26b3e1cc.
//
// Solidity: function addLiquidityWTickLimits(address pool, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, int32 minActiveTick, int32 maxActiveTick, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactorSession) AddLiquidityWTickLimits(pool common.Address, tokenId *big.Int, params []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, minActiveTick int32, maxActiveTick int32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddLiquidityWTickLimits(&_Storage.TransactOpts, pool, tokenId, params, minTokenAAmount, minTokenBAmount, minActiveTick, maxActiveTick, deadline)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageTransactor) ExactInput(opts *bind.TransactOpts, params IRouterExactInputParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageSession) ExactInput(params IRouterExactInputParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactInput(&_Storage.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageTransactorSession) ExactInput(params IRouterExactInputParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactInput(&_Storage.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xa5dcbcdf.
//
// Solidity: function exactInputSingle((address,address,address,address,uint256,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISlimRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xa5dcbcdf.
//
// Solidity: function exactInputSingle((address,address,address,address,uint256,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageSession) ExactInputSingle(params ISlimRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactInputSingle(&_Storage.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xa5dcbcdf.
//
// Solidity: function exactInputSingle((address,address,address,address,uint256,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Storage *StorageTransactorSession) ExactInputSingle(params ISlimRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactInputSingle(&_Storage.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageTransactor) ExactOutput(opts *bind.TransactOpts, params IRouterExactOutputParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageSession) ExactOutput(params IRouterExactOutputParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactOutput(&_Storage.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageTransactorSession) ExactOutput(params IRouterExactOutputParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactOutput(&_Storage.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x27dbbf70.
//
// Solidity: function exactOutputSingle((address,address,address,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISlimRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x27dbbf70.
//
// Solidity: function exactOutputSingle((address,address,address,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageSession) ExactOutputSingle(params ISlimRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactOutputSingle(&_Storage.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x27dbbf70.
//
// Solidity: function exactOutputSingle((address,address,address,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Storage *StorageTransactorSession) ExactOutputSingle(params ISlimRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Storage.Contract.ExactOutputSingle(&_Storage.TransactOpts, params)
}

// GetOrCreatePoolAndAddLiquidity is a paid mutator transaction binding the contract method 0xcc1695c9.
//
// Solidity: function getOrCreatePoolAndAddLiquidity((uint256,uint256,int256,int32,address,address) poolParams, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] addParams, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactor) GetOrCreatePoolAndAddLiquidity(opts *bind.TransactOpts, poolParams IRouterPoolParams, tokenId *big.Int, addParams []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "getOrCreatePoolAndAddLiquidity", poolParams, tokenId, addParams, minTokenAAmount, minTokenBAmount, deadline)
}

// GetOrCreatePoolAndAddLiquidity is a paid mutator transaction binding the contract method 0xcc1695c9.
//
// Solidity: function getOrCreatePoolAndAddLiquidity((uint256,uint256,int256,int32,address,address) poolParams, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] addParams, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageSession) GetOrCreatePoolAndAddLiquidity(poolParams IRouterPoolParams, tokenId *big.Int, addParams []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.GetOrCreatePoolAndAddLiquidity(&_Storage.TransactOpts, poolParams, tokenId, addParams, minTokenAAmount, minTokenBAmount, deadline)
}

// GetOrCreatePoolAndAddLiquidity is a paid mutator transaction binding the contract method 0xcc1695c9.
//
// Solidity: function getOrCreatePoolAndAddLiquidity((uint256,uint256,int256,int32,address,address) poolParams, uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] addParams, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) payable returns(uint256 receivingTokenId, uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactorSession) GetOrCreatePoolAndAddLiquidity(poolParams IRouterPoolParams, tokenId *big.Int, addParams []IPoolAddLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.GetOrCreatePoolAndAddLiquidity(&_Storage.TransactOpts, poolParams, tokenId, addParams, minTokenAAmount, minTokenBAmount, deadline)
}

// MigrateBinsUpStack is a paid mutator transaction binding the contract method 0x3339aad1.
//
// Solidity: function migrateBinsUpStack(address pool, uint128[] binIds, uint32 maxRecursion, uint256 deadline) returns()
func (_Storage *StorageTransactor) MigrateBinsUpStack(opts *bind.TransactOpts, pool common.Address, binIds []*big.Int, maxRecursion uint32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "migrateBinsUpStack", pool, binIds, maxRecursion, deadline)
}

// MigrateBinsUpStack is a paid mutator transaction binding the contract method 0x3339aad1.
//
// Solidity: function migrateBinsUpStack(address pool, uint128[] binIds, uint32 maxRecursion, uint256 deadline) returns()
func (_Storage *StorageSession) MigrateBinsUpStack(pool common.Address, binIds []*big.Int, maxRecursion uint32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MigrateBinsUpStack(&_Storage.TransactOpts, pool, binIds, maxRecursion, deadline)
}

// MigrateBinsUpStack is a paid mutator transaction binding the contract method 0x3339aad1.
//
// Solidity: function migrateBinsUpStack(address pool, uint128[] binIds, uint32 maxRecursion, uint256 deadline) returns()
func (_Storage *StorageTransactorSession) MigrateBinsUpStack(pool common.Address, binIds []*big.Int, maxRecursion uint32, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MigrateBinsUpStack(&_Storage.TransactOpts, pool, binIds, maxRecursion, deadline)
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

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x29ffb48c.
//
// Solidity: function removeLiquidity(address pool, address recipient, uint256 tokenId, (uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactor) RemoveLiquidity(opts *bind.TransactOpts, pool common.Address, recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeLiquidity", pool, recipient, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x29ffb48c.
//
// Solidity: function removeLiquidity(address pool, address recipient, uint256 tokenId, (uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageSession) RemoveLiquidity(pool common.Address, recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, pool, recipient, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x29ffb48c.
//
// Solidity: function removeLiquidity(address pool, address recipient, uint256 tokenId, (uint128,uint128)[] params, uint256 minTokenAAmount, uint256 minTokenBAmount, uint256 deadline) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Storage *StorageTransactorSession) RemoveLiquidity(pool common.Address, recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams, minTokenAAmount *big.Int, minTokenBAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RemoveLiquidity(&_Storage.TransactOpts, pool, recipient, tokenId, params, minTokenAAmount, minTokenBAmount, deadline)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermit(&_Storage.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermit(&_Storage.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitAllowed(&_Storage.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitAllowed(&_Storage.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitAllowedIfNecessary(&_Storage.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitAllowedIfNecessary(&_Storage.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitIfNecessary(&_Storage.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Storage *StorageTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SelfPermitIfNecessary(&_Storage.TransactOpts, token, value, deadline, v, r, s)
}

// SwapCallback is a paid mutator transaction binding the contract method 0x923b8a2a.
//
// Solidity: function swapCallback(uint256 amountToPay, uint256 amountOut, bytes _data) returns()
func (_Storage *StorageTransactor) SwapCallback(opts *bind.TransactOpts, amountToPay *big.Int, amountOut *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapCallback", amountToPay, amountOut, _data)
}

// SwapCallback is a paid mutator transaction binding the contract method 0x923b8a2a.
//
// Solidity: function swapCallback(uint256 amountToPay, uint256 amountOut, bytes _data) returns()
func (_Storage *StorageSession) SwapCallback(amountToPay *big.Int, amountOut *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapCallback(&_Storage.TransactOpts, amountToPay, amountOut, _data)
}

// SwapCallback is a paid mutator transaction binding the contract method 0x923b8a2a.
//
// Solidity: function swapCallback(uint256 amountToPay, uint256 amountOut, bytes _data) returns()
func (_Storage *StorageTransactorSession) SwapCallback(amountToPay *big.Int, amountOut *big.Int, _data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapCallback(&_Storage.TransactOpts, amountToPay, amountOut, _data)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SweepToken(&_Storage.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SweepToken(&_Storage.TransactOpts, token, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UnwrapWETH9(&_Storage.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Storage *StorageTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UnwrapWETH9(&_Storage.TransactOpts, amountMinimum, recipient)
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
