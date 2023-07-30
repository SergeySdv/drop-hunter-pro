package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
)

type TokenLimitCheckerReq struct {
	Token       v1.Token
	WalletPK    string
	Amount      *big.Int
	SpenderAddr common.Address
	Gas         *defi.Gas
}

type TokenLimitCheckerRes struct {
	ApproveTx *defi.Transaction
}

func (c *Client) TokenLimitChecker(ctx context.Context, req *TokenLimitCheckerReq) (*TokenLimitCheckerRes, error) {
	r := &TokenLimitCheckerRes{
		ApproveTx: nil,
	}

	if req.Token == c.Cfg.MainToken {
		return &TokenLimitCheckerRes{
			ApproveTx: nil,
		}, nil
	}

	tx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
		Token:       req.Token,
		WalletAddr:  tx.WalletAddr,
		SpenderAddr: req.SpenderAddr,
	})
	if err != nil {
		return nil, err
	}

	if req.Amount.Cmp(allowed.Allowance) == 1 {
		tx, err := c.TokenApprove(ctx, &ApproveReq{
			Token:       req.Token,
			Wallet:      tx,
			Amount:      req.Amount,
			SpenderAddr: req.SpenderAddr,
		})
		if err != nil {
			return nil, err
		}

		if err := c.WaitTxComplete(ctx, tx.TxHash); err != nil {
			return nil, err
		}

		r.ApproveTx = c.NewTx(tx.TxHash, defi.CodeApprove)
	}

	return r, nil
}

type AllowedReq struct {
	Token       v1.Token
	WalletAddr  common.Address
	SpenderAddr common.Address
	Retry       int
}

type AllowedRes struct {
	Allowance *big.Int
}

func (c *Client) TokenAllowed(ctx context.Context, req *AllowedReq) (*AllowedRes, error) {
	addr, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, defi.NewErrTokenNotSupported(req.Token)
	}

	caller, err := erc_20.NewStorageCaller(addr, c.Provider)
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
	Token       v1.Token
	Wallet      *WalletTransactor
	Amount      *big.Int
	SpenderAddr common.Address
	Retry       int
}

func (r *ApproveReq) Validate(tm map[v1.Token]common.Address) error {

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	_, ok := tm[r.Token]
	if !ok {
		return defi.NewErrTokenNotSupported(r.Token)
	}

	if r.Amount.Cmp(big.NewInt(0)) == 0 {
		return errors.New("zero Amount")
	}

	return nil
}

type ApproveRes struct {
	TxHash common.Hash
}

func (c *Client) TokenApprove(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {

	if err := req.Validate(c.Cfg.TokenMap); err != nil {
		return nil, err
	}

	addr := c.Cfg.TokenMap[req.Token]

	abi, err := erc_20.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := abi.Pack("approve", req.SpenderAddr, req.Amount)
	if err != nil {
		return nil, err
	}

	signer, err := accounts.NewEthSignerFromRawPrivateKey(req.Wallet.PKb, c.NetworkId.Int64())
	if err != nil {
		return nil, err
	}
	w, err := accounts.NewWallet(signer, c.Provider)
	if err != nil {
		return nil, err
	}

	hash, err := w.Execute(addr, data, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "caller.Allowance")
	}

	return &ApproveRes{TxHash: hash}, nil
}
