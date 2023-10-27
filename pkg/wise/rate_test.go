package wise

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	client := NewRestClient()
	for _, c := range cases {
		rate, err := client.QueryRate(ctx, c.source, c.target)
		assert.NoError(t, err)
		assert.Greater(t, rate, 0.0)
	}
}
