package wise

import (
	"fmt"
)

const (
	defaultQuoteCurrency = "TWD"
	defaultCardFeeRate   = 0.015
	defaultRewardRate    = 0.1
)

type Cost struct {
	Price

	CardFeeRate float64
	RewardRate  float64
}

func NewCost(price Price) *Cost {
	return &Cost{
		Price:       price,
		CardFeeRate: defaultCardFeeRate,
		RewardRate:  defaultRewardRate,
	}
}

func (c *Cost) String() string {
	cardFee := c.SourceAmount * c.CardFeeRate
	wiseFeeRate := c.Total / c.SourceAmount
	totalFee := c.Total + cardFee
	totalFeeRate := totalFee / c.SourceAmount
	costPerMile := totalFee / (c.SourceAmount * c.RewardRate)

	s := fmt.Sprintf("Add %.2f %s", c.TargetAmount, c.TargetCurrency)
	s += fmt.Sprintf(", pay with %.2f %s", c.SourceAmount, c.SourceCurrency)
	s += fmt.Sprintf(", wise fee: %.2f %s (%.2f%%)", c.Total, c.SourceCurrency, wiseFeeRate*100)
	s += fmt.Sprintf(", total fee: %.2f %s (%.2f%%)", totalFee, c.SourceCurrency, totalFeeRate*100)
	s += fmt.Sprintf(", cost per mile: %.4f", costPerMile)
	return s
}
