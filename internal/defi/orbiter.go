package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/orbiter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type OrbiterBridgeReq struct {
	OrbiterService *orbiter.Service

	FromNetwork v1.Network
	ToNetwork   v1.Network
	FromToken   v1.Token
	ToToken     v1.Token
	Amount      *big.Int

	WalletPk string

	Gas          *Gas
	EstimateOnly bool
}

type OrbiterBridgeRes struct {
	Tx    *Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) OrbiterBridge(ctx context.Context, req *OrbiterBridgeReq) (*OrbiterBridgeRes, error) {

	r := &OrbiterBridgeRes{}

	opt, err := req.OrbiterService.MakeTx(&orbiter.MakeTxReq{
		FromNetwork: req.FromNetwork,
		ToNetwork:   req.ToNetwork,
		FromToken:   req.FromToken,
		ToToken:     req.ToToken,
		Amount:      req.Amount,
	})

	if err != nil {
		return nil, err
	}
	req.Amount = orbiter.WrapValueWei(req.Amount, opt.Value4DigitCode)

	res, err := c.Transfer(ctx, &TransferReq{
		Pk:           req.WalletPk,
		ToAddr:       common.HexToAddress(opt.MakerReceiverAddr),
		Token:        req.FromToken,
		Amount:       req.Amount,
		Gas:          req.Gas,
		EstimateOnly: req.EstimateOnly,
	})
	if err != nil {
		return nil, err
	}
	r.Tx = res.Tx
	r.ECost = res.ECost

	return r, nil
}
