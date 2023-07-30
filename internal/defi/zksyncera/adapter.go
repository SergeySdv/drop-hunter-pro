package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	tt "github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/orbiter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type GetBalanceReq struct {
	WalletAddress common.Address
	Token         v1.Token
}

type GetBalanceRes struct {
	Req           *GetBalanceReq
	WEI           *big.Int
	ETHER         *big.Float
	HumanReadable string
	Float         float64
}

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {

	if req.Token == c.Cfg.MainToken {
		b, err := c.Provider.GetClient().BalanceAt(ctx, req.WalletAddress, nil)
		if err != nil {
			return nil, err
		}

		f, _ := defi.WEIToEther(b).Float64()
		return &defi.GetBalanceRes{
			Req:           req,
			WEI:           b,
			ETHER:         defi.WEIToEther(b),
			HumanReadable: defi.WEIToEther(b).String(),
			Float:         f,
		}, nil
	}

	ta, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, defi.NewErrTokenNotSupported(req.Token)
	}

	token, err := erc20.NewIERC20(ta, c.Provider)
	if err != nil {
		return nil, err
	}

	b, err := token.BalanceOf(nil, req.WalletAddress)
	if err != nil {
		return nil, err
	}

	f, _ := defi.WEIToEther(b).Float64()
	res := &defi.GetBalanceRes{
		Req:           req,
		WEI:           b,
		ETHER:         defi.WeiToToken(b, req.Token),
		HumanReadable: defi.WeiToToken(b, req.Token).String(),
		Float:         f,
	}

	return res, nil
}

func (c *Client) TxViewFn(id string) string {
	return c.Cfg.TxViewFn(id)
}

func (c *Client) GetNetworkToken() v1.Token {
	return c.Cfg.MainToken
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.Provider.SuggestGasPrice(ctx)
}

func (c *Client) WaitTxComplete(ctx context.Context, tx common.Hash) error {
	res, err := c.Provider.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	if res.Status == tt.ReceiptStatusFailed {
		return errors.Wrap(defi.ErrTxStatusFailed, "invalid status")
	}

	return nil
}

type L1L2BridgeReq struct {
	Amount       *big.Int
	WalletPK     string
	EstimateOnly bool
	Gas          *defi.Gas
}

type L1L2BridgeRes struct {
	TxHash           *defi.Transaction
	EstimatedGasCost *defi.EstimatedGasCost
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {

	r := &defi.OrbiterBridgeRes{}

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

	res, err := c.Transfer(ctx, &defi.TransferReq{
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

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	if r.Token == c.Cfg.MainToken {
		return c.TransferMainToken(ctx, r)
	}
	return c.TransferToken(ctx, r)
}

func (c *Client) TransferToken(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	res := &defi.TransferRes{}

	to, supported := c.Cfg.TokenMap[r.Token]
	if !supported {
		return nil, defi.ErrTokenNotSupported
	}

	wtx, err := NewWalletTransactor(r.Pk, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	tokenAbi, err := abi.JSON(strings.NewReader(erc20.IERC20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}

	data, err := tokenAbi.Pack("transfer", r.ToAddr, r.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}
	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		to,
		big.NewInt(0),
		big.NewInt(0),
		r.Amount,
		data,
		nil, nil,
	)

	gas, err := c.Provider.EstimateGas712(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to EstimateGas: %w", err)
	}

	gasPrice, err := c.Provider.GetGasPrice()
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	nonce, err := c.Provider.NonceAt(ctx, w.GetAddress(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	prepared := &types.Transaction712{
		Nonce:      new(big.Int).SetUint64(nonce),
		GasTipCap:  big.NewInt(100_000_000),
		GasFeeCap:  gasPrice,
		Gas:        gas,
		To:         &tx.To,
		Value:      tx.Value.ToInt(),
		Data:       tx.Data,
		AccessList: nil,
		ChainID:    c.NetworkId,
		From:       &tx.From,
		Meta:       tx.Eip712Meta,
	}

	signature, err := wtx.Signer.SignTypedData(wtx.Signer.GetDomain(), prepared)
	if err != nil {
		return nil, errors.Wrap(err, "Signer.SignTypedData")
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return nil, errors.Wrap(err, "prepared.RLPValues")
	}

	res.ECost = &defi.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}

	if r.EstimateOnly {
		return res, nil
	}

	hash, err := c.Provider.SendRawTransaction(rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	res.Tx = c.NewTx(hash, defi.CodeTransfer)

	return res, nil

}
func (c *Client) TransferMainToken(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {

	res := &defi.TransferRes{}

	wtx, err := NewWalletTransactor(r.Pk, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		r.ToAddr,
		big.NewInt(0),
		big.NewInt(0),
		r.Amount,
		[]byte{},
		nil, nil,
	)

	gas, err := c.Provider.EstimateGas712(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to EstimateGas: %w", err)
	}

	gasPrice, err := c.Provider.GetGasPrice()
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	nonce, err := c.Provider.NonceAt(ctx, w.GetAddress(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	prepared := &types.Transaction712{
		Nonce:      new(big.Int).SetUint64(nonce),
		GasTipCap:  big.NewInt(100000000),
		GasFeeCap:  gasPrice,
		Gas:        gas,
		To:         &tx.To,
		Value:      tx.Value.ToInt(),
		Data:       tx.Data,
		AccessList: nil,
		ChainID:    c.NetworkId,
		From:       &tx.From,
		Meta:       tx.Eip712Meta,
	}

	signature, err := wtx.Signer.SignTypedData(wtx.Signer.GetDomain(), prepared)
	if err != nil {
		return nil, errors.Wrap(err, "Signer.SignTypedData")
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return nil, errors.Wrap(err, "prepared.RLPValues")
	}

	res.ECost = &defi.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}

	if r.EstimateOnly {
		return res, nil
	}

	hash, err := c.Provider.SendRawTransaction(rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	res.Tx = c.NewTx(hash, defi.CodeTransfer)

	return res, nil
}
