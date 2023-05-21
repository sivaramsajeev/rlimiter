package configs

import "fmt"

const (
	HTTP_PORT             = ":8888"
	COUNT_PER_MINUTE      = 10
	RESET_TIME_IN_MINUTES = 1
)

var (
	NORMAL_MSG       = "You aint ratelimited BOSS\n"
	RESET_MSG        = fmt.Sprintf("Bucket size re-set to: %d for %d minutes\n", COUNT_PER_MINUTE, RESET_TIME_IN_MINUTES)
	RATE_LIMITED_MSG = "Oops! You're ratelimited\n"
)
