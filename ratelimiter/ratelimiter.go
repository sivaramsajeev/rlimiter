package ratelimiter

import (
	"log"
	"time"

	"github.com/sivaramsajeev/rlimiter/configs"
	"github.com/sivaramsajeev/rlimiter/storage"
)

type RateLimiter struct {
	Bucket storage.IBucket
}

func (r *RateLimiter) StartRateLimiter() {
	go func() {
		for {
			r.Bucket.SetMaxSize(configs.COUNT_PER_MINUTE)
			log.Println(configs.RESET_MSG)
			time.Sleep(configs.RESET_TIME_IN_MINUTES * time.Minute)
		}
	}()
}
