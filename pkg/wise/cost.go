package wise

import (
	"context"
	"fmt"
)

const (
	defaultQuoteCurrency = "TWD"
	defaultCardFeeRate   = 0.015
	defaultRewardRate    = 0.1
)

type Cost struct {
	*Price

	quoteCurrency string
	sourceRate    float64
	cardFeeRate   float64
	rewardRate    float64
}

func (c *RestClient) NewCost(ctx context.Context, source string, amount float64, target string) (*Cost, error) {
	price, err := c.QueryPrice(ctx, source, amount, target)
	if err != nil {
		return nil, err
	}

	rate, err := c.QueryRate(ctx, price.SourceCurrency, defaultQuoteCurrency)
	if err != nil {
		return nil, err
	}

	return &Cost{
		Price:         price,
		quoteCurrency: defaultQuoteCurrency,
		sourceRate:    rate,
		cardFeeRate:   defaultCardFeeRate,
		rewardRate:    defaultRewardRate,
	}, nil
}

func (c *Cost) CardFee() float64 {
	return c.SourceAmount * c.cardFeeRate
}

func (c *Cost) TotalAmount() float64 {
	return c.SourceAmount + c.CardFee()
}

func (c *Cost) WiseFeeRate() float64 {
	return c.Total / c.SourceAmount
}

func (c *Cost) Miles() float64 {
	return c.SourceAmount * c.sourceRate * c.rewardRate
}

func (c *Cost) TotalFee() float64 {
	return c.CardFee() + c.Total
}

func (c *Cost) TotalFeeRate() float64 {
	return c.TotalFee() / c.TotalAmount()
}

func (c *Cost) MilePrice() float64 {
	return c.TotalFee() * c.sourceRate / c.Miles()
}

func (c *Cost) String() string {
	s := fmt.Sprintf("Add %.2f %s", c.TargetAmount, c.TargetCurrency)
	s += fmt.Sprintf(", pay with %.2f %s", c.SourceAmount, c.SourceCurrency)
	s += fmt.Sprintf(", wise fee: %.2f %s (%.2f%%)", c.Total, c.SourceCurrency, c.WiseFeeRate()*100)
	s += fmt.Sprintf(", total fee: %.2f %s (%.2f%%)", c.TotalFee(), c.SourceCurrency, c.TotalFeeRate()*100)
	s += fmt.Sprintf(", miles: %.2f (%.2f %s/mile)", c.Miles(), c.MilePrice(), c.quoteCurrency)
	return s
}
