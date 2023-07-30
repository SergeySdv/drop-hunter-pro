package defi

import (
	"context"
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
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
		v1.Network_GOERLIETH:   154,
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
	typeFuncSwap uint8 = 1
)

type GasLimit struct {
	TotalGas *big.Int
}

type StargateBridgeSwapReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	FromToken    Token
	ToToken      Token
	Gas          *Gas
	EstimateOnly bool
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
	Tx        *Transaction
	Allowance *Transaction
	ECost     *EstimatedGasCost
}

// https://stargateprotocol.gitbook.io/stargate/developers/how-to-swap
func (c *EtheriumClient) StargateBridgeSwap(ctx context.Context, req *StargateBridgeSwapReq) (*StargateBridgeSwapRes, error) {

	r := &StargateBridgeSwapRes{}

	switch req.FromToken {
	case v1.Token_ETH:
		res, err := c.StargateBridgeSwapEth(ctx, &StargateBridgeSwapEthReq{
			DestChain:    req.DestChain,
			Wallet:       req.Wallet,
			Quantity:     req.Quantity,
			Gas:          req.Gas,
			EstimateOnly: req.EstimateOnly,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapEth")
		}
		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract)
		}
		r.ECost = res.ECost
	case v1.Token_STG:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      req.Wallet,
			Amount:      req.Quantity,
			SpenderAddr: c.Cfg.TokenMap[v1.Token_STG],
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.Allowance = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove)
		}

		res, err := c.StargateBridgeSwapSTG(ctx, &StargateBridgeSwapSTGReq{
			DestChain:    req.DestChain,
			Wallet:       req.Wallet,
			Quantity:     req.Quantity,
			Gas:          req.Gas,
			EstimateOnly: req.EstimateOnly,
		})
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwapSTG")
		}

		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract)
		}
		r.ECost = res.ECost
	default:
		limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
			Token:       req.FromToken,
			Wallet:      req.Wallet,
			Amount:      req.Quantity,
			SpenderAddr: c.Cfg.Dict.Stargate.StargateRouterAddress,
		})
		if err != nil {
			return nil, errors.Wrap(err, "TokenLimitChecker")
		}
		if limitTx.LimitExtended {
			r.Allowance = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove)
		}

		res, err := c.StargateBridgeSwapToken(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "stargateBridgeSwap")
		}
		if res.Tx != nil {
			r.Tx = c.NewTx(res.Tx.Hash(), CodeContract)
		}
		r.ECost = res.ECost
	}

	return r, nil
}
