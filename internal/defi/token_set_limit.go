package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/erc_20"
	"crypto_scripts/internal/lib"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type AllowedReq struct {
	Token       Token
	WalletAddr  common.Address
	SpenderAddr common.Address
}

type AllowedRes struct {
	Allowance *big.Int
}

func (c *EtheriumClient) TokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {
	return lib.Retry[*AllowedReq, *AllowedRes](ctx, c.tokenAllowed, req, &lib.RetryOpt{
		RetryCount: RetryMax,
	})
}
func (c *EtheriumClient) tokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {

	addr, ok := c.c.TokenMap[req.Token]
	if !ok {
		return nil, errTokenNotSupported(req.Token)
	}

	caller, err := erc_20.NewStorageCaller(addr, c.cli)
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
	amount      *big.Int
	SpenderAddr common.Address
}

type ApproveRes struct {
	Tx *types.Transaction
}

type TokenLimitCheckerReq struct {
	Token       Token
	Wallet      *WalletTransactor
	amount      *big.Int
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

	b, err := c.GetBalance(ctx, &GetBalanceReq{
		WalletAddress: req.Wallet.WalletAddr,
		Token:         req.Token,
	})
	if err != nil {
		return nil, err
	}

	if b.WEI.Cmp(req.amount) == -1 {
		return nil, errors.New("not enough ballance to approve token: " + string(req.Token))
	}

	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
		Token:       req.Token,
		WalletAddr:  req.Wallet.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.amount.Cmp(allowed.Allowance) == 1 {
		tx, err := c.TokenApprove(ctx, &ApproveReq{
			Token:       req.Token,
			Wallet:      req.Wallet,
			amount:      req.amount,
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
	return lib.Retry[*ApproveReq, *ApproveRes](ctx, c.tokenApprove, req, &lib.RetryOpt{
		RetryCount: RetryMax,
	})
}

func (r *ApproveReq) Validate(c *ClientConfig) error {

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	_, ok := c.TokenMap[r.Token]
	if !ok {
		return errTokenNotSupported(r.Token)
	}

	if r.amount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("zero amount")
	}

	return nil
}

func (c *EtheriumClient) tokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {

	if err := req.Validate(c.c); err != nil {
		return nil, err
	}

	addr := c.c.TokenMap[req.Token]

	caller, err := erc_20.NewStorageTransactor(addr, c.cli)
	if err != nil {
		return nil, errors.Wrap(err, "stg.NewStgCaller")
	}

	chainID, err := c.cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	tx, err := caller.Approve(opt, req.SpenderAddr, req.amount)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{Tx: tx}, nil
}
