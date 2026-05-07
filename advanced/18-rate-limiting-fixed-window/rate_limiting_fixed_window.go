package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimit struct {
	mu sync.Mutex
	limit int
	count int
	window time.Duration
	resetTime time.Time
}

func NewRateLimit(limit int, window time.Duration) *RateLimit{
	return &RateLimit{
		limit: limit,
		window: window,
		resetTime: time.Now().Truncate(window).Add(window),

	}
}

func (rl *RateLimit) Allow() bool{
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Truncate(rl.window).Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}

	return false
}

func main(){
	var wg sync.WaitGroup
	rateLimit := NewRateLimit(3, time.Second)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rateLimit.Allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
			}
		}()
	}

	wg.Wait()
}