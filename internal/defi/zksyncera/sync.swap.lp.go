package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/syncswaprouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type SyncSwapLiquidityReq struct {
	Amount   *big.Int
	A        v1.Token
	B        v1.Token
	WalletPK string
	Add      bool

	EstimateOnly bool
	Gas          *defi.Gas
	debug        bool
}

type SyncSwapLiquidityRes struct {
	SwapTx           *defi.Transaction
	ApproveATxHash   *defi.Transaction
	ApproveBTxHash   *defi.Transaction
	EstimatedGasCost *defi.EstimatedGasCost
}

func (c *Client) SyncSwapLiquidity(ctx context.Context, req *SyncSwapLiquidityReq) (*SyncSwapLiquidityRes, error) {
	if req.B != v1.Token_ETH {
		return nil, errors.New("token B not supported: " + req.B.String())
	}

	if req.Add {
		return c.syncSwapLiquidityXXXETHAdd(ctx, req)
	}

	return c.syncSwapLiquidityXXXETHRemove(ctx, req)
}

func (c *Client) syncSwapLiquidityXXXETHAdd(ctx context.Context, req *SyncSwapLiquidityReq) (*SyncSwapLiquidityRes, error) {

	A, supported := c.Cfg.TokenMap[req.A]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(req.A)
	}
	ethAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")

	result := &SyncSwapLiquidityRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	// amount token A
	amOut, err := c.getAmountOut(ctx, &getAmountOutReq{
		Addr:      transactor.WalletAddr,
		FromToken: req.A,
		ToToken:   req.B,
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetSyncSwapFee")
	}

	poolAddr, err := c.GetSyncSwapPool(ctx, &GetSyncSwapPoolReq{
		A: req.A,
		B: req.B,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetSyncSwapPool")
	}

	r0, r1, err := c.syncSwapPoolRates(ctx, *poolAddr)
	if err != nil {
		return nil, errors.Wrap(err, "syncSwapPoolRates")
	}

	liquidityA := new(big.Float).Mul(r0, new(big.Float).SetInt(req.Amount))
	liquidityB := new(big.Float).Mul(r1, new(big.Float).SetInt(amOut))
	superLiquidity := new(big.Float).Add(liquidityA, liquidityB)
	averageLiquidity, _ := new(big.Float).Quo(superLiquidity, big.NewFloat(2)).Int(nil)
	minLiquidity := defi.Slippage(averageLiquidity, defi.SlippagePercent001)

	inputs := []syncswaprouter.SyncSwapRouterTokenInput{
		{
			Token:  A,
			Amount: req.Amount,
		},
		{
			Token:  ethAddr,
			Amount: amOut,
		},
	}

	syncswaprouterabi, err := syncswaprouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	l := len(transactor.WalletAddr.Bytes())

	data1 := padOrTrim(transactor.WalletAddr.Bytes(), l+12)

	h := common.BytesToHash(data1)

	if req.debug {
		println(h.String())
		// 0x0000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d43815
		println("pool: ", poolAddr.String())
		println(fmt.Sprintf("inputes %+v", inputs))
		println(fmt.Sprintf("data %+v", string(data1)))
		println(fmt.Sprintf("minLiquidity %+v", minLiquidity.String()))
		println("addr: ", ethAddr.String())
	}

	cbdata := common.Hex2Bytes("0x")
	data, err := syncswaprouterabi.Pack("addLiquidity2", *poolAddr, inputs, data1, minLiquidity, ethAddr, cbdata)
	if err != nil {
		return nil, err
	}

	if req.debug {
		println("0x" + common.Bytes2Hex(data))
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.SyncSwap.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		amOut,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}

	result.EstimatedGasCost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.SwapTx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

func (c *Client) syncSwapLiquidityXXXETHRemove(ctx context.Context, req *SyncSwapLiquidityReq) (*SyncSwapLiquidityRes, error) {

	poolAddr, err := c.GetSyncSwapPool(ctx, &GetSyncSwapPoolReq{
		A: req.A,
		B: req.B,
	})
	if err != nil {
		return nil, err
	}
	ethAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")

	result := &SyncSwapLiquidityRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	syncswaprouterabi, err := syncswaprouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "syncswaprouter.StorageMetaData.GetAbi")
	}

	ra, rb, err := c.syncSwapPoolRates(ctx, *poolAddr)
	if err != nil {
		return nil, errors.Wrap(err, "syncSwapPoolRates")
	}

	minAmountA, _ := new(big.Float).Quo(new(big.Float).SetInt(req.Amount), ra).Int(nil)
	minAmountB, _ := new(big.Float).Quo(new(big.Float).SetInt(req.Amount), rb).Int(nil)

	minAmounts := []*big.Int{
		defi.Slippage(minAmountA, defi.SlippagePercent001),
		defi.Slippage(minAmountB, defi.SlippagePercent001),
	}
	cbdata := common.Hex2Bytes("0x")
	data1 := makeDataForBurnLiquidity(transactor.WalletAddr, false)

	data, err := syncswaprouterabi.Pack("burnLiquidity", *poolAddr, req.Amount, data1, minAmounts, ethAddr, cbdata)
	if err != nil {
		return nil, err
	}

	if req.debug {
		println("0x" + common.Bytes2Hex(data))
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.SyncSwap.RouterSwap,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}

	result.EstimatedGasCost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.SwapTx = c.NewTx(hash, defi.CodeContract)

	return result, nil
}

func makeDataForBurnLiquidity(addr common.Address, wrapETH bool) []byte {
	firstLen := len(addr.Bytes())
	firstArg := padOrTrim(addr.Bytes(), firstLen+12)

	secondAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	secondLen := len(secondAddr.Bytes())
	secondArg := padOrTrim(secondAddr.Bytes(), secondLen+12)

	out := []byte{}
	out = append(out, firstArg...)
	out = append(out, secondArg...)
	if wrapETH {
		out[len(out)-1] = 2
	} else {
		out[len(out)-1] = 1
	}
	return out
}

func padOrTrim(bb []byte, size int) []byte {
	l := len(bb)
	if l == size {
		return bb
	}
	if l > size {
		return bb[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], bb)
	return tmp
}
