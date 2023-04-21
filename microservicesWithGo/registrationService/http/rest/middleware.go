package rest

import (
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	libredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func RateLimiter(next http.Handler) http.Handler {
	rate, err := limiter.NewRateFromFormatted("4-S")
	if err != nil {
		panic(err)
	}

	// Redis client
	option, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(option)

	// Configure store with redis
	store, err := libredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_http_example",
		MaxRetry: 3,
	})
	if err != nil {
		panic(err)
	}

	middleware := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
	return middleware.Handler(next)
}
