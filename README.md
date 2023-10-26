# go-wise

## Usage

```go
ctx := context.Background()

price, err := wise.QueryPrice(ctx, "GBP", 1000, "USD")
if err != nil {
    panic(err)
}

cost, err := wise.NewCost(ctx, price)
if err != nil {
    panic(err)
}

fmt.Printf("%+v", cost)
```
