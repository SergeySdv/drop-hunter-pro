package go1inch

// https://docs.1inch.io/docs/aggregation-protocol/api/swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

const (
	inchURL = "https://api.1inch.io/v5.0/"
)

type Network string

const (
	Eth         Network = "eth"
	Bsc         Network = "bsc"
	Matic       Network = "matic"
	Optimism    Network = "optimism"
	Arbitrum    Network = "arbitrum"
	GnosisChain Network = "gnosis"
	Avalanche   Network = "avalanche"
	Fantom      Network = "fantom"
	Klaytn      Network = "klaytn"
	Auror       Network = "auror"
	ZkSync      Network = "zksync"
)

func GetNetworks() []Network {
	return []Network{
		Eth,
		Bsc,
		Matic,
		Optimism,
		Arbitrum,
		GnosisChain,
		Avalanche,
		Fantom,
		Klaytn,
		Auror,
		ZkSync,
	}
}

var (
	networks = map[Network]string{
		Eth:         "1",
		Bsc:         "56",
		Matic:       "137",
		Optimism:    "10",
		Arbitrum:    "42161",
		GnosisChain: "100",
		Avalanche:   "43114",
		Fantom:      "250",
		Klaytn:      "8217",
		Auror:       "1313161554",
		ZkSync:      "324",
	}
)

type Config struct {
	HttpCli *http.Client
}

func NewClient(cfg *Config) *Client {

	httpCli := cfg.HttpCli
	if httpCli == nil {
		httpCli = &http.Client{}
	}

	c := &Client{
		Http: httpCli,
	}
	return c
}

func setQueryParam(endpoint *string, params []map[string]interface{}) {
	var first = true
	for _, param := range params {
		for i := range param {
			if first {
				*endpoint = fmt.Sprintf("%s?%s=%v", *endpoint, i, param[i])
				first = false
			} else {
				*endpoint = fmt.Sprintf("%s&%s=%v", *endpoint, i, param[i])
			}
		}
	}
}

var rl = ratelimit.New(1, ratelimit.Per(time.Second)) // per second

func (c *Client) doRequest(ctx context.Context, net Network, endpoint, method string, expRes interface{}, reqData interface{}, opts ...map[string]interface{}) (int, error) {
	n, ok := networks[net]
	if !ok {
		return 0, errors.New("invalid network")
	}
	callURL := fmt.Sprintf("%s%s%s", inchURL, n, endpoint)

	var dataReq []byte
	var err error

	rl.Take()

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, err
		}
	}

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&callURL, opts)
	}
	req, err := http.NewRequestWithContext(ctx, method, callURL, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}

	// todo: заменить
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	resp, err := c.Http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil

	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}

//fetch("https://api.1inch.io/v5.0/42161/approve/spender", {
//"headers": {
//"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
//"accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,ka;q=0.6,id;q=0.5",
//"cache-control": "max-age=0",
//"if-none-match": "W/\"38-8E1/rcQh+uMscosWnsIfAHNlcrY\"",
//"sec-ch-ua": "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"",
//"sec-ch-ua-mobile": "?0",
//"sec-ch-ua-platform": "\"macOS\"",
//"sec-fetch-dest": "document",
//"sec-fetch-mode": "navigate",
//"sec-fetch-site": "none",
//"sec-fetch-user": "?1",
//"upgrade-insecure-requests": "1",
//"cookie": "intercom-id-zgn72x6y=b9430c6f-4fe1-4a6f-aa68-2d62192669cc; intercom-session-zgn72x6y=; intercom-device-id-zgn72x6y=d0bdf117-c964-4b5e-a71b-2fef5d38a89d; __cf_bm=7JCq8HgHp9RbFR6gKFMhiUusRc2EV346XNSoogIhmcI-1686932336-0-AUuMZYx3IRCg8lwFZobxop93ddNuYTTItmxGqXHV/0q5NWCjX5BHIRzz2F7UUUka1cQqXkIg8szQ6MgHL13r8LA="
//},
//"referrerPolicy": "strict-origin-when-cross-origin",
//"body": null,
//"method": "GET"
//});
