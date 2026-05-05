package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu  	sync.Mutex
	count 	int
}

func (c *counter) incement(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main(){
	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 10

	for range numGoroutines {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for range 1000 {
				counter.incement()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())
}


// func main() {
// 	var count int
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup
// 	numGoroutines := 5

// 	increment := func() {
// 		defer wg.Done()
// 		for range 1000 {
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}

// 	}

// 	for range numGoroutines {
// 		wg.Add(1)
// 		go increment()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final count:", count)
// }