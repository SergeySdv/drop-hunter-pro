package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/scan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Token = v1.Token

type Network = v1.Network

type Networker interface {
	TxViewFn(id string) string
	GetNetworkId() *big.Int
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	WaitTxComplete(ctx context.Context, tx common.Hash) error
	Balancer
}

type Balancer interface {
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	GetNetworkToken() Token
}

type SwapRes struct {
	Tx        *Transaction
	Allowance *Transaction
	ECost     *EstimatedGasCost
}

type StargateSwapper interface {
	Networker
	StargateBridgeSwap(ctx context.Context, req *StargateBridgeSwapReq) (*StargateBridgeSwapRes, error)
	GetStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error)
}
type TestNetworkBridgeSwapper interface {
	Networker
	TestNetBridgeSwap(ctx context.Context, req *TestNetBridgeSwapReq) (*TestNetBridgeSwapRes, error)
}

type OrbiterSwapper interface {
	Transfer
	OrbiterBridge(ctx context.Context, req *OrbiterBridgeReq) (*OrbiterBridgeRes, error)
}

type Gas struct {
	Network       v1.Network
	GasMultiplier float64
	MaxGas        big.Int

	GasBeforeMultiplier big.Int
	GasLimit            big.Int
	GasPrice            big.Int
	TotalGas            big.Int
}

func (g *Gas) RuleSet() bool {
	return g != nil
}

type SyncSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
}

type DefaultSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
	Debug        bool
}

type DefaultSwapRes struct {
	SwapTx           *Transaction
	ApproveTxHash    *Transaction
	EstimatedGasCost *EstimatedGasCost
}

type SyncSwapRes struct {
	SwapTx           *Transaction
	ApproveTxHash    *Transaction
	EstimatedGasCost *EstimatedGasCost
}

type TxCode = string

const (
	CodeApprove  TxCode = "approve"
	CodeContract TxCode = "contract"
	CodeTransfer TxCode = "transfer"
)

type Transaction struct {
	Code    TxCode
	Network v1.Network
	Id      string
	Url     string
}

func (c *EtheriumClient) NewTx(id common.Hash, code TxCode) *Transaction {
	return &Transaction{
		Code:    code,
		Network: c.Cfg.Network,
		Id:      id.String(),
		Url:     c.TxViewFn(id.String()),
	}
}

type EstimatedGasCost struct {
	GasLimit    *big.Int
	GasPrice    *big.Int
	TotalGasWei *big.Int
}

type SyncSwapper interface {
	Networker
	SyncSwap(ctx context.Context, req *SyncSwapReq) (*SyncSwapRes, error)
	WaitSwapFinish(ctx context.Context, txId common.Hash) (*scan.TxStatus, error)
}

type ZkSyncOfficialWithdrawalEtherium interface {
	Networker
}

type Transfer interface {
	Networker
	Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error)
}

type WETHReq struct {
	Amount *big.Int

	Wrap bool

	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
}

type WETHRes struct {
	Tx    *Transaction
	ECost *EstimatedGasCost
}

type WETH interface {
	Networker
	SwapWETH(ctx context.Context, req *WETHReq) (*WETHRes, error)
}
