package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	muteiorouter "github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/muteio"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type MuteIOSwapReq struct {
	Amount    *big.Int
	FromToken v1.Token
	ToToken   v1.Token

	WalletPK     string
	EstimateOnly bool
	Gas          *defi.Gas
}

type MuteIOSwapRes struct {
	Tx        *defi.Transaction
	Allowance *defi.Transaction
	ECost     *defi.EstimatedGasCost
}

func (c *Client) MuteIOSwap(ctx context.Context, req *MuteIOSwapReq) (*MuteIOSwapRes, error) {
	if req.FromToken == v1.Token_ETH || req.FromToken == v1.Token_WETH {
		return c.muteIoSwapFromEth(ctx, req)
	}

	res, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.Muteio.RouterSwap,
	})
	if err != nil {
		return nil, err
	}

	r, err := c.muteIoSwapToEth(ctx, req)
	if err != nil {
		return nil, err
	}

	if res.ApproveTx != nil {
		r.Allowance = res.ApproveTx
	}
	return r, nil
}

// off https://explorer.zksync.io/tx/0x74dc1805388938923ffbfdbf0b021656f7bc1aca19a6508a64ffccf2632c7c94
// my
func (c *Client) muteIoSwapToEth(ctx context.Context, req *MuteIOSwapReq) (*MuteIOSwapRes, error) {

	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}
	//t, _ := muteiorouter.NewStorageTransactor()
	//t.SwapExactTokensForETHSupportingFeeOnTransferTokens()
	muteiorouterabi, err := muteiorouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	amountOutMin, _, _, err := c.MuteIoPairFee(ctx, req)
	if err != nil {
		return nil, err
	}

	min := req.Amount
	path := []common.Address{c.Cfg.TokenMap[req.FromToken], c.Cfg.Weth}
	to := w.GetAddress()
	deadline := new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
	stableMap := []bool{false, false}

	data, err := muteiorouterabi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", min, amountOutMin, path, to, deadline, stableMap)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		c.Cfg.Muteio.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &MuteIOSwapRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

// off https://explorer.zksync.io/tx/0xf5aa782e37711f8ebdfe19608ce24d5d0219103264d39f20d04738e677756c7d
// my https://explorer.zksync.io/tx/0x5f86de07362b4de327a59f6ae866ba98bb4ddd65e73d19cfb1baeca5fc702292
func (c *Client) muteIoSwapFromEth(ctx context.Context, req *MuteIOSwapReq) (*MuteIOSwapRes, error) {
	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	muteiorouterabi, err := muteiorouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	amountOutMin, _, fee, err := c.MuteIoPairFee(ctx, req)
	if err != nil {
		return nil, err
	}

	path := []common.Address{c.Cfg.Weth, c.Cfg.TokenMap[req.ToToken]}
	to := w.GetAddress()
	deadline := new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
	stableMap := []bool{false, false}

	data, err := muteiorouterabi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline, stableMap)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	value = defi.BigIntSum(value, fee)

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		c.Cfg.Muteio.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &MuteIOSwapRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

func (c *Client) MuteIoPairFee(ctx context.Context, req *MuteIOSwapReq) (AmountOut *big.Int, Stable bool, Fee *big.Int, err error) {
	caller, err := muteiorouter.NewStorageCaller(c.Cfg.Muteio.RouterSwap, c.Provider)
	if err != nil {
		return nil, false, nil, err
	}
	opt := &bind.CallOpts{}
	opt.Context = ctx

	amOut, err := caller.GetAmountOut(opt, req.Amount, c.Cfg.TokenMap[req.FromToken], c.Cfg.TokenMap[req.ToToken])
	if err != nil {
		return nil, false, nil, err
	}
	return defi.Slippage(amOut.AmountOut, defi.SlippagePercent01), amOut.Stable, amOut.Fee, nil
}
