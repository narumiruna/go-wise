package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	v1 "github.com/narumiruna/go-wise/pkg/wise/v1"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("WISE_TOKEN")

	c := v1.NewRestClient()
	c.Auth(token)

	now := time.Now()
	ctx := context.Background()
	rates, err := c.QueryRateHistory(ctx, "GBP", "USD", now.AddDate(0, 0, -1), now, v1.GroupHour)
	if err != nil {
		panic(err)
	}

	for _, rate := range rates {
		log.Infof("%+v", rate)
	}
}
