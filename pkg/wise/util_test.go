package wise

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryPrice(t *testing.T) {
	cases := []struct {
		amount   float64
		currency string
		payWith  string
	}{
		{1000, "USD", "GBP"},
		{1500, "USD", "NOK"},
		{2000, "USD", "EUR"},
	}

	for _, c := range cases {
		ctx := context.Background()
		price, err := QueryPrice(ctx, c.amount, c.currency, c.payWith)
		assert.NoError(t, err)
		assert.IsType(t, &Price{}, price)
		assert.Equal(t, c.currency, price.TargetCurrency)
		assert.Equal(t, c.payWith, price.SourceCurrency)
		assert.Equal(t, c.amount, price.TargetAmount)
	}
}
func Test_QueryMidRate(t *testing.T) {
	cases := []struct {
		from string
		to   string
	}{
		{"GBP", "TWD"},
		{"NOK", "TWD"},
		{"EUR", "TWD"},
	}

	ctx := context.Background()
	for _, c := range cases {
		rate, err := QueryRate(ctx, c.from, c.to)
		assert.NoError(t, err)
		assert.Greater(t, rate, 0.0)
	}
}
