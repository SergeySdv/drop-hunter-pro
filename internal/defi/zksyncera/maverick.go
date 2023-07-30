package zksyncera

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/contracts/maverickrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

// eth -> usdc
// 0xc04b8d59000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150000000000000000000000000000000000000000000000000000000064c11132000000000000000000000000000000000000000000000000001f33bb6f50acd30000000000000000000000000000000000000000000000000000000000f5d994000000000000000000000000000000000000000000000000000000000000003c5aea5775959fbc2557cc8789bc1bf90a239d9a9141c8cf74c27554a8972d3bf3d2bd4a14d8b604ab3355df6d4c9c3035724fd0e3914de96a5a83aaf400000000

// site https://explorer.zksync.io/tx/0x614d510f3ba8b394efbbb7712c5bc94e98239be5e79267ce6c45205ddddc8552 ETH -> USDC
// me https://explorer.zksync.io/tx/0x044fcd4eaf54ddd686589f216f506d9c310ed8af77d8139f2ef4bb5ef45643e5 USDC -> ETH
func (c *Client) MaverickSwap(ctx context.Context, req *defi.DefaultSwapReq) (*defi.DefaultSwapRes, error) {

	result := &defi.DefaultSwapRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	tokenLimitChecker, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		WalletPK:    req.WalletPK,
		Amount:      req.Amount,
		SpenderAddr: c.Cfg.Maverick.Router,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	result.ApproveTxHash = tokenLimitChecker.ApproveTx

	// path ETH -> USDC
	// 0x5aea5775959fbc2557cc8789bc1bf90a239d9a91 41c8cf74c27554a8972d3bf3d2bd4a14d8b604ab 3355df6d4c9c3035724fd0e3914de96a5a83aaf4
	// path USDC -> ETH
	// 0x3355df6d4c9c3035724fd0e3914de96a5a83aaf4 57681331b6cb8df134dccb4b54dc30e8fcdf0ad8 5aea5775959fbc2557cc8789bc1bf90a239d9a91
	// 0x3355df6d4c9c3035724fd0e3914de96a5a83aaf4 57681331b6cb8df134dccb4b54dc30e8fcdf0ad8 5aea5775959fbc2557cc8789bc1bf90a239d9a91
	// 0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4 41C8cf74c27554A8972d3bf3D2BD4a14D8B604AB 5AEa5775959fBC2557Cc8789bC1bf90A239D9a91
	data, err := c.makeMaverickSwapData(ctx, req.FromToken, req.ToToken, transactor.WalletAddr, req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "makeMaverickSwapData")
	}
	if req.Debug {
		println("0x" + common.Bytes2Hex(data))
	}

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		c.Cfg.Maverick.Router,
		big.NewInt(0),
		big.NewInt(0),
		value,
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

func (c *Client) makeMaverickSwapData(ctx context.Context, fromToken, toToken v1.Token, wallet common.Address, amountIn *big.Int) ([]byte, error) {

	fromTokenAddr, supported := c.Cfg.TokenMap[fromToken]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(fromToken)
	}

	toTokenAddr, supported := c.Cfg.TokenMap[toToken]
	if !supported {
		return nil, defi.NewErrTokenNotSupported(toToken)
	}

	am := defi.WeiToToken(amountIn, fromToken)

	d, err := getMaverickSwapData(ctx, fromTokenAddr, toTokenAddr, am.String(), "0.1")
	if err != nil {
		return nil, errors.Wrap(err, "getMaverickSwapData")
	}

	amOut := defi.TokenAmountFloatToWEI(d.OutputAmount, toToken)
	amOut = defi.Slippage(amOut, defi.SlippagePercent05)

	pathb := []byte{}
	pathb = append(pathb, common.HexToAddress(d.Path[0]).Bytes()...)
	pathb = append(pathb, common.HexToAddress(d.Path[1]).Bytes()...)
	pathb = append(pathb, common.HexToAddress(d.Path[2]).Bytes()...)
	//path := common.BytesToHash(pathb)

	// (USDC -> ETH) recipient = ZERO
	// unwrapWETH9 + am 0, recipient - wallet
	// ETH -> LUSD refundETH
	recipient := wallet
	if toToken == v1.Token_ETH {
		recipient = ZEROADDR
	}

	params := maverickrouter.IRouterExactInputParams{
		Path:             pathb,
		Recipient:        recipient,
		Deadline:         new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix()),
		AmountIn:         amountIn,
		AmountOutMinimum: amOut,
	}

	constractabi, err := maverickrouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	p1, err := constractabi.Pack("exactInput", params)
	if err != nil {
		return nil, err
	}

	var p2 []byte
	if toToken == v1.Token_ETH {
		p2, err = constractabi.Pack("unwrapWETH9", big.NewInt(0), wallet)
		if err != nil {
			return nil, err
		}
	}

	if fromToken == v1.Token_ETH {
		p2, err = constractabi.Pack("refundETH")
		if err != nil {
			return nil, err
		}
	}

	if len(p2) == 0 {
		return constractabi.Pack("multicall", [][]byte{p1})
	}

	return constractabi.Pack("multicall", [][]byte{p1, p2})
}

func getMaverickSwapData(ctx context.Context, fromToken, toToken common.Address, amountIn string, slippage string) (*getMaverickSwapDataRes, error) {
	cli := &http.Client{}
	zksyncChainId := "324"
	base := strings.Join([]string{
		"https://api.mav.xyz/api/swapRoutes?chainId=", zksyncChainId,
		"&inTokenAddress=", fromToken.String(),
		"&outTokenAddress=", toToken.String(),
		"&amount=", amountIn,
		"&gasPrice=0.25&slippage=", slippage,
	}, "")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base, nil)
	if err != nil {
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(b))
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result getMaverickSwapDataRes
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	if result.Message != "Path successfully found." {
		return nil, errors.New(result.Message)
	}

	return &result, nil
}

type getMaverickSwapDataRes struct {
	Path         [3]string `json:"path"`
	OutputAmount float64   `json:"outputAmount"`
	Message      string    `json:"message"`
}
