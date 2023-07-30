package zksynclite

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) GetNetworkToken() defi.Token {
	return c.Cfg.mainToken
}

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {

	acc, err := c.AccountInfo(ctx, &AccountInfoReq{
		Addr: req.WalletAddress,
	})
	if err != nil {
		return nil, err
	}

	if acc == nil {
		return nil, errors.New("can't get balance")
	}

	m := map[v1.Token]*defi.GetBalanceRes{}

	for token, balString := range acc.Committed.Balances {
		wei, ok := big.NewInt(0).SetString(balString, 10)
		if !ok {
			return nil, errors.New("invalid value: " + balString)
		}

		_, exist := v1.Token_value[token]
		if !exist {
			continue
		}

		t := v1.Token(v1.Token_value[token])

		f, _ := defi.WEIToEther(wei).Float64()
		m[t] = &defi.GetBalanceRes{
			Req:           req,
			WEI:           wei,
			ETHER:         defi.WeiToToken(wei, req.Token),
			HumanReadable: defi.WeiToToken(wei, req.Token).String(),
			Float:         f,
		}
	}

	b, exist := m[req.Token]
	if !exist {
		return &defi.GetBalanceRes{
			Req:           req,
			WEI:           big.NewInt(0),
			ETHER:         big.NewFloat(0),
			HumanReadable: "0",
			Float:         0,
		}, nil
	}

	return b, nil
}

type AccountInfoReq struct {
	Addr common.Address
}

// "committed": boolean, // Whether block with this operation was committed in L2.
// "verified": boolean, // Whether block with this transaction was already verified.
type AccountInfoRes struct {
	Address    string `json:"address"`
	Id         int    `json:"id"`
	Depositing struct {
		Balances map[string]string `json:"balances"`
	} `json:"depositing"`
	Committed struct {
		Balances map[string]string `json:"balances"`
		//Nfts struct {
		//} `json:"nfts"`
		//MintedNfts struct {
		//} `json:"mintedNfts"`
		Nonce      int    `json:"nonce"`
		PubKeyHash string `json:"pubKeyHash"`
	} `json:"committed"`
	Verified struct {
		Balances map[string]string `json:"balances"`
		//Nfts struct {
		//} `json:"nfts"`
		//MintedNfts struct {
		//} `json:"mintedNfts"`
		Nonce      int    `json:"nonce"`
		PubKeyHash string `json:"pubKeyHash"`
	} `json:"verified"`
	AccountType interface{} `json:"accountType"`
}

func (c *Client) AccountInfo(ctx context.Context, req *AccountInfoReq) (*AccountInfoRes, error) {
	rpccli := c.ethcli.Client()
	var a AccountInfoRes
	err := rpccli.CallContext(ctx, &a, "account_info", req.Addr)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
