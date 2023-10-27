# go-wise

## Usage

```go
client := wise.NewRestClient()
ctx := context.Background()

price, err := client.QueryPrice(ctx, "GBP", 1000, "USD")
if err != nil {
    panic(err)
}

cost, err := wise.NewCost(ctx, price)
if err != nil {
    panic(err)
}

fmt.Printf("%+v\n", cost)
```
