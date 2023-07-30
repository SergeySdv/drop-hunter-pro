package uniclient

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"go.uber.org/zap"
)

func GetAllBalance(ctx context.Context, c *BaseClientConfig, token v1.Token, wallet common.Address) lib.Map[v1.Network, *big.Int] {
	networks := []v1.Network{
		v1.Network_OPTIMISM,
		v1.Network_AVALANCHE,
		v1.Network_ARBITRUM,
		v1.Network_POLIGON,
		v1.Network_POLIGON,
		v1.Network_BinanaceBNB,
	}

	m := lib.NewMap[v1.Network, *big.Int]()
	wg := &sync.WaitGroup{}
	wg.Add(len(networks))
	for _, network := range networks {
		go func(network v1.Network) {
			c, err := NewBaseClient(network, c)
			if err != nil {
				log.Log.Warn("GetAllBalance", zap.Error(err))
				return
			}
			b, err := c.GetBalance(ctx, &defi.GetBalanceReq{
				WalletAddress: wallet,
				Token:         token,
			})
			if err != nil {
				log.Log.Warn("GetBalance", zap.Error(err))
				return
			}
			m.Set(network, b.WEI)
		}(network)
	}

	wg.Wait()

	return *m
}
