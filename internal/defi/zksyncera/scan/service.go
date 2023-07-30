package scan

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	TestNetUrl = "https://zksync2-testnet-explorer.zksync.dev"
	MainNetUrl = "https://zksync2-mainnet-explorer.zksync.io"
)

type Service struct {
	url     string
	httpCli *http.Client
}

func newService(url string) *Service {
	return &Service{
		url:     url,
		httpCli: &http.Client{},
	}
}

func NewTestNetService() *Service {
	return newService(TestNetUrl)
}

func NewMainNetService() *Service {
	return newService(MainNetUrl)
}

type Tx struct {
	TransactionHash string `json:"transactionHash"`
	Status          string `json:"status"`
}

func (s *Service) GetTx(ctx context.Context, id string) (*Tx, error) {
	url := s.url + `/transaction/` + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.httpCli.Do(req)
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
	var response Tx
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type TxStatus = string

const (
	TxStatusVerified   TxStatus = "verified"
	TxStatusProcessing TxStatus = "included"
)

func (s *Service) WaitComplete(ctx context.Context, txId string, interval time.Duration) (*TxStatus, error) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		tx, err := s.GetTx(ctx, txId)
		if err == nil {
			if tx.Status == TxStatusVerified {
				return &tx.Status, nil
			}
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			tmp := TxStatusProcessing
			return &tmp, nil
		}
	}
}
