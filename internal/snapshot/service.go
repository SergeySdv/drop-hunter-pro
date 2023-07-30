package snapshot

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Service struct {
	cli *http.Client
	c   *Config
}

type Config struct {
	Host string
}

func NewService(c *Config) *Service {
	return &Service{
		cli: &http.Client{},
		c:   c,
	}
}

type ActiveProposalsReq struct {
	ProviderRPC string `json:"provider_rpc"`
	Space       string `json:"space"`
	Pk          string `json:"pk"`
}

func (s *Service) ActiveProposals(ctx context.Context, req *ActiveProposalsReq) ([]Proposal, error) {

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, s.c.Host+"/proposals/active", bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := s.cli.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.Wrap(err, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ser []Proposal
	if err := json.Unmarshal(body, &ser); err != nil {
		return nil, err
	}

	println(string(body))

	return ser, nil
}

type Proposal struct {
	Id         string   `json:"id"`
	Discussion string   `json:"discussion"`
	Choices    []string `json:"choices"`

	Network string `json:"network"`
	Type    string `json:"type"`
	Symbol  string `json:"symbol"`
	Privacy string `json:"privacy"`
	Space   struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"space"`
	ScoresState string    `json:"scores_state"`
	Scores      []float64 `json:"scores"`
	Error       any       `json:"error"`
}

type VoteReq struct {
	Type        string `json:"type"`
	ProposalId  string `json:"proposal_id"`
	ProviderRPC string `json:"provider_rpc" json:"provider_rpc"`
	Space       string `json:"space" json:"space"`
	Pk          string `json:"pk" json:"pk"`
}

func (s *Service) Vote(ctx context.Context, req *VoteReq) error {

	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, s.c.Host+"/proposal/vote", bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := s.cli.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}
