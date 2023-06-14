package wise

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-visa-fx-rates/pkg/visa"
)

const (
	defaultQuoteCurrency = "TWD"
	defaultCardFeeRate   = 0.015
	defaultMilesRate     = 0.1
)

type Payment struct {
	SourceAmount   float64 `json:"sourceAmount"`
	SourceCurrency string  `json:"sourceCurrency"`
	TargetAmount   float64 `json:"targetAmount"`
	TargetCurrency string  `json:"targetCurrency"`
}

type Cost struct {
	Payment       Payment `json:"payment"`
	QuoteCurrency string  `json:"quoteCurrency"`

	sourceQuoteAsk float64
	targetQuoteBid float64
}

func NewCost(ctx context.Context, payment Payment, quoteCurrency string) (*Cost, error) {
	client := visa.NewRestClient()
	sourceQuoteAsk, err := client.AskPrice(ctx, payment.SourceCurrency, quoteCurrency)
	if err != nil {
		return nil, err
	}
	targetQuoteBid, err := client.BidPrice(ctx, payment.TargetCurrency, quoteCurrency)
	if err != nil {
		return nil, err
	}

	return &Cost{
		Payment:        payment,
		QuoteCurrency:  defaultQuoteCurrency,
		sourceQuoteAsk: sourceQuoteAsk,
		targetQuoteBid: targetQuoteBid,
	}, nil
}

func (c *Cost) Amount() float64 {
	return c.Payment.SourceAmount * c.sourceQuoteAsk
}

func (c *Cost) CardFee() float64 {
	return c.Amount() * defaultCardFeeRate
}

func (c *Cost) TotalAmount() float64 {
	return c.Amount() + c.CardFee()
}

func (c *Cost) Miles() float64 {
	return c.Amount() * defaultMilesRate
}

func (c *Cost) WiseFee() float64 {
	return c.Amount() - c.Payment.TargetAmount*c.targetQuoteBid
}

func (c *Cost) TotalFee() float64 {
	return c.CardFee() + c.WiseFee()
}

func (c *Cost) TotalFeeRate() float64 {
	return c.TotalFee() / c.TotalAmount()
}

func (c *Cost) MilePrice() float64 {
	return c.TotalFee() / c.Miles()
}

func (c *Cost) String() string {
	return fmt.Sprintf(`Add %.2f %s, pay %.2f %s (%.2f %s), total fees: %.2f %s (%.2f%%), miles: %.2f, mile price: %.2f %s/mile`,
		c.Payment.TargetAmount,
		c.Payment.TargetCurrency,
		c.Payment.SourceAmount,
		c.Payment.SourceCurrency,
		c.Amount(),
		c.QuoteCurrency,
		c.TotalFee(),
		c.QuoteCurrency,
		c.TotalFeeRate()*100.0,
		c.Miles(),
		c.MilePrice(),
		c.QuoteCurrency,
	)
}
