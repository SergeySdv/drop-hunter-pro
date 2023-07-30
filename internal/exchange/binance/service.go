package binance

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/pkg/errors"
)

const (
	// (0:Email Sent,1:Cancelled 2:Awaiting Approval 3:Rejected 4:Processing 5:Failure 6:Completed)
	StatusSend             = 0
	StatusCanceled         = 1
	StatusAwaitingApproval = 2
	StatusRejected         = 3
	StatusProcessing       = 4
	StatusFail             = 5
	StatusCompleted        = 6
)

type Service struct {
	cli *binance.Client
}

func NewService(key, secret string, cli *http.Client) *Service {
	client := binance.NewClient(key, secret)

	client.HTTPClient = cli
	s := &Service{
		cli: client,
	}

	return s
}

func (s *Service) GetUserAsset(ctx context.Context, token string) (float64, error) {
	return lib.Retry[string, float64](ctx, s.getUserAsset, token, &lib.RetryOpt{
		RetryCount: 10,
	})
}

func (s *Service) GetAllUserAssets(ctx context.Context) (map[string]string, error) {
	res, err := s.cli.NewGetUserAsset().Do(ctx)
	if err != nil {
		return nil, err
	}

	m := map[string]string{}
	for _, r := range res {
		m[r.Asset] = r.Free
	}

	return m, err
}

func (s *Service) getUserAsset(ctx context.Context, token string) (float64, error) {
	res, err := s.cli.NewGetUserAsset().Asset(token).Do(ctx)
	if err != nil {
		return 0, err
	}

	for _, r := range res {
		if r.Asset == token {
			float, err := strconv.ParseFloat(r.Free, 10)
			if err != nil {
				return 0, err
			}
			return float, nil
		}
	}

	return 0, err
}

func (s *Service) GetUserAssets(ctx context.Context, tokens []string) (map[string]float64, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(tokens))
	mu := sync.Mutex{}
	out := map[string]float64{}

	errs := make([]error, len(tokens))
	for i := range tokens {
		go func(i int) {
			defer wg.Done()
			amount, err := s.GetUserAsset(ctx, tokens[i])

			if err != nil {
				errs[i] = err
				return
			}

			mu.Lock()
			out[tokens[i]] = amount
			mu.Unlock()
		}(i)

	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

type (
	WithdrawReq struct {
		ToAddress string
		Amount    string
		Network   string
		Coin      string
	}
	WithdrawRes struct {
		Req             *WithdrawReq
		WithDrawOrderId string
		Id              string
	}

	WithdrawsRes struct {
		Res *WithdrawRes
		Err error
	}
)

func (s *Service) Withdraw(ctx context.Context, r *WithdrawReq) (_ *WithdrawRes, err error) {

	id := uuid.New().String()

	res, err := s.cli.NewCreateWithdrawService().Address(r.ToAddress).Amount(r.Amount).Network(r.Network).Coin(r.Coin).WithdrawOrderID(id).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "NewCreateWithdrawService")
	}
	return &WithdrawRes{Id: res.ID, WithDrawOrderId: id, Req: r}, nil
}

func (s *Service) Withdraws(ctx context.Context, r []WithdrawReq) []WithdrawsRes {
	wg := sync.WaitGroup{}
	wg.Add(len(r))

	out := make([]WithdrawsRes, len(r))

	for i := range r {

		go func(i int) {
			defer wg.Done()
			out[i].Res, out[i].Err = s.Withdraw(ctx, &r[i])
		}(i)
	}

	wg.Wait()

	return out
}

type PendingFn func(id string, status string, err error)

func (s *Service) WithdrawStatus(ctx context.Context, withdrawId string) (string, error) {
	transactions, err := s.cli.NewListWithdrawsService().WithdrawOrderId(withdrawId).Do(ctx)
	if err != nil {
		return "", err
	}

	for _, tx := range transactions {
		switch tx.Status {
		case StatusCompleted:
			return "done", nil
		case StatusAwaitingApproval, StatusProcessing:
			return "waiting", nil
		case StatusFail, StatusRejected, StatusCanceled:
			return "error", nil
		case StatusSend:
			return "waiting", nil
		}
	}

	return "", errors.New("unknown tx status")
}

func (s *Service) WithdrawsPending(ctx context.Context, withdrawOrderId string, interval time.Duration) (*string, error) {

	maxPendings := 100
	pendings := 0

	for pendings < maxPendings {
		pendings++
		transactions, err := s.cli.NewListWithdrawsService().WithdrawOrderId(withdrawOrderId).Do(ctx)
		if err != nil {
			return nil, err
		}

		for _, tx := range transactions {
			switch tx.Status {
			case StatusCompleted:
				return &tx.TxID, nil
			case StatusAwaitingApproval, StatusProcessing:
			case StatusFail, StatusRejected, StatusCanceled:
				b, _ := json.Marshal(tx)
				return nil, errors.New(string(b))
			case StatusSend:
				continue
			}
		}

		time.Sleep(interval)
	}

	return nil, errors.New("unknown binance withdraw error")
}

func (s *Service) GetTokenAvailableNetworks(ctx context.Context, tokens []string) (map[string][]string, error) {
	return lib.Retry[[]string, map[string][]string](ctx, s.getTokenAvailableNetworks, tokens, &lib.RetryOpt{
		RetryCount: 5,
	})
}
func (s *Service) getTokenAvailableNetworks(ctx context.Context, tokens []string) (map[string][]string, error) {

	m := make(map[string][]string)

	res, err := s.cli.NewGetAllCoinsInfoService().Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "NewGetAllCoinsInfoService")
	}

	for _, token := range tokens {
		for _, r := range res {
			if r.Coin == token {
				networks := []string{}
				for _, network := range r.NetworkList {
					networks = append(networks, network.Network)
				}
				m[token] = networks
			}
		}
	}

	return m, nil
}

type WithdrawOption struct {
	Token    string
	Networks []WithdrawTokenOption
}

type WithdrawTokenOption struct {
	Network   string
	MinAmount string
	MaxAmount string
	Fee       string
}

func (s *Service) GetWithDrawOptions(ctx context.Context) ([]WithdrawOption, error) {

	m := make([]WithdrawOption, 0)

	res, err := s.cli.NewGetAllCoinsInfoService().Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "NewGetAllCoinsInfoService")
	}

	for _, o := range res {

		networks := make([]WithdrawTokenOption, 0)

		for _, n := range o.NetworkList {

			networks = append(networks, WithdrawTokenOption{
				Network:   n.Network,
				MinAmount: n.WithdrawMin,
				MaxAmount: n.WithdrawMax,
				Fee:       n.WithdrawFee,
			})
		}

		m = append(m, WithdrawOption{
			Token:    o.Coin,
			Networks: networks,
		})
	}

	return m, nil
}

func (a *Service) GetDepositAddr(ctx context.Context, network, coin string) (string, error) {

	res, err := a.cli.NewGetDepositAddressService().Network(network).Coin(coin).Do(ctx)
	if err != nil {
		return "", err
	}

	return res.Address, nil
}
