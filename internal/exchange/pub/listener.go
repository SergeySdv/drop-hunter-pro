package pub

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/exchange/binance"
	"github.com/hardstylez72/cry/internal/log"
)

type Listener struct {
	cli *binance.Ws
}

var p = Pairs{
	AVAX:  14,
	BNB:   250,
	ETH:   1900,
	MATIC: 0.7,
}

type Pairs struct {
	AVAX  float64
	BNB   float64
	ETH   float64
	MATIC float64
}

func Price() *Pairs {
	return &p
}

func NewIListener() *Listener {
	return &Listener{
		cli: binance.NewWsClient(),
	}
}

func (l *Listener) Listen(ctx context.Context) {

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {

			d, err := l.GetData()
			if err != nil {
				log.Log.Error("pub.Listen.GetData", err)
			} else {
				p = *d
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}()

}

func (l *Listener) GetData() (*Pairs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	bnb, err := l.cli.GetCoinPrice(ctx, "BNBUSDT")
	if err != nil {
		return nil, err
	}

	eth, err := l.cli.GetCoinPrice(ctx, "ETHUSDT")
	if err != nil {
		return nil, err
	}

	avax, err := l.cli.GetCoinPrice(ctx, "AVAXUSDT")
	if err != nil {
		return nil, err
	}

	matic, err := l.cli.GetCoinPrice(ctx, "MATICUSDT")
	if err != nil {
		return nil, err
	}

	return &Pairs{
		AVAX:  avax,
		BNB:   bnb,
		ETH:   eth,
		MATIC: matic,
	}, nil
}
