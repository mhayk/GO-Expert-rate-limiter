package model

type RateLimiter struct {
	MaxRequestsPerSecond      int
	BlockDurationSeconds      int
	IPMaxRequestsPerSecond    int
	IPBlockDurationSeconds    int
	TokenMaxRequestsPerSecond int
	TokenBlockDurationSeconds int
}
