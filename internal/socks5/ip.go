package socks5

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetIpStatRes struct {
	Ip              string `json:"ip"`
	IpNumber        string `json:"ip_number"`
	IpVersion       int    `json:"ip_version"`
	CountryName     string `json:"country_name"`
	CountryCode2    string `json:"country_code2"`
	Isp             string `json:"isp"`
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

func (p *Proxy) GetIpStat(ctx context.Context, ip string) (*GetIpStatRes, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.iplocation.net/?cmd=ip-country&ip="+ip, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequestWithContext")
	}
	res, err := p.Cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "p.Cli.Do")
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	var r GetIpStatRes

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.Wrap(err, "json.NewDecoder().Decode()")
	}
	return &r, nil
}

type GetIpRes struct {
	Success bool   `json:"success"`
	Ip      string `json:"ip"`
	Type    string `json:"type"`
}

func (p *Proxy) GetIp(ctx context.Context) (*GetIpRes, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api4.my-ip.io/ip.json", nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequestWithContext")
	}

	res, err := p.Cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "p.Cli.Do")
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var r GetIpRes

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.Wrap(err, "json.NewDecoder().Decode()")
	}
	return &r, nil
}
