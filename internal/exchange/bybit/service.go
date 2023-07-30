package bybit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/cksidharthan/go-bybit"
	_ "github.com/ginarea/gobybit/bybitv5"
	"github.com/google/go-querystring/query"
	"github.com/hardstylez72/cry/internal/exchange"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

const host = "https://api.bybit.com"

type Config struct {
	ApiKey    string
	SecretKey string
	HttpCli   *http.Client
}

type Service struct {
	cfg     *Config
	httpCli *http.Client
}

func NewService(cfg *Config) *Service {

	httpCli := &http.Client{}
	if cfg.HttpCli != nil {
		httpCli = cfg.HttpCli
	}

	return &Service{
		cfg:     cfg,
		httpCli: httpCli,
	}
}

func (S *Service) GetExchangeWithdrawOptions(ctx context.Context, req *v1.GetExchangeWithdrawOptionsRequest) (*v1.GetExchangeWithdrawOptionsResponse, error) {
	panic(1)
}
func (S *Service) Withdraw(ctx context.Context, req *exchange.WithdrawRequest) (*exchange.WithdrawResponse, error) {
	panic(1)
}
func (S *Service) WaitConfirm(ctx context.Context, id exchange.WithdrawId) (*string, error) {
	panic(1)
}
func (S *Service) Ping(ctx context.Context) error {
	panic(1)
}

type GetBalanceResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			TotalEquity            string `json:"totalEquity"`
			AccountIMRate          string `json:"accountIMRate"`
			TotalMarginBalance     string `json:"totalMarginBalance"`
			TotalInitialMargin     string `json:"totalInitialMargin"`
			AccountType            string `json:"accountType"`
			TotalAvailableBalance  string `json:"totalAvailableBalance"`
			AccountMMRate          string `json:"accountMMRate"`
			TotalPerpUPL           string `json:"totalPerpUPL"`
			TotalWalletBalance     string `json:"totalWalletBalance"`
			AccountLTV             string `json:"accountLTV"`
			TotalMaintenanceMargin string `json:"totalMaintenanceMargin"`
			Coin                   []struct {
				AvailableToBorrow   string `json:"availableToBorrow"`
				Bonus               string `json:"bonus"`
				AccruedInterest     string `json:"accruedInterest"`
				AvailableToWithdraw string `json:"availableToWithdraw"`
				TotalOrderIM        string `json:"totalOrderIM"`
				Equity              string `json:"equity"`
				TotalPositionMM     string `json:"totalPositionMM"`
				UsdValue            string `json:"usdValue"`
				UnrealisedPnl       string `json:"unrealisedPnl"`
				BorrowAmount        string `json:"borrowAmount"`
				TotalPositionIM     string `json:"totalPositionIM"`
				WalletBalance       string `json:"walletBalance"`
				CumRealisedPnl      string `json:"cumRealisedPnl"`
				Coin                string `json:"coin"`
			} `json:"coin"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type GetBalanceRequest struct {
	AccountType string `url:"accountType" json:"accountType"`
	Coin        string `url:"coin,omitempty" json:"coin,omitempty"`
}

func (s *Service) GetBalance(ctx context.Context, coin string) (float64, error) {

	params := &GetBalanceRequest{
		AccountType: "FUND",
		Coin:        coin,
	}

	u := host + "/spot/v3/private/account"

	apiPath, err := url.Parse(u)
	if err != nil {
		return 0, nil
	}

	queryString, err := query.Values(params)
	if err != nil {
		return 0, nil
	}
	apiPath.RawQuery = queryString.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiPath.String(), nil)
	if err != nil {
		return 0, err
	}

	req = populateSignature(req, queryString.Encode(), s.cfg.ApiKey, s.cfg.SecretKey)

	println(req.URL.String())
	res, err := s.httpCli.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return 0, errors.New(string(b))
	}

	var response GetBalanceResponse
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(b, &response); err != nil {
		return 0, err
	}

	return 0, nil
}

func signHmac(s, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func populateSignature(req *http.Request, s, apiKey, apiSecret string) *http.Request {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)

	window := strconv.Itoa(20000)

	req.Header.Set("X-BAPI-API-KEY", apiKey)
	req.Header.Set("X-BAPI-RECV-WINDOW", window)
	req.Header.Set("X-BAPI-TIMESTAMP", now)
	req.Header.Set("X-BAPI-SIGN", signHmac(now+apiKey+window+s, apiSecret))
	return req
}
