# go-wise

## Usage

```go
client := wise.NewRestClient()
ctx := context.Background()

cost, err := client.NewCost(ctx, "GBP", 1000, "USD")
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", cost)

rate, err := client.QueryRate(ctx, "GBP", "USD")
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", rate)

rates, err := client.QueryRateHistory(ctx, "GBP", "USD", 14, wise.ResolutionHourly, wise.UnitDay)
if err != nil {
	panic(err)
}
fmt.Printf("len(rates) = %+v\n", len(rates))
```
