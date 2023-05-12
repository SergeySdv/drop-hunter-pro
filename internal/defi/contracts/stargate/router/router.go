// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

// IStargateRouterlzTxObj is an auto generated low-level Go binding around an user-defined struct.
type IStargateRouterlzTxObj struct {
	DstGasForCall   *big.Int
	DstNativeAmount *big.Int
	DstNativeAddr   []byte
}

// PoolCreditObj is an auto generated low-level Go binding around an user-defined struct.
type PoolCreditObj struct {
	Credits      *big.Int
	IdealBalance *big.Int
}

// PoolSwapObj is an auto generated low-level Go binding around an user-defined struct.
type PoolSwapObj struct {
	Amount      *big.Int
	EqFee       *big.Int
	EqReward    *big.Int
	LpFee       *big.Int
	ProtocolFee *big.Int
	LkbRemove   *big.Int
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"CachedSwapSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmountSD\",\"type\":\"uint256\"}],\"name\":\"RedeemLocalCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"bridgeFunctionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Revert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmountSD\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"}],\"name\":\"RevertRedeemLocal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"activateChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cachedSwapLookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_fullMode\",\"type\":\"bool\"}],\"name\":\"callDelta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"clearCachedSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"createChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_sharedDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_localDecimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.CreditObj\",\"name\":\"_c\",\"type\":\"tuple\"}],\"name\":\"creditChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcPoolId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"instantRedeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_functionType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_transferAndCallPayload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargateRouter.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"quoteLayerZeroFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargateRouter.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"redeemLocal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintAmountSD\",\"type\":\"uint256\"}],\"name\":\"redeemLocalCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"}],\"name\":\"redeemLocalCheckOnRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargateRouter.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"redeemRemote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"retryRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"revertLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargateRouter.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"}],\"name\":\"revertRedeemLocal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendCredits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"contractFactory\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"setBridgeAndFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_batched\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_swapDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_defaultSwapMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_defaultLPMode\",\"type\":\"bool\"}],\"name\":\"setDeltaParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeLibraryAddr\",\"type\":\"address\"}],\"name\":\"setFeeLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintFeeBP\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setMintFeeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setProtocolFeeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_swapStop\",\"type\":\"bool\"}],\"name\":\"setSwapStop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_weight\",\"type\":\"uint16\"}],\"name\":\"setWeightForChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstNativeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dstNativeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structIStargateRouter.lzTxObj\",\"name\":\"_lzTxParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstGasForCall\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkbRemove\",\"type\":\"uint256\"}],\"internalType\":\"structPool.SwapObj\",\"name\":\"_s\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"swapRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Router *RouterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Router *RouterSession) Bridge() (common.Address, error) {
	return _Router.Contract.Bridge(&_Router.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Router *RouterCallerSession) Bridge() (common.Address, error) {
	return _Router.Contract.Bridge(&_Router.CallOpts)
}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 , bytes , uint256 ) view returns(address token, uint256 amountLD, address to, bytes payload)
func (_Router *RouterCaller) CachedSwapLookup(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 *big.Int) (struct {
	Token    common.Address
	AmountLD *big.Int
	To       common.Address
	Payload  []byte
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "cachedSwapLookup", arg0, arg1, arg2)

	outstruct := new(struct {
		Token    common.Address
		AmountLD *big.Int
		To       common.Address
		Payload  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AmountLD = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Payload = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 , bytes , uint256 ) view returns(address token, uint256 amountLD, address to, bytes payload)
func (_Router *RouterSession) CachedSwapLookup(arg0 uint16, arg1 []byte, arg2 *big.Int) (struct {
	Token    common.Address
	AmountLD *big.Int
	To       common.Address
	Payload  []byte
}, error) {
	return _Router.Contract.CachedSwapLookup(&_Router.CallOpts, arg0, arg1, arg2)
}

// CachedSwapLookup is a free data retrieval call binding the contract method 0x23fd4647.
//
// Solidity: function cachedSwapLookup(uint16 , bytes , uint256 ) view returns(address token, uint256 amountLD, address to, bytes payload)
func (_Router *RouterCallerSession) CachedSwapLookup(arg0 uint16, arg1 []byte, arg2 *big.Int) (struct {
	Token    common.Address
	AmountLD *big.Int
	To       common.Address
	Payload  []byte
}, error) {
	return _Router.Contract.CachedSwapLookup(&_Router.CallOpts, arg0, arg1, arg2)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// MintFeeOwner is a free data retrieval call binding the contract method 0xaf640d82.
//
// Solidity: function mintFeeOwner() view returns(address)
func (_Router *RouterCaller) MintFeeOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "mintFeeOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintFeeOwner is a free data retrieval call binding the contract method 0xaf640d82.
//
// Solidity: function mintFeeOwner() view returns(address)
func (_Router *RouterSession) MintFeeOwner() (common.Address, error) {
	return _Router.Contract.MintFeeOwner(&_Router.CallOpts)
}

// MintFeeOwner is a free data retrieval call binding the contract method 0xaf640d82.
//
// Solidity: function mintFeeOwner() view returns(address)
func (_Router *RouterCallerSession) MintFeeOwner() (common.Address, error) {
	return _Router.Contract.MintFeeOwner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// ProtocolFeeOwner is a free data retrieval call binding the contract method 0xa96fbed4.
//
// Solidity: function protocolFeeOwner() view returns(address)
func (_Router *RouterCaller) ProtocolFeeOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "protocolFeeOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeOwner is a free data retrieval call binding the contract method 0xa96fbed4.
//
// Solidity: function protocolFeeOwner() view returns(address)
func (_Router *RouterSession) ProtocolFeeOwner() (common.Address, error) {
	return _Router.Contract.ProtocolFeeOwner(&_Router.CallOpts)
}

// ProtocolFeeOwner is a free data retrieval call binding the contract method 0xa96fbed4.
//
// Solidity: function protocolFeeOwner() view returns(address)
func (_Router *RouterCallerSession) ProtocolFeeOwner() (common.Address, error) {
	return _Router.Contract.ProtocolFeeOwner(&_Router.CallOpts)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_Router *RouterCaller) QuoteLayerZeroFee(opts *bind.CallOpts, _dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargateRouterlzTxObj) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "quoteLayerZeroFee", _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_Router *RouterSession) QuoteLayerZeroFee(_dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargateRouterlzTxObj) (*big.Int, *big.Int, error) {
	return _Router.Contract.QuoteLayerZeroFee(&_Router.CallOpts, _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)
}

// QuoteLayerZeroFee is a free data retrieval call binding the contract method 0x0a512369.
//
// Solidity: function quoteLayerZeroFee(uint16 _dstChainId, uint8 _functionType, bytes _toAddress, bytes _transferAndCallPayload, (uint256,uint256,bytes) _lzTxParams) view returns(uint256, uint256)
func (_Router *RouterCallerSession) QuoteLayerZeroFee(_dstChainId uint16, _functionType uint8, _toAddress []byte, _transferAndCallPayload []byte, _lzTxParams IStargateRouterlzTxObj) (*big.Int, *big.Int, error) {
	return _Router.Contract.QuoteLayerZeroFee(&_Router.CallOpts, _dstChainId, _functionType, _toAddress, _transferAndCallPayload, _lzTxParams)
}

// RevertLookup is a free data retrieval call binding the contract method 0xda133a06.
//
// Solidity: function revertLookup(uint16 , bytes , uint256 ) view returns(bytes)
func (_Router *RouterCaller) RevertLookup(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "revertLookup", arg0, arg1, arg2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RevertLookup is a free data retrieval call binding the contract method 0xda133a06.
//
// Solidity: function revertLookup(uint16 , bytes , uint256 ) view returns(bytes)
func (_Router *RouterSession) RevertLookup(arg0 uint16, arg1 []byte, arg2 *big.Int) ([]byte, error) {
	return _Router.Contract.RevertLookup(&_Router.CallOpts, arg0, arg1, arg2)
}

// RevertLookup is a free data retrieval call binding the contract method 0xda133a06.
//
// Solidity: function revertLookup(uint16 , bytes , uint256 ) view returns(bytes)
func (_Router *RouterCallerSession) RevertLookup(arg0 uint16, arg1 []byte, arg2 *big.Int) ([]byte, error) {
	return _Router.Contract.RevertLookup(&_Router.CallOpts, arg0, arg1, arg2)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0xc7d968e3.
//
// Solidity: function activateChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Router *RouterTransactor) ActivateChainPath(opts *bind.TransactOpts, _poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "activateChainPath", _poolId, _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0xc7d968e3.
//
// Solidity: function activateChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Router *RouterSession) ActivateChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ActivateChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0xc7d968e3.
//
// Solidity: function activateChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Router *RouterTransactorSession) ActivateChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ActivateChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, _poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", _poolId, _amountLD, _to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_Router *RouterSession) AddLiquidity(_poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, _poolId, _amountLD, _to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x87b21efc.
//
// Solidity: function addLiquidity(uint256 _poolId, uint256 _amountLD, address _to) returns()
func (_Router *RouterTransactorSession) AddLiquidity(_poolId *big.Int, _amountLD *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, _poolId, _amountLD, _to)
}

// CallDelta is a paid mutator transaction binding the contract method 0xfba6e280.
//
// Solidity: function callDelta(uint256 _poolId, bool _fullMode) returns()
func (_Router *RouterTransactor) CallDelta(opts *bind.TransactOpts, _poolId *big.Int, _fullMode bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "callDelta", _poolId, _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0xfba6e280.
//
// Solidity: function callDelta(uint256 _poolId, bool _fullMode) returns()
func (_Router *RouterSession) CallDelta(_poolId *big.Int, _fullMode bool) (*types.Transaction, error) {
	return _Router.Contract.CallDelta(&_Router.TransactOpts, _poolId, _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0xfba6e280.
//
// Solidity: function callDelta(uint256 _poolId, bool _fullMode) returns()
func (_Router *RouterTransactorSession) CallDelta(_poolId *big.Int, _fullMode bool) (*types.Transaction, error) {
	return _Router.Contract.CallDelta(&_Router.TransactOpts, _poolId, _fullMode)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_Router *RouterTransactor) ClearCachedSwap(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "clearCachedSwap", _srcChainId, _srcAddress, _nonce)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_Router *RouterSession) ClearCachedSwap(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ClearCachedSwap(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// ClearCachedSwap is a paid mutator transaction binding the contract method 0x403a9f7a.
//
// Solidity: function clearCachedSwap(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) returns()
func (_Router *RouterTransactorSession) ClearCachedSwap(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.Contract.ClearCachedSwap(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x16fb60f5.
//
// Solidity: function createChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Router *RouterTransactor) CreateChainPath(opts *bind.TransactOpts, _poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "createChainPath", _poolId, _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x16fb60f5.
//
// Solidity: function createChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Router *RouterSession) CreateChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Router.Contract.CreateChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x16fb60f5.
//
// Solidity: function createChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Router *RouterTransactorSession) CreateChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Router.Contract.CreateChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId, _weight)
}

// CreatePool is a paid mutator transaction binding the contract method 0x7af935a1.
//
// Solidity: function createPool(uint256 _poolId, address _token, uint8 _sharedDecimals, uint8 _localDecimals, string _name, string _symbol) returns(address)
func (_Router *RouterTransactor) CreatePool(opts *bind.TransactOpts, _poolId *big.Int, _token common.Address, _sharedDecimals uint8, _localDecimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "createPool", _poolId, _token, _sharedDecimals, _localDecimals, _name, _symbol)
}

// CreatePool is a paid mutator transaction binding the contract method 0x7af935a1.
//
// Solidity: function createPool(uint256 _poolId, address _token, uint8 _sharedDecimals, uint8 _localDecimals, string _name, string _symbol) returns(address)
func (_Router *RouterSession) CreatePool(_poolId *big.Int, _token common.Address, _sharedDecimals uint8, _localDecimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Router.Contract.CreatePool(&_Router.TransactOpts, _poolId, _token, _sharedDecimals, _localDecimals, _name, _symbol)
}

// CreatePool is a paid mutator transaction binding the contract method 0x7af935a1.
//
// Solidity: function createPool(uint256 _poolId, address _token, uint8 _sharedDecimals, uint8 _localDecimals, string _name, string _symbol) returns(address)
func (_Router *RouterTransactorSession) CreatePool(_poolId *big.Int, _token common.Address, _sharedDecimals uint8, _localDecimals uint8, _name string, _symbol string) (*types.Transaction, error) {
	return _Router.Contract.CreatePool(&_Router.TransactOpts, _poolId, _token, _sharedDecimals, _localDecimals, _name, _symbol)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xa18fa804.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _srcPoolId, (uint256,uint256) _c) returns()
func (_Router *RouterTransactor) CreditChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _srcPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "creditChainPath", _dstChainId, _dstPoolId, _srcPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xa18fa804.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _srcPoolId, (uint256,uint256) _c) returns()
func (_Router *RouterSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _srcPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Router.Contract.CreditChainPath(&_Router.TransactOpts, _dstChainId, _dstPoolId, _srcPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xa18fa804.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _srcPoolId, (uint256,uint256) _c) returns()
func (_Router *RouterTransactorSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _srcPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Router.Contract.CreditChainPath(&_Router.TransactOpts, _dstChainId, _dstPoolId, _srcPoolId, _c)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Router *RouterTransactor) InstantRedeemLocal(opts *bind.TransactOpts, _srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "instantRedeemLocal", _srcPoolId, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Router *RouterSession) InstantRedeemLocal(_srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.InstantRedeemLocal(&_Router.TransactOpts, _srcPoolId, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0xc4de93a5.
//
// Solidity: function instantRedeemLocal(uint16 _srcPoolId, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Router *RouterTransactorSession) InstantRedeemLocal(_srcPoolId uint16, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.InstantRedeemLocal(&_Router.TransactOpts, _srcPoolId, _amountLP, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactor) RedeemLocal(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemLocal", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterSession) RedeemLocal(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocal(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0x8f2e1d18.
//
// Solidity: function redeemLocal(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactorSession) RedeemLocal(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocal(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _to, _lzTxParams)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0x7f721298.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, address _to, uint256 _amountSD, uint256 _mintAmountSD) returns()
func (_Router *RouterTransactor) RedeemLocalCallback(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _to common.Address, _amountSD *big.Int, _mintAmountSD *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemLocalCallback", _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _to, _amountSD, _mintAmountSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0x7f721298.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, address _to, uint256 _amountSD, uint256 _mintAmountSD) returns()
func (_Router *RouterSession) RedeemLocalCallback(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _to common.Address, _amountSD *big.Int, _mintAmountSD *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocalCallback(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _to, _amountSD, _mintAmountSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0x7f721298.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, address _to, uint256 _amountSD, uint256 _mintAmountSD) returns()
func (_Router *RouterTransactorSession) RedeemLocalCallback(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _to common.Address, _amountSD *big.Int, _mintAmountSD *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocalCallback(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _to, _amountSD, _mintAmountSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0x0403bce5.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _amountSD, bytes _to) returns()
func (_Router *RouterTransactor) RedeemLocalCheckOnRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _amountSD *big.Int, _to []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemLocalCheckOnRemote", _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _amountSD, _to)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0x0403bce5.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _amountSD, bytes _to) returns()
func (_Router *RouterSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _amountSD *big.Int, _to []byte) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocalCheckOnRemote(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _amountSD, _to)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0x0403bce5.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _amountSD, bytes _to) returns()
func (_Router *RouterTransactorSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _amountSD *big.Int, _to []byte) (*types.Transaction, error) {
	return _Router.Contract.RedeemLocalCheckOnRemote(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _amountSD, _to)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactor) RedeemRemote(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemRemote", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterSession) RedeemRemote(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RedeemRemote(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x84d0dba3.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLP, uint256 _minAmountLD, bytes _to, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactorSession) RedeemRemote(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLP *big.Int, _minAmountLD *big.Int, _to []byte, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RedeemRemote(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLP, _minAmountLD, _to, _lzTxParams)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// RetryRevert is a paid mutator transaction binding the contract method 0x60a3b95c.
//
// Solidity: function retryRevert(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) payable returns()
func (_Router *RouterTransactor) RetryRevert(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "retryRevert", _srcChainId, _srcAddress, _nonce)
}

// RetryRevert is a paid mutator transaction binding the contract method 0x60a3b95c.
//
// Solidity: function retryRevert(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) payable returns()
func (_Router *RouterSession) RetryRevert(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RetryRevert(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// RetryRevert is a paid mutator transaction binding the contract method 0x60a3b95c.
//
// Solidity: function retryRevert(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce) payable returns()
func (_Router *RouterTransactorSession) RetryRevert(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RetryRevert(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce)
}

// RevertRedeemLocal is a paid mutator transaction binding the contract method 0x6a7982da.
//
// Solidity: function revertRedeemLocal(uint16 _dstChainId, bytes _srcAddress, uint256 _nonce, address _refundAddress, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactor) RevertRedeemLocal(opts *bind.TransactOpts, _dstChainId uint16, _srcAddress []byte, _nonce *big.Int, _refundAddress common.Address, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "revertRedeemLocal", _dstChainId, _srcAddress, _nonce, _refundAddress, _lzTxParams)
}

// RevertRedeemLocal is a paid mutator transaction binding the contract method 0x6a7982da.
//
// Solidity: function revertRedeemLocal(uint16 _dstChainId, bytes _srcAddress, uint256 _nonce, address _refundAddress, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterSession) RevertRedeemLocal(_dstChainId uint16, _srcAddress []byte, _nonce *big.Int, _refundAddress common.Address, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RevertRedeemLocal(&_Router.TransactOpts, _dstChainId, _srcAddress, _nonce, _refundAddress, _lzTxParams)
}

// RevertRedeemLocal is a paid mutator transaction binding the contract method 0x6a7982da.
//
// Solidity: function revertRedeemLocal(uint16 _dstChainId, bytes _srcAddress, uint256 _nonce, address _refundAddress, (uint256,uint256,bytes) _lzTxParams) payable returns()
func (_Router *RouterTransactorSession) RevertRedeemLocal(_dstChainId uint16, _srcAddress []byte, _nonce *big.Int, _refundAddress common.Address, _lzTxParams IStargateRouterlzTxObj) (*types.Transaction, error) {
	return _Router.Contract.RevertRedeemLocal(&_Router.TransactOpts, _dstChainId, _srcAddress, _nonce, _refundAddress, _lzTxParams)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_Router *RouterTransactor) SendCredits(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "sendCredits", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_Router *RouterSession) SendCredits(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Router.Contract.SendCredits(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// SendCredits is a paid mutator transaction binding the contract method 0x9ba3aa74.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress) payable returns()
func (_Router *RouterTransactorSession) SendCredits(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Router.Contract.SendCredits(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress)
}

// SetBridgeAndFactory is a paid mutator transaction binding the contract method 0x424c9119.
//
// Solidity: function setBridgeAndFactory(address _bridge, address _factory) returns()
func (_Router *RouterTransactor) SetBridgeAndFactory(opts *bind.TransactOpts, _bridge common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setBridgeAndFactory", _bridge, _factory)
}

// SetBridgeAndFactory is a paid mutator transaction binding the contract method 0x424c9119.
//
// Solidity: function setBridgeAndFactory(address _bridge, address _factory) returns()
func (_Router *RouterSession) SetBridgeAndFactory(_bridge common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetBridgeAndFactory(&_Router.TransactOpts, _bridge, _factory)
}

// SetBridgeAndFactory is a paid mutator transaction binding the contract method 0x424c9119.
//
// Solidity: function setBridgeAndFactory(address _bridge, address _factory) returns()
func (_Router *RouterTransactorSession) SetBridgeAndFactory(_bridge common.Address, _factory common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetBridgeAndFactory(&_Router.TransactOpts, _bridge, _factory)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0x43a30630.
//
// Solidity: function setDeltaParam(uint256 _poolId, bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Router *RouterTransactor) SetDeltaParam(opts *bind.TransactOpts, _poolId *big.Int, _batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setDeltaParam", _poolId, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0x43a30630.
//
// Solidity: function setDeltaParam(uint256 _poolId, bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Router *RouterSession) SetDeltaParam(_poolId *big.Int, _batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Router.Contract.SetDeltaParam(&_Router.TransactOpts, _poolId, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0x43a30630.
//
// Solidity: function setDeltaParam(uint256 _poolId, bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Router *RouterTransactorSession) SetDeltaParam(_poolId *big.Int, _batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Router.Contract.SetDeltaParam(&_Router.TransactOpts, _poolId, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0xc6a27624.
//
// Solidity: function setFeeLibrary(uint256 _poolId, address _feeLibraryAddr) returns()
func (_Router *RouterTransactor) SetFeeLibrary(opts *bind.TransactOpts, _poolId *big.Int, _feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setFeeLibrary", _poolId, _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0xc6a27624.
//
// Solidity: function setFeeLibrary(uint256 _poolId, address _feeLibraryAddr) returns()
func (_Router *RouterSession) SetFeeLibrary(_poolId *big.Int, _feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetFeeLibrary(&_Router.TransactOpts, _poolId, _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0xc6a27624.
//
// Solidity: function setFeeLibrary(uint256 _poolId, address _feeLibraryAddr) returns()
func (_Router *RouterTransactorSession) SetFeeLibrary(_poolId *big.Int, _feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetFeeLibrary(&_Router.TransactOpts, _poolId, _feeLibraryAddr)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 _poolId, uint256 _mintFeeBP) returns()
func (_Router *RouterTransactor) SetFees(opts *bind.TransactOpts, _poolId *big.Int, _mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setFees", _poolId, _mintFeeBP)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 _poolId, uint256 _mintFeeBP) returns()
func (_Router *RouterSession) SetFees(_poolId *big.Int, _mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetFees(&_Router.TransactOpts, _poolId, _mintFeeBP)
}

// SetFees is a paid mutator transaction binding the contract method 0x0b78f9c0.
//
// Solidity: function setFees(uint256 _poolId, uint256 _mintFeeBP) returns()
func (_Router *RouterTransactorSession) SetFees(_poolId *big.Int, _mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetFees(&_Router.TransactOpts, _poolId, _mintFeeBP)
}

// SetMintFeeOwner is a paid mutator transaction binding the contract method 0xcefbdde2.
//
// Solidity: function setMintFeeOwner(address _owner) returns()
func (_Router *RouterTransactor) SetMintFeeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setMintFeeOwner", _owner)
}

// SetMintFeeOwner is a paid mutator transaction binding the contract method 0xcefbdde2.
//
// Solidity: function setMintFeeOwner(address _owner) returns()
func (_Router *RouterSession) SetMintFeeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetMintFeeOwner(&_Router.TransactOpts, _owner)
}

// SetMintFeeOwner is a paid mutator transaction binding the contract method 0xcefbdde2.
//
// Solidity: function setMintFeeOwner(address _owner) returns()
func (_Router *RouterTransactorSession) SetMintFeeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetMintFeeOwner(&_Router.TransactOpts, _owner)
}

// SetProtocolFeeOwner is a paid mutator transaction binding the contract method 0x34aba410.
//
// Solidity: function setProtocolFeeOwner(address _owner) returns()
func (_Router *RouterTransactor) SetProtocolFeeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setProtocolFeeOwner", _owner)
}

// SetProtocolFeeOwner is a paid mutator transaction binding the contract method 0x34aba410.
//
// Solidity: function setProtocolFeeOwner(address _owner) returns()
func (_Router *RouterSession) SetProtocolFeeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeOwner(&_Router.TransactOpts, _owner)
}

// SetProtocolFeeOwner is a paid mutator transaction binding the contract method 0x34aba410.
//
// Solidity: function setProtocolFeeOwner(address _owner) returns()
func (_Router *RouterTransactorSession) SetProtocolFeeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeOwner(&_Router.TransactOpts, _owner)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0x98e391a1.
//
// Solidity: function setSwapStop(uint256 _poolId, bool _swapStop) returns()
func (_Router *RouterTransactor) SetSwapStop(opts *bind.TransactOpts, _poolId *big.Int, _swapStop bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setSwapStop", _poolId, _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0x98e391a1.
//
// Solidity: function setSwapStop(uint256 _poolId, bool _swapStop) returns()
func (_Router *RouterSession) SetSwapStop(_poolId *big.Int, _swapStop bool) (*types.Transaction, error) {
	return _Router.Contract.SetSwapStop(&_Router.TransactOpts, _poolId, _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0x98e391a1.
//
// Solidity: function setSwapStop(uint256 _poolId, bool _swapStop) returns()
func (_Router *RouterTransactorSession) SetSwapStop(_poolId *big.Int, _swapStop bool) (*types.Transaction, error) {
	return _Router.Contract.SetSwapStop(&_Router.TransactOpts, _poolId, _swapStop)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0x7b84d287.
//
// Solidity: function setWeightForChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Router *RouterTransactor) SetWeightForChainPath(opts *bind.TransactOpts, _poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setWeightForChainPath", _poolId, _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0x7b84d287.
//
// Solidity: function setWeightForChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Router *RouterSession) SetWeightForChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Router.Contract.SetWeightForChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0x7b84d287.
//
// Solidity: function setWeightForChainPath(uint256 _poolId, uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Router *RouterTransactorSession) SetWeightForChainPath(_poolId *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Router.Contract.SetWeightForChainPath(&_Router.TransactOpts, _poolId, _dstChainId, _dstPoolId, _weight)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_Router *RouterTransactor) Swap(opts *bind.TransactOpts, _dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargateRouterlzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swap", _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_Router *RouterSession) Swap(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargateRouterlzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// Swap is a paid mutator transaction binding the contract method 0x9fbf10fc.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _srcPoolId, uint256 _dstPoolId, address _refundAddress, uint256 _amountLD, uint256 _minAmountLD, (uint256,uint256,bytes) _lzTxParams, bytes _to, bytes _payload) payable returns()
func (_Router *RouterTransactorSession) Swap(_dstChainId uint16, _srcPoolId *big.Int, _dstPoolId *big.Int, _refundAddress common.Address, _amountLD *big.Int, _minAmountLD *big.Int, _lzTxParams IStargateRouterlzTxObj, _to []byte, _payload []byte) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, _dstChainId, _srcPoolId, _dstPoolId, _refundAddress, _amountLD, _minAmountLD, _lzTxParams, _to, _payload)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x2f925555.
//
// Solidity: function swapRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _dstGasForCall, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s, bytes _payload) returns()
func (_Router *RouterTransactor) SwapRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _dstGasForCall *big.Int, _to common.Address, _s PoolSwapObj, _payload []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapRemote", _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _dstGasForCall, _to, _s, _payload)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x2f925555.
//
// Solidity: function swapRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _dstGasForCall, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s, bytes _payload) returns()
func (_Router *RouterSession) SwapRemote(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _dstGasForCall *big.Int, _to common.Address, _s PoolSwapObj, _payload []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapRemote(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _dstGasForCall, _to, _s, _payload)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x2f925555.
//
// Solidity: function swapRemote(uint16 _srcChainId, bytes _srcAddress, uint256 _nonce, uint256 _srcPoolId, uint256 _dstPoolId, uint256 _dstGasForCall, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s, bytes _payload) returns()
func (_Router *RouterTransactorSession) SwapRemote(_srcChainId uint16, _srcAddress []byte, _nonce *big.Int, _srcPoolId *big.Int, _dstPoolId *big.Int, _dstGasForCall *big.Int, _to common.Address, _s PoolSwapObj, _payload []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapRemote(&_Router.TransactOpts, _srcChainId, _srcAddress, _nonce, _srcPoolId, _dstPoolId, _dstGasForCall, _to, _s, _payload)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// WithdrawMintFee is a paid mutator transaction binding the contract method 0x5500585c.
//
// Solidity: function withdrawMintFee(uint256 _poolId, address _to) returns()
func (_Router *RouterTransactor) WithdrawMintFee(opts *bind.TransactOpts, _poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawMintFee", _poolId, _to)
}

// WithdrawMintFee is a paid mutator transaction binding the contract method 0x5500585c.
//
// Solidity: function withdrawMintFee(uint256 _poolId, address _to) returns()
func (_Router *RouterSession) WithdrawMintFee(_poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawMintFee(&_Router.TransactOpts, _poolId, _to)
}

// WithdrawMintFee is a paid mutator transaction binding the contract method 0x5500585c.
//
// Solidity: function withdrawMintFee(uint256 _poolId, address _to) returns()
func (_Router *RouterTransactorSession) WithdrawMintFee(_poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawMintFee(&_Router.TransactOpts, _poolId, _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0xc8adf12d.
//
// Solidity: function withdrawProtocolFee(uint256 _poolId, address _to) returns()
func (_Router *RouterTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, _poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawProtocolFee", _poolId, _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0xc8adf12d.
//
// Solidity: function withdrawProtocolFee(uint256 _poolId, address _to) returns()
func (_Router *RouterSession) WithdrawProtocolFee(_poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawProtocolFee(&_Router.TransactOpts, _poolId, _to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0xc8adf12d.
//
// Solidity: function withdrawProtocolFee(uint256 _poolId, address _to) returns()
func (_Router *RouterTransactorSession) WithdrawProtocolFee(_poolId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Router.Contract.WithdrawProtocolFee(&_Router.TransactOpts, _poolId, _to)
}

// RouterCachedSwapSavedIterator is returned from FilterCachedSwapSaved and is used to iterate over the raw logs and unpacked data for CachedSwapSaved events raised by the Router contract.
type RouterCachedSwapSavedIterator struct {
	Event *RouterCachedSwapSaved // Event containing the contract specifics and raw log

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
func (it *RouterCachedSwapSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterCachedSwapSaved)
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
		it.Event = new(RouterCachedSwapSaved)
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
func (it *RouterCachedSwapSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterCachedSwapSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterCachedSwapSaved represents a CachedSwapSaved event raised by the Router contract.
type RouterCachedSwapSaved struct {
	ChainId    uint16
	SrcAddress []byte
	Nonce      *big.Int
	Token      common.Address
	AmountLD   *big.Int
	To         common.Address
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCachedSwapSaved is a free log retrieval operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_Router *RouterFilterer) FilterCachedSwapSaved(opts *bind.FilterOpts) (*RouterCachedSwapSavedIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "CachedSwapSaved")
	if err != nil {
		return nil, err
	}
	return &RouterCachedSwapSavedIterator{contract: _Router.contract, event: "CachedSwapSaved", logs: logs, sub: sub}, nil
}

// WatchCachedSwapSaved is a free log subscription operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_Router *RouterFilterer) WatchCachedSwapSaved(opts *bind.WatchOpts, sink chan<- *RouterCachedSwapSaved) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "CachedSwapSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterCachedSwapSaved)
				if err := _Router.contract.UnpackLog(event, "CachedSwapSaved", log); err != nil {
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

// ParseCachedSwapSaved is a log parse operation binding the contract event 0x8186389e97ff190cd5e17304ed8188a4a98a6c8add46e6df94462ac7f7e8dd34.
//
// Solidity: event CachedSwapSaved(uint16 chainId, bytes srcAddress, uint256 nonce, address token, uint256 amountLD, address to, bytes payload, bytes reason)
func (_Router *RouterFilterer) ParseCachedSwapSaved(log types.Log) (*RouterCachedSwapSaved, error) {
	event := new(RouterCachedSwapSaved)
	if err := _Router.contract.UnpackLog(event, "CachedSwapSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Router contract.
type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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
func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnershipTransferred represents a OwnershipTransferred event raised by the Router contract.
type RouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRedeemLocalCallbackIterator is returned from FilterRedeemLocalCallback and is used to iterate over the raw logs and unpacked data for RedeemLocalCallback events raised by the Router contract.
type RouterRedeemLocalCallbackIterator struct {
	Event *RouterRedeemLocalCallback // Event containing the contract specifics and raw log

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
func (it *RouterRedeemLocalCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRedeemLocalCallback)
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
		it.Event = new(RouterRedeemLocalCallback)
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
func (it *RouterRedeemLocalCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRedeemLocalCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRedeemLocalCallback represents a RedeemLocalCallback event raised by the Router contract.
type RouterRedeemLocalCallback struct {
	SrcChainId   uint16
	SrcAddress   common.Hash
	Nonce        *big.Int
	SrcPoolId    *big.Int
	DstPoolId    *big.Int
	To           common.Address
	AmountSD     *big.Int
	MintAmountSD *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocalCallback is a free log retrieval operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_Router *RouterFilterer) FilterRedeemLocalCallback(opts *bind.FilterOpts, srcAddress [][]byte, nonce []*big.Int) (*RouterRedeemLocalCallbackIterator, error) {

	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "RedeemLocalCallback", srcAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &RouterRedeemLocalCallbackIterator{contract: _Router.contract, event: "RedeemLocalCallback", logs: logs, sub: sub}, nil
}

// WatchRedeemLocalCallback is a free log subscription operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_Router *RouterFilterer) WatchRedeemLocalCallback(opts *bind.WatchOpts, sink chan<- *RouterRedeemLocalCallback, srcAddress [][]byte, nonce []*big.Int) (event.Subscription, error) {

	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "RedeemLocalCallback", srcAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRedeemLocalCallback)
				if err := _Router.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
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

// ParseRedeemLocalCallback is a log parse operation binding the contract event 0xc7379a02e530fbd0a46ea1ce6fd91987e96535798231a796bdc0e1a688a50873.
//
// Solidity: event RedeemLocalCallback(uint16 srcChainId, bytes indexed srcAddress, uint256 indexed nonce, uint256 srcPoolId, uint256 dstPoolId, address to, uint256 amountSD, uint256 mintAmountSD)
func (_Router *RouterFilterer) ParseRedeemLocalCallback(log types.Log) (*RouterRedeemLocalCallback, error) {
	event := new(RouterRedeemLocalCallback)
	if err := _Router.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRevertIterator is returned from FilterRevert and is used to iterate over the raw logs and unpacked data for Revert events raised by the Router contract.
type RouterRevertIterator struct {
	Event *RouterRevert // Event containing the contract specifics and raw log

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
func (it *RouterRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRevert)
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
		it.Event = new(RouterRevert)
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
func (it *RouterRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRevert represents a Revert event raised by the Router contract.
type RouterRevert struct {
	BridgeFunctionType uint8
	ChainId            uint16
	SrcAddress         []byte
	Nonce              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRevert is a free log retrieval operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_Router *RouterFilterer) FilterRevert(opts *bind.FilterOpts) (*RouterRevertIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Revert")
	if err != nil {
		return nil, err
	}
	return &RouterRevertIterator{contract: _Router.contract, event: "Revert", logs: logs, sub: sub}, nil
}

// WatchRevert is a free log subscription operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_Router *RouterFilterer) WatchRevert(opts *bind.WatchOpts, sink chan<- *RouterRevert) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Revert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRevert)
				if err := _Router.contract.UnpackLog(event, "Revert", log); err != nil {
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

// ParseRevert is a log parse operation binding the contract event 0xa5d2ba6de30cc2f2e91c5a29ba66b148c27826954217e2f67cb8983541da21cf.
//
// Solidity: event Revert(uint8 bridgeFunctionType, uint16 chainId, bytes srcAddress, uint256 nonce)
func (_Router *RouterFilterer) ParseRevert(log types.Log) (*RouterRevert, error) {
	event := new(RouterRevert)
	if err := _Router.contract.UnpackLog(event, "Revert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRevertRedeemLocalIterator is returned from FilterRevertRedeemLocal and is used to iterate over the raw logs and unpacked data for RevertRedeemLocal events raised by the Router contract.
type RouterRevertRedeemLocalIterator struct {
	Event *RouterRevertRedeemLocal // Event containing the contract specifics and raw log

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
func (it *RouterRevertRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRevertRedeemLocal)
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
		it.Event = new(RouterRevertRedeemLocal)
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
func (it *RouterRevertRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRevertRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRevertRedeemLocal represents a RevertRedeemLocal event raised by the Router contract.
type RouterRevertRedeemLocal struct {
	SrcChainId     uint16
	SrcPoolId      *big.Int
	DstPoolId      *big.Int
	To             []byte
	RedeemAmountSD *big.Int
	MintAmountSD   *big.Int
	Nonce          *big.Int
	SrcAddress     common.Hash
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevertRedeemLocal is a free log retrieval operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_Router *RouterFilterer) FilterRevertRedeemLocal(opts *bind.FilterOpts, nonce []*big.Int, srcAddress [][]byte) (*RouterRevertRedeemLocalIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "RevertRedeemLocal", nonceRule, srcAddressRule)
	if err != nil {
		return nil, err
	}
	return &RouterRevertRedeemLocalIterator{contract: _Router.contract, event: "RevertRedeemLocal", logs: logs, sub: sub}, nil
}

// WatchRevertRedeemLocal is a free log subscription operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_Router *RouterFilterer) WatchRevertRedeemLocal(opts *bind.WatchOpts, sink chan<- *RouterRevertRedeemLocal, nonce []*big.Int, srcAddress [][]byte) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var srcAddressRule []interface{}
	for _, srcAddressItem := range srcAddress {
		srcAddressRule = append(srcAddressRule, srcAddressItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "RevertRedeemLocal", nonceRule, srcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRevertRedeemLocal)
				if err := _Router.contract.UnpackLog(event, "RevertRedeemLocal", log); err != nil {
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

// ParseRevertRedeemLocal is a log parse operation binding the contract event 0x6ace246fa15cf1d5decabf654b1e8581a4422e0fcf4c1ed4bf83f41687caec19.
//
// Solidity: event RevertRedeemLocal(uint16 srcChainId, uint256 _srcPoolId, uint256 _dstPoolId, bytes to, uint256 redeemAmountSD, uint256 mintAmountSD, uint256 indexed nonce, bytes indexed srcAddress)
func (_Router *RouterFilterer) ParseRevertRedeemLocal(log types.Log) (*RouterRevertRedeemLocal, error) {
	event := new(RouterRevertRedeemLocal)
	if err := _Router.contract.UnpackLog(event, "RevertRedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
