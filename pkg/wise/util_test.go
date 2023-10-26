package wise

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryPrice(t *testing.T) {
	cases := []struct {
		source string
		amount float64
		target string
	}{
		{"GBP", 1000, "USD"},
		{"NOK", 1500, "USD"},
		{"EUR", 2000, "USD"},
	}

	for _, c := range cases {
		ctx := context.Background()
		price, err := QueryPrice(ctx, c.source, c.amount, c.target)
		assert.NoError(t, err)
		assert.IsType(t, &Price{}, price)
		assert.Equal(t, c.target, price.TargetCurrency)
		assert.Equal(t, c.source, price.SourceCurrency)
		assert.Equal(t, c.amount, price.TargetAmount)
	}
}
func Test_QueryMidRate(t *testing.T) {
	cases := []struct {
		source string
		target string
	}{
		{"GBP", "TWD"},
		{"NOK", "TWD"},
		{"EUR", "TWD"},
	}

	ctx := context.Background()
	for _, c := range cases {
		rate, err := QueryRate(ctx, c.source, c.target)
		assert.NoError(t, err)
		assert.Greater(t, rate, 0.0)
	}
}
