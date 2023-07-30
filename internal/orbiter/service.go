package orbiter

import (
	_ "embed"
	"encoding/json"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

//go:embed chain.json
var chainsJson []byte

//go:embed maker-1.json
var maker1 []byte

////go:embed chain.json
//var maker2 []byte
//
////go:embed chain.json
//var maker3 []byte
//
////go:embed chain.json
//var maker4 []byte

type Service struct {
	chainMap *lib.BiMap[v1.Network, string]
	makerMap map[string]map[string]MakerPair
}

func NewService() (*Service, error) {

	supportedTokens.Set(v1.Token_ETH)
	supportedNetworks.Set(v1.Network_ZKSYNCLITE)
	supportedNetworks.Set(v1.Network_OPTIMISM)
	supportedNetworks.Set(v1.Network_POLIGON)
	supportedNetworks.Set(v1.Network_ARBITRUM)
	supportedNetworks.Set(v1.Network_Etherium)
	supportedNetworks.Set(v1.Network_ZKSYNCERA)

	chains := make([]Chain, 0)
	if err := json.Unmarshal(chainsJson, &chains); err != nil {
		return nil, err
	}

	makerMap := map[string]map[string]MakerPair{}
	if err := json.Unmarshal(maker1, &makerMap); err != nil {
		return nil, err
	}

	chainMap := MakeChainMap(chains)

	return &Service{
		chainMap: chainMap,
		makerMap: makerMap,
	}, nil
}

func (s *Service) SwapOptions(from, to v1.Network, a, b v1.Token) (*MakerPair, bool) {

	_, supported := s.GetToken(a.String())
	if !supported {
		return nil, false
	}

	_, supported = s.GetToken(b.String())
	if !supported {
		return nil, false
	}

	fromId, supported := s.GetNetwork(from)
	if !supported {
		return nil, false
	}

	toId, supported := s.GetNetwork(to)
	if !supported {
		return nil, false
	}
	key := fromId + "-" + toId
	coinPairMap, supported := s.makerMap[key]
	if !supported {
		return nil, false
	}
	coinPair := a.String() + "-" + b.String()
	opt, supported := coinPairMap[coinPair]
	if !supported {
		return nil, false
	}
	return &opt, true
}
func (s *Service) GetNetwork(n v1.Network) (string, bool) {
	fromNetwork, exist := s.chainMap.Get(n)
	if !exist {
		return "", false
	}

	if !supportedNetworks.Get(n) {
		return "", false
	}

	return fromNetwork, true
}
func (s *Service) GetNetworks() []v1.Network {
	return supportedNetworks.Keys()
}
func (s *Service) GetToken(t string) (v1.Token, bool) {
	token, ok := tokenFromSting(t)
	if !ok {
		return v1.Token_BNB, false
	}
	supported := supportedTokens.Get(token)
	if !supported {
		return v1.Token_BNB, false
	}
	return token, true
}
func (s *Service) GetToTokens(from, to v1.Network, t v1.Token) []v1.Token {

	out := make([]v1.Token, 0)

	if from == v1.Network_ZKSYNCLITE {
		return out
	}

	fromNetwork, exist := s.GetNetwork(from)
	if !exist {
		return out
	}

	toNetwork, exist := s.GetNetwork(to)
	if !exist {
		return out
	}

	key := fromNetwork + "-" + toNetwork

	tm, exist := s.makerMap[key]
	if !exist {
		return out
	}

	for pair := range tm {
		sub := strings.Split(pair, "-")
		if len(sub) != 2 {
			continue
		}

		token, supported := tokenFromSting(sub[0])
		if !supported {
			continue
		}

		if token != t {
			continue
		}

		tokenTo, supported := tokenFromSting(sub[1])
		if !supported {
			continue
		}

		out = append(out, tokenTo)
	}

	return out
}
func (s *Service) GetFromTokens(from, to v1.Network) []v1.Token {

	out := make([]v1.Token, 0)

	if from == v1.Network_ZKSYNCLITE {
		return out
	}

	fromNetwork, exist := s.chainMap.Get(from)
	if !exist {
		return out
	}

	toNetwork, exist := s.chainMap.Get(to)
	if !exist {
		return out
	}

	key := fromNetwork + "-" + toNetwork

	tm, exist := s.makerMap[key]
	if !exist {
		return out
	}

	for pair := range tm {
		sub := strings.Split(pair, "-")
		if len(sub) != 2 {
			continue
		}

		token, supported := tokenFromSting(sub[0])
		if !supported {
			continue
		}

		out = append(out, token)
	}

	return out
}

type Token struct {
	MakerPair
}

func MakeChainMap(in []Chain) *lib.BiMap[v1.Network, string] {
	out := lib.NewBiMap[v1.Network, string]()

	for _, c := range in {
		switch c.ChainId {
		case "1":
			out.Set(v1.Network_Etherium, c.InternalId)
		case "42161":
			out.Set(v1.Network_ARBITRUM, c.InternalId)
		case "zksync":
			out.Set(v1.Network_ZKSYNCLITE, c.InternalId)
		case "137":
			out.Set(v1.Network_POLIGON, c.InternalId)
		case "10":
			out.Set(v1.Network_OPTIMISM, c.InternalId)
		case "324":
			out.Set(v1.Network_ZKSYNCERA, c.InternalId)
		case "56":
			out.Set(v1.Network_BinanaceBNB, c.InternalId)
		}
	}

	return out
}

type MakeTxReq struct {
	FromNetwork v1.Network
	ToNetwork   v1.Network
	FromToken   v1.Token
	ToToken     v1.Token
	Amount      *big.Int
}

type Chain struct {
	Api struct {
		Url string `json:"url"`
		Key string `json:"key"`
	} `json:"api"`
	ChainId        string `json:"chainId"`
	NetworkId      string `json:"networkId"`
	InternalId     string `json:"internalId"`
	Name           string `json:"name"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
	} `json:"nativeCurrency"`
	Rpc       []string `json:"rpc"`
	Watch     []string `json:"watch"`
	Contracts []string `json:"contracts"`
	Tokens    []struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
	} `json:"tokens"`
	XvmList []string `json:"xvmList"`
	InfoURL string   `json:"infoURL"`
}

type MakerPair struct {
	MakerAddress string `json:"makerAddress"`
	Sender       string `json:"sender"`
	//MaxPrice     StringJson `json:"maxPrice"`
	MaxPrice   float64    `json:"maxPrice"`
	MinPrice   float64    `json:"minPrice"`
	TradingFee float64    `json:"tradingFee"`
	GasFee     StringJson `json:"gasFee"`
	StartTime  int        `json:"startTime"`
	EndTime    int64      `json:"endTime"`
}

type StringJson string

func (s *StringJson) UnmarshalJSON(in []byte) error {

	str := string(in)

	*s = StringJson(str)

	return nil
}

type Tx struct {
	Value4DigitCode   string
	MakerReceiverAddr string
	MakerSenderAddr   string
}

func (s *Service) MakeTx(req *MakeTxReq) (*Tx, error) {
	opt, ok := s.SwapOptions(req.FromNetwork, req.ToNetwork, req.FromToken, req.ToToken)
	if !ok {
		return nil, errors.New("swap is not possible")
	}

	am, _ := WeiToToken(req.Amount, req.FromToken).Float64()

	if am <= opt.MinPrice {
		return nil, errors.New("amount to transfer is less than minimum: " + lib.FloatToString(opt.MinPrice))
	}

	if am >= opt.MaxPrice {
		return nil, errors.New("amount to transfer is more than maximum: " + lib.FloatToString(opt.MaxPrice))
	}

	code, ok := s.chainMap.Get(req.ToNetwork)
	if !ok {
		return nil, errors.New("swap is not possible")
	}

	c, err := strconv.Atoi(code)
	if err != nil {
		return nil, errors.New("swap is not possible")
	}

	tmp := 9000 + c

	return &Tx{
		Value4DigitCode:   strconv.Itoa(tmp),
		MakerReceiverAddr: opt.MakerAddress,
		MakerSenderAddr:   opt.Sender,
	}, nil
}

func WeiToToken(wei *big.Int, token v1.Token) *big.Float {
	switch token {
	case v1.Token_USDT, v1.Token_USDC:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether*1e-12))
	default:
		return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	}
}

func WrapValueWei(in *big.Int, code string) *big.Int {

	str := in.String()
	n := str
	if len(str) < len(code) {
		n = code
	} else {
		n = n[:len(n)-len(code)] + code
	}

	out, ok := new(big.Int).SetString(n, 10)
	if ok {
		return out
	}

	return out
}
