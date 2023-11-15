package wise

import (
	"fmt"
)

const (
	rewardRate  = 0.1
	cardFeeRate = 0.015
)

type Cost struct {
	Price

	RewardRate   float64
	CardFee      float64
	CardFeeRate  float64
	WiseFee      float64
	WiseFeeRate  float64
	TotalFee     float64
	TotalFeeRate float64
	CostPerMile  float64
}

func NewCost(p Price) *Cost {
	cardFee := p.SourceAmount * cardFeeRate
	fee := p.Total + cardFee

	return &Cost{
		Price:        p,
		RewardRate:   rewardRate,
		CardFee:      cardFee,
		CardFeeRate:  cardFeeRate,
		WiseFee:      p.Total,
		WiseFeeRate:  p.Total / p.SourceAmount,
		TotalFee:     fee,
		TotalFeeRate: fee / p.SourceAmount,
		CostPerMile:  fee / (p.SourceAmount * rewardRate),
	}
}

func (c *Cost) String() string {
	s := fmt.Sprintf("Add %.2f %s", c.TargetAmount, c.TargetCurrency)
	s += fmt.Sprintf(", pay with %.2f %s", c.SourceAmount, c.SourceCurrency)
	s += fmt.Sprintf(", wise fee: %.2f %s (%.2f%%)", c.Total, c.SourceCurrency, c.WiseFeeRate*100)
	s += fmt.Sprintf(", total fee: %.2f %s (%.2f%%)", c.TotalFee, c.SourceCurrency, c.TotalFeeRate*100)
	s += fmt.Sprintf(", cost per mile: %.4f", c.CostPerMile)
	return s
}
