package main

import (
	"github.com/sivaramsajeev/rlimiter/ratelimiter"
	"github.com/sivaramsajeev/rlimiter/server"
	"github.com/sivaramsajeev/rlimiter/storage"
)

func main() {
	(&server.Server{
		RateLimiter: &ratelimiter.RateLimiter{
			Bucket: storage.NewInMemBucket(),
		},
	}).Serve()

}
