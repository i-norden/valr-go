# Valr Golang API Client

Go API wrapper for the [valr exchange](https://www.valr.com)

Based on the [luno exchange API wrapper](https://github.com/luno/luno-go)

### Installation

```
go get github.com/i-norden/valr-go
```

### Authentication

Public and private API keys can be generated within your account at the [exchange](https://www.valr.com)

### Example usage

```go
import (
	"context"
	
	valr "github.com/i-norden/valr-go/api"
)
valrClient := valr.NewClient()
valrClient.SetAuth("api_key_public", "api_key_secret")

req := valr.GetOrderBookRequest{Pair: "XBTZAR"}
res, err := valrClient.GetOrderBook(context.Background(), &req)
if err != nil {
  log.Fatal(err)
}
log.Println(res)
```

### License

[MIT](https://github.com/i-norden/valr-go/blob/master/LICENSE)