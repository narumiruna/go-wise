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

	ctx := context.Background()
	client := NewRestClient()
	for _, c := range cases {
		prices, err := client.QueryPrice(ctx, c.source, c.amount, c.target)
		assert.NoError(t, err)

		price, err := findPrice(prices, "VISA_CREDIT", "BALANCE")
		assert.NoError(t, err)

		assert.IsType(t, &Price{}, price)
		assert.Equal(t, c.target, price.TargetCurrency)
		assert.Equal(t, c.source, price.SourceCurrency)
		assert.Equal(t, c.amount, price.TargetAmount)
	}
}
