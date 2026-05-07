package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimit struct {
	rate float64
	tokens float64
	capacity float64
	lastRefill time.Time
	mu sync.Mutex
}

func NewRateLimit(rate int, capacity int) *RateLimit{
	rl := &RateLimit{
		rate: float64(rate),
		capacity:  float64(capacity),
		lastRefill: time.Now(),
		tokens: float64(capacity),
	}

	return rl
}

func (rl *RateLimit) Allow() bool{
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()

	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity	
	}	

	rl.lastRefill = now

	if rl.tokens >= 1.0 {
		rl.tokens--
		return true
	}

	return false
	
}

func main() {
	rl := NewRateLimit(2, 5)

	for range 10 {
		if rl.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}

		time.Sleep(200 * time.Millisecond)
	}

}

//===== TOKEN BUCKET
// type RateLimit struct {
// 	tokens		chan struct{}
// 	refillTime 	time.Duration
// }

// func NewRateLimit(rateLimit int, refillTime time.Duration) *RateLimit {
// 	rl := &RateLimit{
// 		tokens: make(chan struct{}, rateLimit),
// 		refillTime: refillTime,
// 	}

// 	for range rateLimit {
// 	 	rl.tokens <- struct{}{}
// 	}

// 	go rl.startRefill()

// 	return rl
// }

// func (rl *RateLimit) startRefill() {
// 	ticker := time.NewTicker(rl.refillTime)
// 	defer ticker.Stop()
	
// 	for range ticker.C{
// 		select {
// 			case rl.tokens <- struct{}{}:
// 			default:
// 		}
// 	}
// }

// func (rl *RateLimit) allow() bool {
// 	select {
// 	case <-rl.tokens:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func main(){
// 	rateLimit := NewRateLimit(5, time.Second)

// 	for range 10 {
// 		if rateLimit.allow(){
// 			fmt.Println("Request allowed")
// 		} else {
// 			fmt.Println("Request denied")
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }