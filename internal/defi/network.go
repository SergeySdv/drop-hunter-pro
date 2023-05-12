package defi

import (
	"context"
	"math/big"
)

func (c *EtheriumClient) GetNetworkId(ctx context.Context) (*big.Int, error) {
	id, err := c.cli.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	c.c.networkId = id
	return id, nil
}
func (c *EtheriumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.cli.SuggestGasPrice(ctx)
}
