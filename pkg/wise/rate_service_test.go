package wise

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RateService_QueryRate(t *testing.T) {
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
		rate, err := client.NewRatesService().QueryRate(ctx, c.source, c.target)
		assert.NoError(t, err)

		assert.IsType(t, &Rate{}, rate)
		assert.Equal(t, c.source, rate.Source)
		assert.Equal(t, c.target, rate.Target)
		assert.Greater(t, rate.Value, 0.0)
	}
}

func Test_RateService_QueryRateHistory(t *testing.T) {
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
		rates, err := client.NewRatesService().QueryRateHistory(ctx, c.source, c.target, c.length, c.resolution, c.unit)
		assert.NoError(t, err)

		for _, rate := range rates {
			assert.IsType(t, Rate{}, rate)
			assert.Equal(t, c.source, rate.Source)
			assert.Equal(t, c.target, rate.Target)
			assert.Greater(t, rate.Value, 0.0)
		}
	}
}
