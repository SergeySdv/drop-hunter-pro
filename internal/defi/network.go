package defi

import (
	"context"
	"math/big"
)

func (c *EtheriumClient) GetNetworkId(ctx context.Context) (*big.Int, error) {
	id, err := c.Cli.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	c.Cfg.networkId = id
	return id, nil
}
func (c *EtheriumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.Cli.SuggestGasPrice(ctx)
}
