package okex

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/okex"
	"github.com/hardstylez72/okex/requests/rest/funding"
	responses "github.com/hardstylez72/okex/responses/funding"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *Service) Currency(ctx context.Context, ccy string) (responses.GetCurrencies, error) {
	return s.cli.Rest.Funding.GetCurrency(ctx, ccy)
}
func (s *Service) Withdraw(ctx context.Context, req *exchange.WithdrawRequest) (*exchange.WithdrawResponse, error) {

	ccy, err := s.Currency(ctx, req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "s.Currency")
	}

	var fee string
	for _, c := range ccy.Currencies {
		if c.Ccy == req.Token && req.Network == c.Chain {
			fee = c.MinFee
			break
		}
	}

	am, err := lib.StringToFloat(req.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "lib.StringToFloat(am)")
	}

	feee, err := lib.StringToFloat(fee)
	if err != nil {
		return nil, errors.Wrap(err, "lib.StringToFloat(fee)")
	}

	res, err := s.cli.Rest.Funding.Withdrawal(ctx, funding.Withdrawal{
		Ccy:    req.Token,
		Chain:  req.Network,
		ToAddr: req.ToAddress,
		Pwd:    "-",
		Amt:    am - 2*feee,
		Fee:    feee,
		Dest:   okex.WithdrawalDigitalAddressDestination,
	})
	if err != nil {
		return nil, errors.Wrap(err, "cli.Rest.Funding.Withdrawal")
	}

	var id string
	if len(res.Withdrawals) > 0 {
		id = strconv.Itoa(int(res.Withdrawals[0].WdID))
	}

	return &exchange.WithdrawResponse{
		WithdrawId: id,
	}, nil
}

func (s *Service) WithdrawStatus(ctx context.Context, withdrawId string) (string, error) {
	res, err := s.cli.Rest.Funding.GetWithdrawalHistory(ctx, funding.GetWithdrawalHistory{
		WdId: withdrawId,
	})
	if err != nil {
		return "", err
	}

	for _, tx := range res.WithdrawalHistories {
		switch tx.State {
		case okex.WithdrawalSent:
			return "done", nil
		case okex.WithdrawalPending, okex.WithdrawalSending:
			return "waiting", nil
		case okex.WithdrawalFailed, okex.WithdrawalCanceled, okex.WithdrawalPendingCancel,
			4, 5, 6, 8, 9, 12:
			return "error", nil
		}
	}

	return "", errors.New("tx has unknown status")
}

func (s *Service) WaitConfirm(ctx context.Context, id exchange.WithdrawId) (*string, error) {

	maxPendings := 100
	pendings := 0

	for pendings < maxPendings {
		pendings++
		res, err := s.cli.Rest.Funding.GetWithdrawalHistory(ctx, funding.GetWithdrawalHistory{
			WdId: id,
		})
		if err != nil {
			return nil, err
		}

		for _, tx := range res.WithdrawalHistories {
			switch tx.State {
			case okex.WithdrawalSent:
				return &tx.TxID, nil
			case okex.WithdrawalPending, okex.WithdrawalSending:
			case okex.WithdrawalFailed, okex.WithdrawalCanceled, okex.WithdrawalPendingCancel,
				4, 5, 6, 8, 9, 12:
				b, _ := json.Marshal(tx)
				return nil, errors.New(string(b))
			}
		}

		time.Sleep(time.Second * 10)
	}

	return nil, errors.New("unknown binance withdraw error")
}

func (s *Service) GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error) {
	res, err := s.cli.Rest.Funding.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	tokens := make([]*v1.ExchangeWithdrawOptions, 0)

	m := map[string]*v1.ExchangeWithdrawOptions{}
	for _, r := range res.Currencies {

		_, exist := m[r.Ccy]
		if !exist {
			m[r.Ccy] = &v1.ExchangeWithdrawOptions{
				Token: r.Ccy,
				Networks: []*v1.ExchangeWithdrawNetwork{
					{
						Network:   r.Chain,
						MinAmount: r.MinWd,
						MaxAmount: r.MaxWd,
						Fee:       r.MinFee + "-" + r.MaxFee,
					},
				},
			}
		} else {
			m[r.Ccy].Networks = append(m[r.Ccy].Networks, &v1.ExchangeWithdrawNetwork{
				Network:   r.Chain,
				MinAmount: r.MinWd,
				MaxAmount: r.MaxWd,
				Fee:       r.MinFee + "-" + r.MaxFee,
			})
		}
	}

	packLen := 20
	packs := make([][]string, 0)
	pack := make([]string, packLen)
	i := 0
	j := 0
	for token, stat := range m {
		pack[i] = token
		tokens = append(tokens, stat)
		if i == packLen-1 {
			i = 0
			j++
			packs = append(packs, pack)
			pack = make([]string, packLen)
		}
		i++
	}

	mu := sync.Mutex{}
	bm := map[string]string{}
	wg := sync.WaitGroup{}
	wg.Add(len(packs))

	for i := range packs {
		p := packs[i]
		time.Sleep(time.Second / 5)
		go func(p []string) {
			defer wg.Done()
			balance, err := s.cli.Rest.Funding.GetBalance(ctx, funding.GetBalance{
				Ccy: p,
			})
			if err != nil {
				log.Log.Warn("Rest.Funding.GetBalance", zap.Error(err))
				return
			}
			for _, b := range balance.Balances {
				mu.Lock()
				bm[b.Ccy] = b.Bal
				mu.Unlock()
			}
		}(p)
	}

	wg.Wait()

	for _, token := range tokens {
		if bm[token.Token] == "" {
			continue
		}
		token.Amount = bm[token.Token]
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].Amount > tokens[i].Amount
	})

	return &v1.GetExchangeWithdrawOptionsResponse{
		Tokens: tokens,
	}, nil
}
