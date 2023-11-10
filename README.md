# go-wise

## Usage

```go
client := wise.NewRestClient()
ctx := context.Background()

cost, err := client.NewCost(ctx, "GBP", 1000, "USD")
if err != nil {
	panic(err)
}

fmt.Println(cost.String())
```
