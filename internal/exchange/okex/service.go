package okex

import (
	"context"
	"net/http"

	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/okex"
	_ "github.com/hardstylez72/okex"
	"github.com/hardstylez72/okex/api"
	"github.com/hardstylez72/okex/requests/rest/funding"
	"github.com/hardstylez72/okex/requests/rest/subaccount"
	"github.com/pkg/errors"
)

type Service struct {
	cli *api.Client
}

type Config struct {
	ApiKey     string
	SecretKey  string
	PassPhrase string
	HttpClient *http.Client
}

type SubToMainReq struct {
	Ccy    string
	Amount float64
	SubAcc string
}

type SubsToMainReq struct {
	Ccy string
}

func NewService(ctx context.Context, c *Config) (*Service, error) {
	client, err := api.NewClient(ctx, c.ApiKey, c.SecretKey, c.PassPhrase, okex.NormalServer, c.HttpClient)
	if err != nil {
		return nil, err
	}

	s := &Service{
		cli: client,
	}
	return s, nil
}

func (s *Service) IsSubAccount(ctx context.Context) (bool, error) {

	res, err := s.cli.Rest.Account.GetConfig(ctx)
	if err != nil {
		return false, err
	}

	for _, c := range res.Configs {
		if c.MainUid != c.Uid {
			return true, err
		}
	}

	return false, err
}

func (s *Service) GetBalance(ctx context.Context, ccy string) (float64, error) {
	balance, err := s.cli.Rest.Funding.GetBalance(ctx, funding.GetBalance{
		Ccy: []string{ccy},
	})

	if err != nil {
		return -1, err
	}

	for _, b := range balance.Balances {
		if b.Ccy == ccy {
			return lib.StringToFloat(b.AvailBal)
		}
	}

	return 0, nil
}

func (s *Service) GetSubBalance(ctx context.Context, sub, ccy string) (float64, error) {
	balance, err := s.cli.Rest.SubAccount.GetBalance(ctx, subaccount.GetBalance{
		SubAcct: sub,
	})

	if err != nil {
		return -1, err
	}

	for _, b := range balance.Balances {
		if b.Ccy == ccy {
			return lib.StringToFloat(b.AvailBal)
		}
	}

	return 0, nil
}

func (s *Service) SubsToMain(ctx context.Context, req *SubsToMainReq) error {

	isSubAcc, err := s.IsSubAccount(ctx)
	if err != nil {
		return err
	}

	if isSubAcc {
		return errors.New("account must be main, not sub account")
	}

	subAccs, err := s.cli.Rest.SubAccount.ViewList(subaccount.ViewList{})
	if err != nil {
		return err
	}

	for _, sub := range subAccs.SubAccounts {

		b, err := s.GetSubBalance(ctx, sub.SubAcct, req.Ccy)
		if err != nil {
			return err
		}

		if b == 0 {
			continue
		}

		err = s.SubToMain(ctx, &SubToMainReq{
			Ccy:    req.Ccy,
			Amount: b,
			SubAcc: sub.SubAcct,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) SubToMain(ctx context.Context, req *SubToMainReq) error {
	_, err := s.cli.Rest.Funding.FundsTransfer(ctx, funding.FundsTransfer{
		Ccy:     req.Ccy,
		Amt:     req.Amount,
		SubAcct: req.SubAcc,
		Type:    okex.MasterSubAccountToAccount,
		From:    okex.FundingAccount,
		To:      okex.FundingAccount,
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) Ping(ctx context.Context) error {
	_, err := s.cli.Rest.Account.GetConfig(ctx)

	return err
}

type addr = string

func (s *Service) GetDepositAddress(ctx context.Context, ccy string) (map[addr]*DepositAddress, error) {
	bb, err := s.cli.Rest.Funding.GetDepositAddress(ctx, funding.GetDepositAddress{
		Ccy: ccy,
	})
	if err != nil {
		return nil, err
	}
	m := make(map[addr]*DepositAddress)
	for _, item := range bb.DepositAddresses {
		_, exist := m[item.Addr]
		if !exist {
			m[item.Addr] = &DepositAddress{
				Tag:      item.Tag,
				Addr:     item.Addr,
				Networks: []string{item.Chain},
			}
		} else {
			networks := append(m[item.Addr].Networks, item.Chain)
			m[item.Addr].Networks = networks
		}
	}

	return m, err
}

type DepositAddress struct {
	Tag      string
	Addr     string
	Networks []string
}
