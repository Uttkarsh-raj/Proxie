package model

import "time"

type RateLimiterModel struct {
	Requests map[string]time.Time
}

var RateLimiter = &RateLimiterModel{
	Requests: make(map[string]time.Time),
}
