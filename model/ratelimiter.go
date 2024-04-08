package model

import (
	"sync"
	"time"
)

type RateLimiterModel struct {
	Mutex    sync.RWMutex
	Requests map[string]time.Time
}

var RateLimiter = &RateLimiterModel{
	Requests: make(map[string]time.Time),
}
