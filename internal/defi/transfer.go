package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	"github.com/pkg/errors"
)

type TransferRes struct {
	Tx    *Transaction
	ECost *EstimatedGasCost
}

type TransferReq struct {
	Pk     string
	ToAddr common.Address
	Token  Token
	Amount *big.Int

	Gas          *Gas
	EstimateOnly bool
}

func (c *EtheriumClient) Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error) {

	if err := r.Validate(c.Cfg.TokenMap); err != nil {
		return nil, err
	}

	wallet, err := NewWalletTransactor(r.Pk)
	if err != nil {
		return nil, err
	}

	if r.Token == c.Cfg.MainToken {
		return c.TransferMainToken(ctx, &TransferMainTokenReq{
			Wallet:       wallet,
			ToAddr:       r.ToAddr,
			Amount:       r.Amount,
			Gas:          r.Gas,
			EstimateOnly: r.EstimateOnly,
		})
	}

	tokenAddr := c.Cfg.TokenMap[r.Token]

	trx, err := erc_20.NewStorageTransactor(tokenAddr, c.Cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	opt.NoSend = r.EstimateOnly
	if r.Gas.RuleSet() {
		opt.GasLimit = r.Gas.GasLimit.Uint64()
		opt.GasPrice = &r.Gas.GasPrice
	}

	tx, err := trx.Transfer(opt, r.ToAddr, r.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "Transfer")
	}

	eCost := &EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(tx.Gas()),
		GasPrice:    tx.GasPrice(),
		TotalGasWei: MinerGasLegacy(tx.GasPrice(), tx.Gas()),
	}

	return &TransferRes{
		Tx:    c.NewTx(tx.Hash(), CodeTransfer),
		ECost: eCost,
	}, nil
}

func (r *TransferReq) Validate(tm map[Token]common.Address) error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Pk == "" {
		return errors.New("empty wallet")
	}

	if r.Amount.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("Amount is 0")
	}

	_, ok := tm[r.Token]
	if !ok {
		return errors.New("unknown token: " + r.Token.String())
	}

	return nil
}

type TransferMainTokenReq struct {
	Wallet       *WalletTransactor
	ToAddr       common.Address
	Amount       *big.Int
	Gas          *Gas
	EstimateOnly bool
}

func (c *EtheriumClient) TransferMainToken(ctx context.Context, r *TransferMainTokenReq) (*TransferRes, error) {
	gasPrice, err := c.Cli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, r.Wallet.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.GetBalance(ctx, &GetBalanceReq{
		WalletAddress: r.Wallet.WalletAddr,
		Token:         c.Cfg.MainToken,
	})
	if err != nil {
		return nil, err
	}

	data := []byte{}

	gas, err := c.Cli.EstimateGas(ctx, ethereum.CallMsg{
		From:       r.Wallet.WalletAddr,
		To:         &r.ToAddr,
		Gas:        0,
		GasPrice:   gasPrice,
		GasFeeCap:  nil,
		GasTipCap:  nil,
		Value:      big.NewInt(1),
		Data:       data,
		AccessList: nil,
	})
	if err != nil {
		return nil, err
	}

	estimate := &EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(gas),
		GasPrice:    gasPrice,
		TotalGasWei: MinerGasLegacy(gasPrice, gas),
	}

	if r.EstimateOnly {
		return &TransferRes{
			Tx:    nil,
			ECost: estimate,
		}, nil
	}

	if r.Gas.RuleSet() {
		gasPrice = &r.Gas.GasPrice
		gas = r.Gas.GasLimit.Uint64()
	}

	am := r.Amount
	if am.Cmp(b.WEI) == 0 {
		am = new(big.Int).Sub(b.WEI, MinerGasLegacy(gasPrice, gas))
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       &r.ToAddr,
		Value:    am,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.Cfg.networkId), r.Wallet.PrivateKey)
	if err != nil {
		return nil, err
	}

	err = c.Cli.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}
	return &TransferRes{
		Tx:    c.NewTx(signedTx.Hash(), CodeTransfer),
		ECost: estimate,
	}, nil
}
