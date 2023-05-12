package binance

import (
	"context"
	"crypto_scripts/internal/tests"
	"testing"
)

func TestService(t *testing.T) {
	cfg := tests.GetConfig()
	c := NewService(cfg.BinanceKey, cfg.BinanceSecret, cfg.Cli)

	t.Run("assets", func(t *testing.T) {
		ctx := context.Background()
		c.GetUserAsset(ctx, "ETH")
	})

	//t.Run("networks", func(t *testing.T) {
	//	ctx := context.Background()
	//	c.GetTokenAvailableNetworks(ctx, []string{"ETH"})
	//})

}
