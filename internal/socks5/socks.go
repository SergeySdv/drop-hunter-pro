package socks5

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/proxy"
)

type Config struct {
	Host     string
	Login    string
	Password string
	Disable  bool
}

type Proxy struct {
	Config *Config
	Stat   *GetIpStatRes
	Cli    *http.Client
}

const mask = "ip:port:login:password"

var ErrParseError = errors.New("parse error")

// NewSock5ProxyString
// example "168.91.123.48:5346:meqdlqu:q6mnwsgrwef2"
// mask "ip:port:login:password"
func NewSock5ProxyString(s string) (*Proxy, error) {
	c, err := ParseProxy(s)
	if err != nil {
		return nil, errors.Wrap(ErrParseError, err.Error())
	}

	return NewSock5Proxy(c)
}

func (c Config) Validate() error {
	if c.Disable {
		return nil
	}
	ss := strings.Split(c.Host, ":")
	if len(ss) != 2 {
		return errors.New("invalid host format. " + c.Host + " mask XX.XX.XX.XX:XX")
	}

	if ss[0] == "" || ss[1] == "" {
		return errors.New("invalid host format. " + c.Host + " mask XX.XX.XX.XX:XX")
	}

	return nil
}

func NewSock5Proxy(c *Config) (*Proxy, error) {

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if c.Disable {
		return &Proxy{
			Config: c,
			Cli: &http.Client{
				Transport: http.DefaultTransport,
				Timeout:   time.Second * 30,
			},
		}, nil
	}

	auth := proxy.Auth{
		User:     c.Login,
		Password: c.Password,
	}
	dialer, err := proxy.SOCKS5("tcp", c.Host, &auth, proxy.Direct)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		Dial: dialer.Dial,
	}
	p := &Proxy{
		Config: c,
		Cli: &http.Client{
			Transport: tr,
			Timeout:   time.Second * 30,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	ip, err := p.GetIp(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "GetIp()")
	}

	if ip.Ip != strings.Split(c.Host, ":")[0] {
		return nil, errors.New("proxy ip and real ip does not match. please use private proxy")
	}

	stat, err := p.GetIpStat(ctx, strings.Split(c.Host, ":")[0])
	if err != nil {
		return nil, errors.Wrap(err, "GetIpStat")
	}
	p.Stat = stat

	return p, nil
}

func ParseProxy(s string) (*Config, error) {

	if s == "" {
		return &Config{
			Disable: true,
		}, nil
	}

	str := strings.Split(s, ":")
	if len(str) != 4 {
		return nil, errors.New("invalid proxy format: " + s + " format: " + mask)
	}

	if str[0] == "" || str[1] == "" {
		return nil, errors.New("invalid format. " + " format: " + mask)
	}

	return &Config{
		Host:     str[0] + ":" + str[1],
		Login:    str[2],
		Password: str[3],
		Disable:  false,
	}, nil
}
