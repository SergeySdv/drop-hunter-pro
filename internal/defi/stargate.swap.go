package defi

import (
	"context"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var (
	ChainIdMap = map[v1.Network]uint16{
		v1.Network_Etherium:    101,
		v1.Network_BinanaceBNB: 102,
		v1.Network_AVALANCHE:   106,
		v1.Network_POLIGON:     109,
		v1.Network_ARBITRUM:    110,
		v1.Network_OPTIMISM:    111,
		//SBNameFantom:        112,
		//SBNameMetis:         151,
	}

	PoolIdMap = map[v1.Network]map[Token]int64{
		v1.Network_ARBITRUM: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//"FRAX":    7,
			v1.Token_ETH: 13,
			//"LUSD":    15,
			//"MAI":     16,
		},
		v1.Network_OPTIMISM: {
			v1.Token_USDC: 1,
			//"DAI":     3,
			//"FRAX":    7,
			v1.Token_ETH: 13,
			//"sUSD":    14,
			//"LUSD":    15,
			//"MAI":     16,
		},
		v1.Network_BinanaceBNB: {
			v1.Token_USDT: 2,
			//"BUSD":    5,
			//"USDD":    11,
			//"MAI":     16,
			//"METIS":   17,
		},
		v1.Network_Etherium: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			v1.Token_ETH:  13,
			//DAI: 3
			//FRAX: 7
			//USDD: 11
			//sUSD: 14
			//LUSD: 15
			//MAI: 16
			//METIS: 17
			//metis.USDT: 19
		},
		v1.Network_POLIGON: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//DAI: 3
			//MAI: 16
		},
		v1.Network_AVALANCHE: {
			v1.Token_USDC: 1,
			v1.Token_USDT: 2,
			//FRAX: 7
			//MAI: 16
			//metis.USDT: 19
		},
	}
)

type (
	StargateBridgeChainName string
)

const (
	typeFuncSwap    uint8                   = 1
	SBNameEthereum  StargateBridgeChainName = "Ethereum"
	SBNameBNB       StargateBridgeChainName = "BNB"
	SBNameAvalanche StargateBridgeChainName = "Avalanche"
	SBNamePolygon   StargateBridgeChainName = "Polygon"
	SBNameArbitrum  StargateBridgeChainName = "Arbitrum"
	SBNameOptimism  StargateBridgeChainName = "Optimism"
	SBNameFantom    StargateBridgeChainName = "Fantom"
	SBNameMetis     StargateBridgeChainName = "Metis"
)

type GasLimit struct {
	TotalGas *big.Int
}

type StargateBridgeSwapReq struct {
	DestChain v1.Network
	Wallet    *WalletTransactor
	Quantity  *big.Int
	FromToken Token
	ToToken   Token
	Fee       *big.Int
	Gas       *GasLimit
	Opt       *lib.RetryOpt
}

func (r *StargateBridgeSwapReq) Validate(currentChain v1.Network) error {
	if r == nil {
		return errors.New("empty input")
	}

	if r.Quantity.Cmp(big.NewInt(0)) == 0 {
		return errors.New("quantity 0")
	}

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	if r.DestChain == currentChain {
		return errors.New("invalid chain, same chain")
	}

	_, ok := ChainIdMap[r.DestChain]
	if !ok {
		return errors.New("invalid chain: " + string(r.DestChain))
	}

	_, ok = PoolIdMap[r.DestChain][r.ToToken]
	if !ok {
		return errors.New("invalid dest chain token: " + string(r.ToToken))
	}

	_, ok = PoolIdMap[currentChain][r.FromToken]
	if !ok {
		return errors.New("invalid current chain token: " + string(r.FromToken))
	}

	return nil
}

type StargateBridgeSwapRes struct {
	Tx []*types.Transaction
}

// https://stargateprotocol.gitbook.io/stargate/developers/how-to-swap
func (c *EtheriumClient) StargateBridgeSwap(ctx context.Context, req *StargateBridgeSwapReq) (*StargateBridgeSwapRes, error) {

	cliChain := c.c.Network

	r := &StargateBridgeSwapRes{
		Tx: []*types.Transaction{},
	}

	switch req.FromToken {
	case v1.Token_ETH:

		switch true {
		case req.DestChain == v1.Network_ARBITRUM && cliChain == v1.Network_OPTIMISM:
		case req.DestChain == v1.Network_OPTIMISM && cliChain == v1.Network_ARBITRUM:
		default:
			return nil, errors.New("invalid input params. check website: https://stargate.finance/transfer")
		}

		res, err := c.StargateBridgeSwapEth(ctx, &StargateBridgeSwapEthReq{
			DestChain: req.DestChain,
			Wallet:    req.Wallet,
			Quantity:  req.Quantity,
			Opt:       req.Opt,
			Gas:       req.Gas,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapEth")
		}
		r.Tx = append(r.Tx, res.Tx)
		_ = c.WaitTxComplete(ctx, res.Tx.Hash())
	case v1.Token_STG:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      req.Wallet,
			amount:      req.Quantity,
			SpenderAddr: c.c.Dict.Stargate.StargateRouterAddress,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.Tx = append(r.Tx, limitTx.ApproveTx)
		}

		res, err := c.StargateBridgeSwapSTG(ctx, &StargateBridgeSwapSTGReq{
			DestChain: req.DestChain,
			Wallet:    req.Wallet,
			Quantity:  req.Quantity,
			Opt:       req.Opt,
			Gas:       req.Gas,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapSTG")
		}
		_ = c.WaitTxComplete(ctx, res.Tx.Hash())
		r.Tx = append(r.Tx, res.Tx)
	default:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      req.Wallet,
			amount:      req.Quantity,
			SpenderAddr: c.c.Dict.Stargate.StargateRouterAddress,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.Tx = append(r.Tx, limitTx.ApproveTx)
		}

		res, err := c.StargateBridgeSwapToken(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwap")
		}
		_ = c.WaitTxComplete(ctx, res.Tx.Hash())
		r.Tx = append(r.Tx, res.Tx)
	}

	return r, nil
}
