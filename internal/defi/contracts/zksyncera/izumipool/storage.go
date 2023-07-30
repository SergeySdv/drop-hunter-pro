// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izumipool

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

// IiZiSwapPoolLimitOrderStruct is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapPoolLimitOrderStruct struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_pointDelta\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"AddLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"DecLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidY\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountX\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountY\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignX\",\"type\":\"uint128\"}],\"name\":\"assignLimOrderEarnX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignY\",\"type\":\"uint128\"}],\"name\":\"assignLimOrderEarnY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amountXLim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLim\",\"type\":\"uint256\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAmountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualAmountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFeeCharged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"collectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"collectEarn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isEarnY\",\"type\":\"bool\"}],\"name\":\"collectLimOrder\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualCollectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"actualCollectEarn\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaX\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaY\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newNextQueueLen\",\"type\":\"uint16\"}],\"name\":\"expandObservationQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeChargePercent\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeScaleX_128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeScaleY_128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leftMostPt\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"limitOrderData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"limitOrderSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwapPool.LimitOrderStruct[]\",\"name\":\"limitOrders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"liquiditySnapshot\",\"outputs\":[{\"internalType\":\"int128[]\",\"name\":\"deltaLiquidities\",\"type\":\"int128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidPt\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"accPoint\",\"type\":\"int56\"},{\"internalType\":\"bool\",\"name\":\"init\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"accPoints\",\"type\":\"int56[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"orderOrEndpoint\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"pointBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pointDelta\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"points\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidSum\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"accFeeXOut_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accFeeYOut_128\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEndpt\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rightMostPt\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sqrtRate_96\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPrice_96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationCurrentIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationNextQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityX\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireY\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2X\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireX\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeXCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeYCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userEarnX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userEarnY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Storage *StorageCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Storage *StorageSession) Fee() (*big.Int, error) {
	return _Storage.Contract.Fee(&_Storage.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Storage *StorageCallerSession) Fee() (*big.Int, error) {
	return _Storage.Contract.Fee(&_Storage.CallOpts)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Storage *StorageCaller) FeeChargePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "feeChargePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Storage *StorageSession) FeeChargePercent() (*big.Int, error) {
	return _Storage.Contract.FeeChargePercent(&_Storage.CallOpts)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Storage *StorageCallerSession) FeeChargePercent() (*big.Int, error) {
	return _Storage.Contract.FeeChargePercent(&_Storage.CallOpts)
}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Storage *StorageCaller) FeeScaleX128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "feeScaleX_128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Storage *StorageSession) FeeScaleX128() (*big.Int, error) {
	return _Storage.Contract.FeeScaleX128(&_Storage.CallOpts)
}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Storage *StorageCallerSession) FeeScaleX128() (*big.Int, error) {
	return _Storage.Contract.FeeScaleX128(&_Storage.CallOpts)
}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Storage *StorageCaller) FeeScaleY128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "feeScaleY_128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Storage *StorageSession) FeeScaleY128() (*big.Int, error) {
	return _Storage.Contract.FeeScaleY128(&_Storage.CallOpts)
}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Storage *StorageCallerSession) FeeScaleY128() (*big.Int, error) {
	return _Storage.Contract.FeeScaleY128(&_Storage.CallOpts)
}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Storage *StorageCaller) LeftMostPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "leftMostPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Storage *StorageSession) LeftMostPt() (*big.Int, error) {
	return _Storage.Contract.LeftMostPt(&_Storage.CallOpts)
}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Storage *StorageCallerSession) LeftMostPt() (*big.Int, error) {
	return _Storage.Contract.LeftMostPt(&_Storage.CallOpts)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint128 sellingY, uint128 earnX, uint256 accEarnX)
func (_Storage *StorageCaller) LimitOrderData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "limitOrderData", arg0)

	outstruct := new(struct {
		SellingX *big.Int
		EarnY    *big.Int
		AccEarnY *big.Int
		SellingY *big.Int
		EarnX    *big.Int
		AccEarnX *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellingX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EarnY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccEarnY = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SellingY = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EarnX = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccEarnX = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint128 sellingY, uint128 earnX, uint256 accEarnX)
func (_Storage *StorageSession) LimitOrderData(arg0 *big.Int) (struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}, error) {
	return _Storage.Contract.LimitOrderData(&_Storage.CallOpts, arg0)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint128 sellingY, uint128 earnX, uint256 accEarnX)
func (_Storage *StorageCallerSession) LimitOrderData(arg0 *big.Int) (struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}, error) {
	return _Storage.Contract.LimitOrderData(&_Storage.CallOpts, arg0)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Storage *StorageCaller) LimitOrderSnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "limitOrderSnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]IiZiSwapPoolLimitOrderStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]IiZiSwapPoolLimitOrderStruct)).(*[]IiZiSwapPoolLimitOrderStruct)

	return out0, err

}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Storage *StorageSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _Storage.Contract.LimitOrderSnapshot(&_Storage.CallOpts, leftPoint, rightPoint)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Storage *StorageCallerSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _Storage.Contract.LimitOrderSnapshot(&_Storage.CallOpts, leftPoint, rightPoint)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Storage *StorageCaller) Liquidity(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "liquidity", arg0)

	outstruct := new(struct {
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		TokenOwedX       *big.Int
		TokenOwedY       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedX = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Storage *StorageSession) Liquidity(arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _Storage.Contract.Liquidity(&_Storage.CallOpts, arg0)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Storage *StorageCallerSession) Liquidity(arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _Storage.Contract.Liquidity(&_Storage.CallOpts, arg0)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Storage *StorageCaller) LiquiditySnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "liquiditySnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Storage *StorageSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _Storage.Contract.LiquiditySnapshot(&_Storage.CallOpts, leftPoint, rightPoint)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Storage *StorageCallerSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _Storage.Contract.LiquiditySnapshot(&_Storage.CallOpts, leftPoint, rightPoint)
}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Storage *StorageCaller) MaxLiquidPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "maxLiquidPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Storage *StorageSession) MaxLiquidPt() (*big.Int, error) {
	return _Storage.Contract.MaxLiquidPt(&_Storage.CallOpts)
}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Storage *StorageCallerSession) MaxLiquidPt() (*big.Int, error) {
	return _Storage.Contract.MaxLiquidPt(&_Storage.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Storage *StorageCaller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "observations", arg0)

	outstruct := new(struct {
		Timestamp uint32
		AccPoint  *big.Int
		Init      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AccPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Init = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Storage *StorageSession) Observations(arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _Storage.Contract.Observations(&_Storage.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Storage *StorageCallerSession) Observations(arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _Storage.Contract.Observations(&_Storage.CallOpts, arg0)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Storage *StorageCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "observe", secondsAgos)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Storage *StorageSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _Storage.Contract.Observe(&_Storage.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Storage *StorageCallerSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _Storage.Contract.Observe(&_Storage.CallOpts, secondsAgos)
}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Storage *StorageCaller) OrderOrEndpoint(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "orderOrEndpoint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Storage *StorageSession) OrderOrEndpoint(arg0 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OrderOrEndpoint(&_Storage.CallOpts, arg0)
}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Storage *StorageCallerSession) OrderOrEndpoint(arg0 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OrderOrEndpoint(&_Storage.CallOpts, arg0)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Storage *StorageCaller) PointBitmap(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pointBitmap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Storage *StorageSession) PointBitmap(arg0 int16) (*big.Int, error) {
	return _Storage.Contract.PointBitmap(&_Storage.CallOpts, arg0)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Storage *StorageCallerSession) PointBitmap(arg0 int16) (*big.Int, error) {
	return _Storage.Contract.PointBitmap(&_Storage.CallOpts, arg0)
}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Storage *StorageCaller) PointDelta(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pointDelta")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Storage *StorageSession) PointDelta() (*big.Int, error) {
	return _Storage.Contract.PointDelta(&_Storage.CallOpts)
}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Storage *StorageCallerSession) PointDelta() (*big.Int, error) {
	return _Storage.Contract.PointDelta(&_Storage.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Storage *StorageCaller) Points(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "points", arg0)

	outstruct := new(struct {
		LiquidSum     *big.Int
		LiquidDelta   *big.Int
		AccFeeXOut128 *big.Int
		AccFeeYOut128 *big.Int
		IsEndpt       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidSum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccFeeXOut128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccFeeYOut128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsEndpt = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Storage *StorageSession) Points(arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _Storage.Contract.Points(&_Storage.CallOpts, arg0)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Storage *StorageCallerSession) Points(arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _Storage.Contract.Points(&_Storage.CallOpts, arg0)
}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Storage *StorageCaller) RightMostPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "rightMostPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Storage *StorageSession) RightMostPt() (*big.Int, error) {
	return _Storage.Contract.RightMostPt(&_Storage.CallOpts)
}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Storage *StorageCallerSession) RightMostPt() (*big.Int, error) {
	return _Storage.Contract.RightMostPt(&_Storage.CallOpts)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Storage *StorageCaller) SqrtRate96(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "sqrtRate_96")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Storage *StorageSession) SqrtRate96() (*big.Int, error) {
	return _Storage.Contract.SqrtRate96(&_Storage.CallOpts)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Storage *StorageCallerSession) SqrtRate96() (*big.Int, error) {
	return _Storage.Contract.SqrtRate96(&_Storage.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Storage *StorageCaller) State(opts *bind.CallOpts) (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SqrtPrice96             *big.Int
		CurrentPoint            *big.Int
		ObservationCurrentIndex uint16
		ObservationQueueLen     uint16
		ObservationNextQueueLen uint16
		Locked                  bool
		Liquidity               *big.Int
		LiquidityX              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationCurrentIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationQueueLen = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationNextQueueLen = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Locked = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Liquidity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Storage *StorageSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _Storage.Contract.State(&_Storage.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Storage *StorageCallerSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _Storage.Contract.State(&_Storage.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Storage *StorageCaller) TokenX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "tokenX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Storage *StorageSession) TokenX() (common.Address, error) {
	return _Storage.Contract.TokenX(&_Storage.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Storage *StorageCallerSession) TokenX() (common.Address, error) {
	return _Storage.Contract.TokenX(&_Storage.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Storage *StorageCaller) TokenY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "tokenY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Storage *StorageSession) TokenY() (common.Address, error) {
	return _Storage.Contract.TokenY(&_Storage.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Storage *StorageCallerSession) TokenY() (common.Address, error) {
	return _Storage.Contract.TokenY(&_Storage.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Storage *StorageCaller) TotalFeeXCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalFeeXCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Storage *StorageSession) TotalFeeXCharged() (*big.Int, error) {
	return _Storage.Contract.TotalFeeXCharged(&_Storage.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Storage *StorageCallerSession) TotalFeeXCharged() (*big.Int, error) {
	return _Storage.Contract.TotalFeeXCharged(&_Storage.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Storage *StorageCaller) TotalFeeYCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalFeeYCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Storage *StorageSession) TotalFeeYCharged() (*big.Int, error) {
	return _Storage.Contract.TotalFeeYCharged(&_Storage.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Storage *StorageCallerSession) TotalFeeYCharged() (*big.Int, error) {
	return _Storage.Contract.TotalFeeYCharged(&_Storage.CallOpts)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageCaller) UserEarnX(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "userEarnX", arg0)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageSession) UserEarnX(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Storage.Contract.UserEarnX(&_Storage.CallOpts, arg0)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageCallerSession) UserEarnX(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Storage.Contract.UserEarnX(&_Storage.CallOpts, arg0)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageCaller) UserEarnY(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "userEarnY", arg0)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageSession) UserEarnY(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Storage.Contract.UserEarnY(&_Storage.CallOpts, arg0)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 earnAssign)
func (_Storage *StorageCallerSession) UserEarnY(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Storage.Contract.UserEarnY(&_Storage.CallOpts, arg0)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Storage *StorageTransactor) AddLimOrderWithX(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLimOrderWithX", recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Storage *StorageSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLimOrderWithX(&_Storage.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Storage *StorageTransactorSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLimOrderWithX(&_Storage.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Storage *StorageTransactor) AddLimOrderWithY(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addLimOrderWithY", recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Storage *StorageSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLimOrderWithY(&_Storage.TransactOpts, recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Storage *StorageTransactorSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.AddLimOrderWithY(&_Storage.TransactOpts, recipient, point, amountY, data)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xcedb2a6c.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX) returns(uint128 actualAssignX)
func (_Storage *StorageTransactor) AssignLimOrderEarnX(opts *bind.TransactOpts, point *big.Int, assignX *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "assignLimOrderEarnX", point, assignX)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xcedb2a6c.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX) returns(uint128 actualAssignX)
func (_Storage *StorageSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AssignLimOrderEarnX(&_Storage.TransactOpts, point, assignX)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xcedb2a6c.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX) returns(uint128 actualAssignX)
func (_Storage *StorageTransactorSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AssignLimOrderEarnX(&_Storage.TransactOpts, point, assignX)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0xd88a74d2.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY) returns(uint128 actualAssignY)
func (_Storage *StorageTransactor) AssignLimOrderEarnY(opts *bind.TransactOpts, point *big.Int, assignY *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "assignLimOrderEarnY", point, assignY)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0xd88a74d2.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY) returns(uint128 actualAssignY)
func (_Storage *StorageSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AssignLimOrderEarnY(&_Storage.TransactOpts, point, assignY)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0xd88a74d2.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY) returns(uint128 actualAssignY)
func (_Storage *StorageTransactorSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AssignLimOrderEarnY(&_Storage.TransactOpts, point, assignY)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) Burn(opts *bind.TransactOpts, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "burn", leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Burn(&_Storage.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Burn(&_Storage.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Storage *StorageTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "collect", recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Storage *StorageSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Collect(&_Storage.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Storage *StorageTransactorSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Collect(&_Storage.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Storage *StorageTransactor) CollectFeeCharged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "collectFeeCharged")
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Storage *StorageSession) CollectFeeCharged() (*types.Transaction, error) {
	return _Storage.Contract.CollectFeeCharged(&_Storage.TransactOpts)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Storage *StorageTransactorSession) CollectFeeCharged() (*types.Transaction, error) {
	return _Storage.Contract.CollectFeeCharged(&_Storage.TransactOpts)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Storage *StorageTransactor) CollectLimOrder(opts *bind.TransactOpts, recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "collectLimOrder", recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Storage *StorageSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Storage.Contract.CollectLimOrder(&_Storage.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Storage *StorageTransactorSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Storage.Contract.CollectLimOrder(&_Storage.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX)
func (_Storage *StorageTransactor) DecLimOrderWithX(opts *bind.TransactOpts, point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "decLimOrderWithX", point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX)
func (_Storage *StorageSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecLimOrderWithX(&_Storage.TransactOpts, point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX)
func (_Storage *StorageTransactorSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecLimOrderWithX(&_Storage.TransactOpts, point, deltaX)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY)
func (_Storage *StorageTransactor) DecLimOrderWithY(opts *bind.TransactOpts, point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "decLimOrderWithY", point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY)
func (_Storage *StorageSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecLimOrderWithY(&_Storage.TransactOpts, point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY)
func (_Storage *StorageTransactorSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DecLimOrderWithY(&_Storage.TransactOpts, point, deltaY)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Storage *StorageTransactor) ExpandObservationQueue(opts *bind.TransactOpts, newNextQueueLen uint16) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "expandObservationQueue", newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Storage *StorageSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _Storage.Contract.ExpandObservationQueue(&_Storage.TransactOpts, newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Storage *StorageTransactorSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _Storage.Contract.ExpandObservationQueue(&_Storage.TransactOpts, newNextQueueLen)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Storage *StorageTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "flash", recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Storage *StorageSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Flash(&_Storage.TransactOpts, recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Storage *StorageTransactorSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Flash(&_Storage.TransactOpts, recipient, amountX, amountY, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mint", recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Mint(&_Storage.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.Mint(&_Storage.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) SwapX2Y(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2Y", recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2Y(&_Storage.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) SwapX2YDesireY(opts *bind.TransactOpts, recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapX2YDesireY", recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapX2YDesireY(&_Storage.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) SwapY2X(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2X", recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2X(&_Storage.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactor) SwapY2XDesireX(opts *bind.TransactOpts, recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "swapY2XDesireX", recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Storage *StorageTransactorSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.SwapY2XDesireX(&_Storage.TransactOpts, recipient, desireX, highPt, data)
}

// StorageAddLimitOrderIterator is returned from FilterAddLimitOrder and is used to iterate over the raw logs and unpacked data for AddLimitOrder events raised by the Storage contract.
type StorageAddLimitOrderIterator struct {
	Event *StorageAddLimitOrder // Event containing the contract specifics and raw log

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
func (it *StorageAddLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddLimitOrder)
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
		it.Event = new(StorageAddLimitOrder)
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
func (it *StorageAddLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddLimitOrder represents a AddLimitOrder event raised by the Storage contract.
type StorageAddLimitOrder struct {
	Amount     *big.Int
	Point      *big.Int
	SellXEarnY bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLimitOrder is a free log retrieval operation binding the contract event 0x79f3811dd9e0c504721e5c7fb14b12345c3c4d493ccef81f3ce2c681e510893d.
//
// Solidity: event AddLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) FilterAddLimitOrder(opts *bind.FilterOpts) (*StorageAddLimitOrderIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddLimitOrder")
	if err != nil {
		return nil, err
	}
	return &StorageAddLimitOrderIterator{contract: _Storage.contract, event: "AddLimitOrder", logs: logs, sub: sub}, nil
}

// WatchAddLimitOrder is a free log subscription operation binding the contract event 0x79f3811dd9e0c504721e5c7fb14b12345c3c4d493ccef81f3ce2c681e510893d.
//
// Solidity: event AddLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) WatchAddLimitOrder(opts *bind.WatchOpts, sink chan<- *StorageAddLimitOrder) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddLimitOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddLimitOrder)
				if err := _Storage.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
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

// ParseAddLimitOrder is a log parse operation binding the contract event 0x79f3811dd9e0c504721e5c7fb14b12345c3c4d493ccef81f3ce2c681e510893d.
//
// Solidity: event AddLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) ParseAddLimitOrder(log types.Log) (*StorageAddLimitOrder, error) {
	event := new(StorageAddLimitOrder)
	if err := _Storage.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Storage contract.
type StorageBurnIterator struct {
	Event *StorageBurn // Event containing the contract specifics and raw log

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
func (it *StorageBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageBurn)
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
		it.Event = new(StorageBurn)
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
func (it *StorageBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageBurn represents a Burn event raised by the Storage contract.
type StorageBurn struct {
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*StorageBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &StorageBurnIterator{contract: _Storage.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StorageBurn, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageBurn)
				if err := _Storage.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) ParseBurn(log types.Log) (*StorageBurn, error) {
	event := new(StorageBurn)
	if err := _Storage.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDecLimitOrderIterator is returned from FilterDecLimitOrder and is used to iterate over the raw logs and unpacked data for DecLimitOrder events raised by the Storage contract.
type StorageDecLimitOrderIterator struct {
	Event *StorageDecLimitOrder // Event containing the contract specifics and raw log

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
func (it *StorageDecLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDecLimitOrder)
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
		it.Event = new(StorageDecLimitOrder)
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
func (it *StorageDecLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDecLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDecLimitOrder represents a DecLimitOrder event raised by the Storage contract.
type StorageDecLimitOrder struct {
	Amount     *big.Int
	Point      *big.Int
	SellXEarnY bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecLimitOrder is a free log retrieval operation binding the contract event 0x1337302f348e2dfa78ea23d72aeb8b90c33c3199ed8ea4ccc43bb4b8c8e05eea.
//
// Solidity: event DecLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) FilterDecLimitOrder(opts *bind.FilterOpts) (*StorageDecLimitOrderIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DecLimitOrder")
	if err != nil {
		return nil, err
	}
	return &StorageDecLimitOrderIterator{contract: _Storage.contract, event: "DecLimitOrder", logs: logs, sub: sub}, nil
}

// WatchDecLimitOrder is a free log subscription operation binding the contract event 0x1337302f348e2dfa78ea23d72aeb8b90c33c3199ed8ea4ccc43bb4b8c8e05eea.
//
// Solidity: event DecLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) WatchDecLimitOrder(opts *bind.WatchOpts, sink chan<- *StorageDecLimitOrder) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DecLimitOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDecLimitOrder)
				if err := _Storage.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
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

// ParseDecLimitOrder is a log parse operation binding the contract event 0x1337302f348e2dfa78ea23d72aeb8b90c33c3199ed8ea4ccc43bb4b8c8e05eea.
//
// Solidity: event DecLimitOrder(uint256 amount, int24 point, bool sellXEarnY)
func (_Storage *StorageFilterer) ParseDecLimitOrder(log types.Log) (*StorageDecLimitOrder, error) {
	event := new(StorageDecLimitOrder)
	if err := _Storage.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the Storage contract.
type StorageFlashIterator struct {
	Event *StorageFlash // Event containing the contract specifics and raw log

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
func (it *StorageFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFlash)
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
		it.Event = new(StorageFlash)
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
func (it *StorageFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFlash represents a Flash event raised by the Storage contract.
type StorageFlash struct {
	Sender    common.Address
	Recipient common.Address
	AmountX   *big.Int
	AmountY   *big.Int
	PaidX     *big.Int
	PaidY     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Storage *StorageFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*StorageFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StorageFlashIterator{contract: _Storage.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Storage *StorageFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *StorageFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFlash)
				if err := _Storage.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Storage *StorageFilterer) ParseFlash(log types.Log) (*StorageFlash, error) {
	event := new(StorageFlash)
	if err := _Storage.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Storage contract.
type StorageMintIterator struct {
	Event *StorageMint // Event containing the contract specifics and raw log

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
func (it *StorageMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMint)
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
		it.Event = new(StorageMint)
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
func (it *StorageMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMint represents a Mint event raised by the Storage contract.
type StorageMint struct {
	Sender     common.Address
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*StorageMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &StorageMintIterator{contract: _Storage.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *StorageMint, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMint)
				if err := _Storage.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) ParseMint(log types.Log) (*StorageMint, error) {
	event := new(StorageMint)
	if err := _Storage.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Storage contract.
type StorageSwapIterator struct {
	Event *StorageSwap // Event containing the contract specifics and raw log

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
func (it *StorageSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSwap)
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
		it.Event = new(StorageSwap)
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
func (it *StorageSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSwap represents a Swap event raised by the Storage contract.
type StorageSwap struct {
	TokenX     common.Address
	TokenY     common.Address
	Fee        *big.Int
	SellXEarnY bool
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xe7779a36a28ae0e49bcbd9fcf57286fb607699c0c339c202e92495640505613e.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) FilterSwap(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (*StorageSwapIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &StorageSwapIterator{contract: _Storage.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xe7779a36a28ae0e49bcbd9fcf57286fb607699c0c339c202e92495640505613e.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *StorageSwap, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSwap)
				if err := _Storage.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xe7779a36a28ae0e49bcbd9fcf57286fb607699c0c339c202e92495640505613e.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY)
func (_Storage *StorageFilterer) ParseSwap(log types.Log) (*StorageSwap, error) {
	event := new(StorageSwap)
	if err := _Storage.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
