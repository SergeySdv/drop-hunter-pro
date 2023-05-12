package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/erc_20"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type GetBalanceReq struct {
	WalletAddress common.Address
	Token         Token
	RetryCount    int
}

type GetBalanceRes struct {
	Req           *GetBalanceReq
	WEI           *big.Int
	ETHER         *big.Float
	HumanReadable string
	Float         float64
}

func (c *EtheriumClient) GetNetworkToken() Token {
	return c.c.MainToken
}

func (c *EtheriumClient) GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error) {
	return lib.Retry[*GetBalanceReq, *GetBalanceRes](ctx, c.getBalance, req, &lib.RetryOpt{
		RetryCount: req.RetryCount,
	})
}
func (c *EtheriumClient) getBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error) {

	fromAddress := req.WalletAddress

	if req.Token == c.c.MainToken {
		b, err := c.cli.BalanceAt(ctx, fromAddress, nil)
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

	ta, ok := c.c.TokenMap[req.Token]
	if !ok {
		return nil, errTokenNotSupported(req.Token)
	}

	s, err := erc_20.NewStorageCaller(ta, c.cli)
	if err != nil {
		return nil, err
	}

	b, err := s.BalanceOf(nil, fromAddress)
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
	if c.c.Network == v1.Network_BinanaceBNB {
		res.HumanReadable = WEIToEther(b).String()
	}

	return res, nil
}

func (c *EtheriumClient) TxViewFn(id string) string {
	return c.c.TxViewFn(id)
}
