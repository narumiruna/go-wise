# go-wise

## Usage

```go
ctx := context.Background()

price, err := wise.QueryPrice(ctx, 1000, "USD", "GBP")
if err != nil {
    panic(err)
}

cost, err := wise.NewCost(ctx, price)
if err != nil {
    panic(err)
}

fmt.Printf("%+v", cost)
```
