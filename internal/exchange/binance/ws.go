package binance

import (
	"context"

	"github.com/adshao/go-binance/v2"
	_ "github.com/adshao/go-binance/v2"
	"github.com/hardstylez72/cry/internal/lib"
)

type Ws struct {
	cli *binance.Client
}

func NewWsClient() *Ws {
	client := binance.NewClient("", "")
	return &Ws{client}
}

func (ws *Ws) GetCoinPrice(ctx context.Context, coin string) (float64, error) {
	res, err := ws.cli.NewAveragePriceService().Symbol(coin).Do(ctx)
	if err != nil {
		return 0, err
	}

	return lib.StringToFloat(res.Price)
}
