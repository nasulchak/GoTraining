package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	mu         sync.Mutex
	capacity   time.Duration
	leakRate   time.Duration
	nextFreeAt time.Time
}

func NewLeakyBucket(capacity int, rps float64) *LeakyBucket {
	leakRate := time.Duration(float64(time.Second) / rps)
	return &LeakyBucket{
		leakRate:   leakRate,
		capacity:   time.Duration(capacity) * leakRate,
		nextFreeAt: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()

	if lb.nextFreeAt.Before(now) {
		lb.nextFreeAt = now
	}

	currentQueueDepth := lb.nextFreeAt.Sub(now)

	if currentQueueDepth > lb.capacity {
		return false
	}
	lb.nextFreeAt = lb.nextFreeAt.Add(lb.leakRate)
	return true
}

func main() {
	bucket := NewLeakyBucket(5, 2)

	fmt.Println("=== Мгновенный всплеск 10 запросов ===")
	for i := 1; i <= 10; i++ {
		if bucket.Allow() {
			fmt.Printf("Запрос %2d: [OK]  Принят в очередь\n", i)
		} else {
			fmt.Printf("Запрос %2d: [FAIL] Отклонен (ведро переполнено)\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// type LeakyBucket struct {
// 	capacity int
// 	tokens int
// 	rateLeak time.Duration
// 	lastLeak time.Time
// 	mu sync.Mutex
// }

// func NewLeakyBucket(capacity int, rateLeak time.Duration) *LeakyBucket {
// 	return &LeakyBucket{
// 		capacity: capacity,
// 		rateLeak: rateLeak,
// 		lastLeak: time.Now(),
// 	}
// }

// func (lb *LeakyBucket) Allow() bool {
// 	lb.mu.Lock()
// 	defer lb.mu.Unlock()

// 	now := time.Now()
// 	elapsed := now.Sub(lb.lastLeak)

// 	tokensAdd := int(elapsed / lb.rateLeak)

// 	lb.tokens -= tokensAdd

// 	if lb.tokens < 0 {
// 		lb.tokens = 0
// 	}

// 	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensAdd) * lb.rateLeak)

// 	if lb.tokens > lb.capacity {
// 		return false
// 	}

// 	lb.tokens++
// 	return true
// }

// func main(){
// 	leakyBucket := NewLeakyBucket(5, 500 * time.Millisecond)

// 	for range 10 {
// 		if leakyBucket.Allow() {
// 			fmt.Println("Request allowed")
// 		} else {
// 			fmt.Println("Request denied")
// 		}
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }