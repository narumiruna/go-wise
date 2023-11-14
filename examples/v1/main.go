package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/narumiruna/go-wise/pkg/transferwise"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("WISE_TOKEN")

	c := transferwise.NewClient(token)
	now := time.Now()
	ctx := context.Background()
	rates, err := c.NewRateService().QueryRateHistory(ctx, "GBP", "USD", now.AddDate(0, 0, -1), now, transferwise.GroupHour)
	if err != nil {
		panic(err)
	}

	for _, rate := range rates {
		log.Infof("%+v", rate)
	}
}
