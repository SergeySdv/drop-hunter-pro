package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/erc_20"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type TransferRes struct {
	Tx *types.Transaction
}

type TransferReq struct {
	Wallet *WalletTransactor
	ToAddr common.Address
	Token  Token
	Amount *big.Int
}

func (r *TransferReq) Validate(tm map[Token]common.Address) error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	if r.Amount.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("amount is 0")
	}

	_, ok := tm[r.Token]
	if !ok {
		return errors.New("unknown token: " + r.Token.String())
	}

	return nil
}

func (c *EtheriumClient) Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error) {

	if err := r.Validate(c.c.TokenMap); err != nil {
		return nil, err
	}

	tokenAddr := c.c.TokenMap[r.Token]

	trx, err := erc_20.NewStorageTransactor(tokenAddr, c.cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(r.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	tx, err := trx.Transfer(opt, r.ToAddr, r.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "Transfer")
	}

	return &TransferRes{Tx: tx}, nil
}
