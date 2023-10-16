package wise

import (
	"context"
	"fmt"
)

const (
	defaultQuoteCurrency = "TWD"
	defaultCardFeeRate   = 0.015
	defaultMilesRate     = 0.1
)

type Cost struct {
	price         *Price
	quoteCurrency string
	sourceMidRate float64
}

func NewCost(ctx context.Context, price *Price) (*Cost, error) {
	sourceMidRate, err := QueryMidRate(ctx, price.SourceCurrency, defaultQuoteCurrency)
	if err != nil {
		return nil, err
	}

	return &Cost{
		price:         price,
		quoteCurrency: defaultQuoteCurrency,
		sourceMidRate: sourceMidRate,
	}, nil
}

func (c *Cost) CardFee() float64 {
	return c.price.SourceAmount * defaultCardFeeRate
}

func (c *Cost) TotalAmount() float64 {
	return c.price.SourceAmount + c.CardFee()
}

func (c *Cost) WiseFeeRate() float64 {
	return c.price.Total / c.price.SourceAmount
}

func (c *Cost) Miles() float64 {
	return c.price.SourceAmount * c.sourceMidRate * defaultMilesRate
}

func (c *Cost) TotalFee() float64 {
	return c.CardFee() + c.price.Total
}

func (c *Cost) TotalFeeRate() float64 {
	return c.TotalFee() / c.TotalAmount()
}

func (c *Cost) MilePrice() float64 {
	return c.TotalFee() * c.sourceMidRate / c.Miles()
}

func (c *Cost) String() string {
	return fmt.Sprintf(`Add %.2f %s, pay with %.2f %s, wise fee: %.2f %s (%.2f%%), total fee: %.2f %s (%.2f%%), miles: %.2f, mile price: %.2f TWD/mile`,
		c.price.TargetAmount,
		c.price.TargetCurrency,
		c.price.SourceAmount,
		c.price.SourceCurrency,
		c.price.Total,
		c.price.SourceCurrency,
		c.WiseFeeRate()*100,
		c.TotalFee(),
		c.price.SourceCurrency,
		c.TotalFeeRate()*100.0,
		c.Miles(),
		c.MilePrice(),
	)
}
