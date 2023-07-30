package zksyncera

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/contracts/ethtoken"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

// транзакция через офф бридж https://explorer.zksync.io/tx/0x8e6f82d93aac1af142d3e3fd6e58d54b6df294109b5f3b364ba7710f2eea7c99
// транзакция через софт https://explorer.zksync.io/tx/0xf471fea71aa35a1fc0716b40a41aaf286d186df1b5220ce13d16fd4a4b53591b
func (c *Client) BridgeToEthereumNetwork(ctx context.Context, req *L1L2BridgeReq) (*L1L2BridgeRes, error) {
	wtx, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "NewWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.Signer, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	ethTokenAbi, err := abi.JSON(strings.NewReader(ethtoken.IEthTokenMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}

	data, err := ethTokenAbi.Pack("withdraw", w.GetAddress())
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}
	tx := utils.CreateFunctionCallTransaction(
		w.GetAddress(),
		utils.L2EthTokenAddress,
		big.NewInt(0),
		big.NewInt(0),
		req.Amount,
		data,
		nil, nil,
	)

	var gas, gasPrice *big.Int
	if req.Gas.RuleSet() {
		gas = &req.Gas.GasLimit
		gasPrice = &req.Gas.GasPrice
		gasPrice = defi.ResolveGasPriceZksync(&req.Gas.MaxGas, gas, gasPrice)
	} else {
		gas, err = c.Provider.EstimateGas712(tx)
		if err != nil {
			return nil, fmt.Errorf("failed to EstimateGas: %w", err)
		}
		gasPrice, err = c.Provider.GetGasPrice()
		if err != nil {
			return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
		}
	}

	nonce, err := c.Provider.NonceAt(ctx, w.GetAddress(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	prepared := &types.Transaction712{
		Nonce:      new(big.Int).SetUint64(nonce),
		GasTipCap:  big.NewInt(100_000_000), // TODO: Estimate correct one
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

	result := &L1L2BridgeRes{}
	result.EstimatedGasCost = &defi.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	return &L1L2BridgeRes{
		TxHash: c.NewTx(hash, defi.CodeContract),
	}, nil
}

// сайт https://etherscan.io/tx/0x8d05dfa50a8be054b5c27d50f36e4076a1f8eda1185fd23e12eb365a71bf50a3
// код
func (c *Client) BridgeFromEthereumNetwork(ctx context.Context, req *L1L2BridgeReq) (*L1L2BridgeRes, error) {

	chainID, err := c.Provider.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(req.WalletPK), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	w, err := accounts.NewWallet(es, c.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	provider, err := w.CreateEthereumProvider(c.EthRpc)
	if err != nil {
		return nil, errors.Wrap(err, "CreateEthereumProvider")
	}

	// Perform deposit
	tx, err := provider.Deposit(
		utils.CreateETH(),
		req.Amount,
		w.GetAddress(),
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "provider.Deposit")
	}

	return &L1L2BridgeRes{
		TxHash: c.NewTx(tx.Hash(), defi.CodeContract),
	}, nil
}

func (c *Client) OfficialBridgeWait(ctx context.Context, tx common.Hash, pk string) error {

	chainID, err := c.Provider.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(pk), chainID.Int64())
	if err != nil {
		return err
	}

	w, err := accounts.NewWallet(es, c.Provider)
	if err != nil {
		return errors.Wrap(err, "zksync2.NewWallet")
	}

	provider, err := w.CreateEthereumProvider(c.EthRpc)
	if err != nil {
		return errors.Wrap(err, "CreateEthereumProvider")
	}

	l1Receipt, err := provider.GetClient().TransactionReceipt(ctx, tx)
	if err != nil {
		return err
	}

	// Get deposit transaction hash on L2 network
	l2Hash, err := provider.GetL2HashFromPriorityOp(l1Receipt)
	if err != nil {
		return err
	}

	// Wait for deposit transaction to be finalized on L2 network (5-7 minutes)
	_, err = provider.WaitMined(ctx, l2Hash)
	if err != nil {
		return err
	}
	return nil
}
