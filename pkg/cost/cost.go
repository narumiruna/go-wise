package cost

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

const (
	defaultQuoteCurrency = "TWD"
	defaultCardFeeRate   = 0.015
	defaultRewardRate    = 0.1
)

type Cost struct {
	*wise.Price

	quoteCurrency string
	sourceRate    float64
	cardFeeRate   float64
	rewardRate    float64

	cardFee      float64
	totalAmount  float64
	wiseFeeRate  float64
	miles        float64
	totalFee     float64
	totalFeeRate float64
	milePrice    float64
}

func NewCost(ctx context.Context, source string, amount float64, target string) (*Cost, error) {
	c := wise.NewClient()

	prices, err := c.NewPriceService().QueryPrice(ctx, source, amount, target)
	if err != nil {
		return nil, err
	}

	price, err := wise.FindPrice(prices, wise.PayInMethodVisaCredit, wise.PayOutMethodBalance)
	if err != nil {
		return nil, err
	}

	rate, err := c.NewRatesService().QueryRate(ctx, price.SourceCurrency, defaultQuoteCurrency)
	if err != nil {
		return nil, err
	}

	cost := &Cost{
		Price:         price,
		quoteCurrency: defaultQuoteCurrency,
		sourceRate:    rate.Value,
		cardFeeRate:   defaultCardFeeRate,
		rewardRate:    defaultRewardRate,
	}
	cost.Initialize()
	return cost, nil
}

func (c *Cost) Initialize() *Cost {
	c.cardFee = c.SourceAmount * c.cardFeeRate
	c.totalAmount = c.SourceAmount + c.cardFee
	c.wiseFeeRate = c.Total / c.SourceAmount
	c.miles = c.SourceAmount * c.sourceRate * c.rewardRate
	c.totalFee = c.Total + c.cardFee
	c.totalFeeRate = c.totalFee / c.SourceAmount
	c.milePrice = c.totalFee * c.sourceRate / c.miles
	return c
}

func (c *Cost) String() string {
	s := fmt.Sprintf("Add %.2f %s", c.TargetAmount, c.TargetCurrency)
	s += fmt.Sprintf(", pay with %.2f %s", c.SourceAmount, c.SourceCurrency)
	s += fmt.Sprintf(", wise fee: %.2f %s (%.2f%%)", c.Total, c.SourceCurrency, c.wiseFeeRate*100)
	s += fmt.Sprintf(", total fee: %.2f %s (%.2f%%)", c.totalFee, c.SourceCurrency, c.totalFeeRate*100)
	s += fmt.Sprintf(", miles: %.2f (%.4f %s/mile)", c.miles, c.milePrice, c.quoteCurrency)
	return s
}
