package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	"github.com/pkg/errors"
)

type AllowedReq struct {
	Token       Token
	WalletAddr  common.Address
	SpenderAddr common.Address
	Retry       int
}

type AllowedRes struct {
	Allowance *big.Int
}

func (c *EtheriumClient) TokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {
	return c.tokenAllowed(ctx, req)
}
func (c *EtheriumClient) tokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {

	addr, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, NewErrTokenNotSupported(req.Token)
	}

	caller, err := erc_20.NewStorageCaller(addr, c.Cli)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	opt := &bind.CallOpts{Context: ctx}
	allowance, err := caller.Allowance(opt, req.WalletAddr, req.SpenderAddr)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &AllowedRes{Allowance: allowance}, nil
}

type ApproveReq struct {
	Token       Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
	Retry       int
}

type ApproveRes struct {
	Tx     *types.Transaction
	TxHash common.Hash
}

type TokenLimitCheckerReq struct {
	Token       Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
}

type TokenLimitCheckerRes struct {
	LimitExtended bool
	ApproveTx     *types.Transaction
}

func (c *EtheriumClient) TokenLimitChecker(ctx context.Context, req *TokenLimitCheckerReq) (*TokenLimitCheckerRes, error) {
	r := &TokenLimitCheckerRes{
		LimitExtended: false,
		ApproveTx:     nil,
	}

	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
		Token:       req.Token,
		WalletAddr:  req.Wallet.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.Amount.Cmp(allowed.Allowance) == 1 {
		tx, err := c.TokenApprove(ctx, &ApproveReq{
			Token:       req.Token,
			Wallet:      req.Wallet,
			Amount:      req.Amount,
			SpenderAddr: req.SpenderAddr,
		})
		if err != nil {
			return nil, err
		}
		r.LimitExtended = true
		r.ApproveTx = tx.Tx

		_ = c.WaitTxComplete(ctx, tx.Tx.Hash())
	}

	return r, nil
}

func (c *EtheriumClient) TokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {
	return c.tokenApprove(ctx, req)
}

func (r *ApproveReq) Validate(c *ClientConfig) error {

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	_, ok := c.TokenMap[r.Token]
	if !ok {
		return NewErrTokenNotSupported(r.Token)
	}

	if r.Amount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("zero Amount")
	}

	return nil
}

func (c *EtheriumClient) tokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {

	if err := req.Validate(c.Cfg); err != nil {
		return nil, err
	}

	addr := c.Cfg.TokenMap[req.Token]

	caller, err := erc_20.NewStorageTransactor(addr, c.Cli)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	tx, err := caller.Approve(opt, req.SpenderAddr, req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{Tx: tx}, nil
}
