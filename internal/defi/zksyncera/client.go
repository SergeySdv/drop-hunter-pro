package zksyncera

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/etherium"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/scan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/clients"
)

const (
	MainNetURL = "https://rpc.ankr.com/zksync_era" // mainnet
	TestNetURL = "https://testnet.era.zksync.dev"
)

var ZEROADDR = common.HexToAddress("0x0000000000000000000000000000000000000000")

type SyncSwap struct {
	RouterSwap         common.Address
	ClassicPoolFactory common.Address
	ClassicPool        common.Address
}

type Muteio struct {
	RouterSwap common.Address
}

type Client struct {
	EthRpc    *rpc.Client
	Provider  *clients.DefaultProvider
	NetworkId *big.Int
	Cfg       *Config
	scanner   scan.Scanner
}

type Maverick struct {
	Router common.Address
}

type SpaceFI struct {
	Router common.Address
}

type Velocore struct {
	Router common.Address
}

type IZUMI struct {
	Router common.Address
	Quoter common.Address
}

type VeSync struct {
	Router common.Address
}

type Ezkalibur struct {
	Router common.Address
}

type ZkSwap struct {
	Router common.Address
}

type Config struct {
	Weth      common.Address
	Network   v1.Network
	MainToken v1.Token
	MainNet   string
	TokenMap  map[v1.Token]common.Address
	SyncSwap  SyncSwap
	Maverick  Maverick
	Httpcli   *http.Client
	TxViewFn  func(s string) string
	networkId *big.Int
	Muteio    Muteio
	SpaceFI   SpaceFI
	Velocore  Velocore
	IZUMI     IZUMI
	VeSync    VeSync
	Ezkalibur Ezkalibur
	ZkSwap    ZkSwap
}

type ClientConfig struct {
	HttpCli     *http.Client
	RPCEndpoint string
}

func TxViewer(txId string) string {
	return "https://explorer.zksync.io/tx/" + txId //mainnet
}

func TxTestNetViewer(txId string) string {
	return "https://goerli.explorer.zksync.io/tx/" + txId
}

func NewTestNetClient(c *ClientConfig) (*Client, error) {

	syncSwap := SyncSwap{
		RouterSwap:         common.HexToAddress("0xB3b7fCbb8Db37bC6f572634299A58f51622A847e"), // staging
		ClassicPoolFactory: common.HexToAddress("0xf2FD2bc2fBC12842aAb6FbB8b1159a6a83E72006"), // staging
	}

	var TokenAddress = map[v1.Token]common.Address{
		v1.Token_ETH:  common.HexToAddress("0x20b28b1e4665fff290650586ad76e977eab90c5d"), // staging
		v1.Token_USDC: common.HexToAddress("0x0faF6df7054946141266420b43783387A78d82A9"), //staging
	}
	muteio := Muteio{
		RouterSwap: common.HexToAddress(""),
	}

	maverick := Maverick{
		Router: common.HexToAddress(""),
	}

	spaceFI := SpaceFI{
		Router: common.HexToAddress(""),
	}

	velocore := Velocore{
		Router: common.HexToAddress(""),
	}

	izumi := IZUMI{
		Router: common.HexToAddress(""),
	}

	vesync := VeSync{
		Router: common.HexToAddress(""),
	}
	ezkalibur := Ezkalibur{
		Router: common.HexToAddress(""),
	}

	ZkSwap := ZkSwap{
		Router: common.HexToAddress(""),
	}
	return newClient(
		c,
		syncSwap,
		TokenAddress,
		TxTestNetViewer,
		v1.Network_ZKSYNCERATESTNET,
		scan.NewTestNetService(),
		muteio,
		maverick,
		spaceFI,
		velocore,
		izumi,
		vesync,
		ezkalibur,
		ZkSwap,
	)
}

func NewMainNetClient(c *ClientConfig) (*Client, error) {

	syncSwap := SyncSwap{
		RouterSwap:         common.HexToAddress("0x2da10A1e27bF85cEdD8FFb1AbBe97e53391C0295"), // mainnet
		ClassicPoolFactory: common.HexToAddress("0xf2DAd89f2788a8CD54625C60b55cD3d2D0ACa7Cb"), // mainnet
		ClassicPool:        common.HexToAddress(""),
	}

	var TokenAddress = map[v1.Token]common.Address{
		v1.Token_ETH:   common.HexToAddress("0x5aea5775959fbc2557cc8789bc1bf90a239d9a91"), // mainnet
		v1.Token_USDC:  common.HexToAddress("0x3355df6d4c9c3035724fd0e3914de96a5a83aaf4"), // mainnet
		v1.Token_USDT:  common.HexToAddress("0x493257fd37edb34451f62edf8d2a0c418852ba4c"),
		v1.Token_WETH:  common.HexToAddress("0x5AEa5775959fBC2557Cc8789bC1bf90A239D9a91"),
		v1.Token_LSD:   common.HexToAddress("0x458A2E32eAbc7626187E6b75f29D7030a5202bD4"),
		v1.Token_LUSD:  common.HexToAddress("0x503234F203fC7Eb888EEC8513210612a43Cf6115"),
		v1.Token_MUTE:  common.HexToAddress("0x0e97C7a0F8B2C9885C8ac9fC6136e829CbC21d42"),
		v1.Token_MAV:   common.HexToAddress("0x787c09494Ec8Bcb24DcAf8659E7d5D69979eE508"),
		v1.Token_SPACE: common.HexToAddress("0x47260090cE5e83454d5f05A0AbbB2C953835f777"),
		v1.Token_VC:    common.HexToAddress("0x85D84c774CF8e9fF85342684b0E795Df72A24908"),
		v1.Token_IZI:   common.HexToAddress("16a9494e257703797d747540f01683952547ee5b"),
	}

	muteio := Muteio{
		RouterSwap: common.HexToAddress("0x8B791913eB07C32779a16750e3868aA8495F5964"),
	}

	maverick := Maverick{
		Router: common.HexToAddress("0x39E098A153Ad69834a9Dac32f0FCa92066aD03f4"),
	}

	spaceFI := SpaceFI{
		Router: common.HexToAddress("0xbE7D1FD1f6748bbDefC4fbaCafBb11C6Fc506d1d"),
	}

	velocore := Velocore{
		Router: common.HexToAddress("0xd999E16e68476bC749A28FC14a0c3b6d7073F50c"),
	}

	izumi := IZUMI{
		Router: common.HexToAddress("0x943ac2310D9BC703d6AB5e5e76876e212100f894"),
		Quoter: common.HexToAddress("0x30C089574551516e5F1169C32C6D429C92bf3CD7"),
	}

	vesync := VeSync{
		Router: common.HexToAddress("0x6C31035D62541ceba2Ac587ea09891d1645D6D07"),
	}

	ezkalibur := Ezkalibur{
		Router: common.HexToAddress("0x498f7bB59c61307De7dEA005877220e4406470e9"),
	}

	ZkSwap := ZkSwap{
		Router: common.HexToAddress("0x18381c0f738146Fb694DE18D1106BdE2BE040Fa4"),
	}

	return newClient(
		c,
		syncSwap,
		TokenAddress,
		TxViewer,
		v1.Network_ZKSYNCERA,
		scan.NewMainNetService(),
		muteio,
		maverick,
		spaceFI,
		velocore,
		izumi,
		vesync,
		ezkalibur,
		ZkSwap,
	)
}

func newClient(
	c *ClientConfig,
	syncSwap SyncSwap,
	tm map[v1.Token]common.Address,
	viewer func(s string) string,
	n v1.Network,
	scanner scan.Scanner,
	Muteio Muteio,
	maverick Maverick,
	spaceFI SpaceFI,
	velocore Velocore,
	izumi IZUMI,
	VeSync VeSync,
	ezkalibur Ezkalibur,
	ZkSwap ZkSwap,
) (*Client, error) {

	config := &ClientConfig{
		HttpCli: &http.Client{},
	}
	if c != nil {
		config = c
	}

	provider, err := clients.NewDefaultProvider(c.RPCEndpoint)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialOptions(context.Background(), c.RPCEndpoint, rpc.WithHTTPClient(config.HttpCli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.RPCEndpoint)
	}

	provider.Client = ethclient.NewClient(rpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	chainId, err := provider.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	ethRpc, err := rpc.DialOptions(context.Background(), etherium.MainNetURL, rpc.WithHTTPClient(config.HttpCli))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to ETH: "+c.RPCEndpoint)
	}

	return &Client{
		EthRpc:    ethRpc,
		Provider:  provider,
		NetworkId: chainId,
		scanner:   scanner,
		Cfg: &Config{
			Weth:      tm[v1.Token_WETH],
			Network:   n,
			MainToken: v1.Token_ETH,
			MainNet:   c.RPCEndpoint,
			TokenMap:  tm,
			SyncSwap:  syncSwap,
			Httpcli:   config.HttpCli,
			TxViewFn:  viewer,
			networkId: chainId,
			Muteio:    Muteio,
			Maverick:  maverick,
			SpaceFI:   spaceFI,
			Velocore:  velocore,
			IZUMI:     izumi,
			VeSync:    VeSync,
			Ezkalibur: ezkalibur,
			ZkSwap:    ZkSwap,
		},
	}, nil
}

func (c *Client) NewTx(id common.Hash, code defi.TxCode) *defi.Transaction {
	return &defi.Transaction{
		Code:    code,
		Network: c.Cfg.Network,
		Id:      id.String(),
		Url:     c.TxViewFn(id.String()),
	}
}
