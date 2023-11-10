package main

import (
	"context"
	"os"

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

	req := c.NewRatesRequest().Source("GBP").Target("USD").From("2022-01-01").To("2022-01-02").Group(v1.GroupHour)
	ctx := context.Background()
	rates, err := req.Do(ctx)
	if err != nil {
		panic(err)
	}

	for _, rate := range rates {
		log.Infof("%+v", rate)
	}
}
