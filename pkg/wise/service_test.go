package wise

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Service_QueryPrice(t *testing.T) {
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
	client := NewClient()
	for _, c := range cases {
		prices, err := client.NewService().QueryPrice(ctx, c.source, c.amount, c.target)
		assert.NoError(t, err)

		for _, price := range prices {
			assert.IsType(t, Price{}, price)
			assert.Equal(t, c.target, price.TargetCurrency)
			assert.Equal(t, c.source, price.SourceCurrency)
			assert.Equal(t, c.amount, price.TargetAmount)
		}
	}
}

func Test_Service_QueryRate(t *testing.T) {
	cases := []struct {
		source string
		target string
	}{
		{"GBP", "TWD"},
		{"NOK", "TWD"},
		{"EUR", "TWD"},
	}

	ctx := context.Background()
	client := NewClient()
	for _, c := range cases {
		rate, err := client.NewService().QueryRate(ctx, c.source, c.target)
		assert.NoError(t, err)

		assert.IsType(t, &Rate{}, rate)
		assert.Equal(t, c.source, rate.Source)
		assert.Equal(t, c.target, rate.Target)
		assert.Greater(t, rate.Value, 0.0)
	}
}

func Test_Service_QueryRateHistory(t *testing.T) {
	cases := []struct {
		source     string
		target     string
		length     int
		resolution Resolution
		unit       Unit
	}{
		{"GBP", "TWD", 1, ResolutionHourly, UnitDay},
		{"NOK", "TWD", 1, ResolutionHourly, UnitDay},
		{"EUR", "TWD", 1, ResolutionHourly, UnitDay},
	}

	ctx := context.Background()
	client := NewClient()
	for _, c := range cases {
		rates, err := client.NewService().QueryRateHistory(ctx, c.source, c.target, c.length, c.resolution, c.unit)
		assert.NoError(t, err)

		for _, rate := range rates {
			assert.IsType(t, Rate{}, rate)
			assert.Equal(t, c.source, rate.Source)
			assert.Equal(t, c.target, rate.Target)
			assert.Greater(t, rate.Value, 0.0)
		}
	}
}
