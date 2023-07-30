package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type GetBalanceReq struct {
	WalletAddress common.Address
	Token         Token
}

type GetBalanceRes struct {
	Req           *GetBalanceReq
	WEI           *big.Int
	ETHER         *big.Float
	HumanReadable string
	Float         float64
}

func (c *EtheriumClient) GetNetworkToken() Token {
	return c.Cfg.MainToken
}

func (c *EtheriumClient) GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error) {
	return c.getBalance(ctx, req)
}
func (c *EtheriumClient) getBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error) {

	fromAddress := req.WalletAddress

	if req.Token == c.Cfg.MainToken {
		b, err := c.Cli.BalanceAt(ctx, fromAddress, nil)
		if err != nil {
			return nil, err
		}

		f, _ := WEIToEther(b).Float64()
		return &GetBalanceRes{
			Req:           req,
			WEI:           b,
			ETHER:         WEIToEther(b),
			HumanReadable: WEIToEther(b).String(),
			Float:         f,
		}, nil
	}

	ta, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, NewErrTokenNotSupported(req.Token)
	}

	s, err := erc_20.NewStorageCaller(ta, c.Cli)
	if err != nil {
		return nil, err
	}

	opt := &bind.CallOpts{
		Context: ctx,
	}

	b, err := s.BalanceOf(opt, fromAddress)
	if err != nil {
		return nil, err
	}

	f, _ := WEIToEther(b).Float64()
	res := &GetBalanceRes{
		Req:           req,
		WEI:           b,
		ETHER:         WeiToToken(b, req.Token),
		HumanReadable: WeiToToken(b, req.Token).String(),
		Float:         f,
	}
	if c.Cfg.Network == v1.Network_BinanaceBNB {
		res.HumanReadable = WEIToEther(b).String()
	}

	return res, nil
}

func (c *EtheriumClient) TxViewFn(id string) string {
	return c.Cfg.TxViewFn(id)
}
