package optimism

import (
	"context"
	"crypto_scripts/internal/defi"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {
	return c.defi.GetBalance(ctx, req)
}

func (c *Client) TxViewFn(id string) string {
	return c.defi.TxViewFn(id)
}

func (c *Client) StakeSTG(ctx context.Context, req *defi.StakeSTGReq) (*defi.StakeSTGRes, error) {
	return c.defi.StakeSTG(ctx, req)
}

func (c *Client) StargateBridgeSwap(ctx context.Context, req *defi.StargateBridgeSwapReq) (*defi.StargateBridgeSwapRes, error) {
	return c.defi.StargateBridgeSwap(ctx, req)
}
func (c *Client) GetStargateBridgeFee(ctx context.Context, req *defi.GetStargateBridgeFeeReq) (*defi.GetStargateBridgeFeeRes, error) {
	return c.defi.GetStargateBridgeFee(ctx, req)
}

func (c *Client) GetNetworkToken() defi.Token {
	return c.defi.GetNetworkToken()
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	return c.defi.Transfer(ctx, r)
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.defi.SuggestGasPrice(ctx)
}

func (c *Client) WaitTxComplete(ctx context.Context, tx common.Hash) error {
	return c.defi.WaitTxComplete(ctx, tx)
}
