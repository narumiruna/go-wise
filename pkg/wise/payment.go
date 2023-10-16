package wise

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/max"
)

const (
	defaultCardFeeRate = 0.015
	defaultMilesRate   = 0.1
)

type Cost struct {
	Price   *Price  `json:"payment"`
	USDTTWD float64 `json:"usdt_twd"`
}

func NewCost(ctx context.Context, price *Price, quoteCurrency string) (*Cost, error) {
	bidAsk, err := max.QueryUSDTTWD()
	if err != nil {
		return nil, err
	}
	return &Cost{
		Price:   price,
		USDTTWD: bidAsk.Ask,
	}, nil
}

func (c *Cost) SourceAmount() float64 {
	return c.Price.SourceAmount
}

func (c *Cost) CardFee() float64 {
	return c.SourceAmount() * defaultCardFeeRate
}

func (c *Cost) TotalAmount() float64 {
	return c.SourceAmount() + c.CardFee()
}

func (c *Cost) WiseFeeRate() float64 {
	return c.Price.Total / c.Price.SourceAmount
}

func (c *Cost) Miles() float64 {
	return c.SourceAmount() * c.USDTTWD * defaultMilesRate
}

func (c *Cost) TotalFee() float64 {
	return c.CardFee() + c.Price.Total
}

func (c *Cost) TotalFeeRate() float64 {
	return c.TotalFee() / c.TotalAmount()
}

func (c *Cost) MilePrice() float64 {
	return c.TotalFee() * c.USDTTWD / c.Miles()
}

func (c *Cost) String() string {
	return fmt.Sprintf(`Add %.2f %s, pay %.2f %s, wise fee: %.2f %s (%.2f%%), total fee: %.2f %s (%.2f%%), miles: %.2f, mile price: %.2f TWD/mile`,
		c.Price.TargetAmount,
		c.Price.TargetCurrency,
		c.Price.SourceAmount,
		c.Price.SourceCurrency,
		c.Price.Total,
		c.Price.SourceCurrency,
		c.WiseFeeRate()*100,
		c.TotalFee(),
		c.Price.SourceCurrency,
		c.TotalFeeRate()*100.0,
		c.Miles(),
		c.MilePrice(),
	)
}
