package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	AddMutex()
	AddAtomic()
	StoreLoadSwap()
	compareAndSwap()
}

func AddMutex() {
	start := time.Now()

	var (
		counter int64
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Finaly count:", counter)
	fmt.Println("With mutex:", time.Since(start).Seconds())
}

func AddAtomic() {
	start := time.Now()

	var (
		counter atomic.Int64
		wg sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println("Finaly count:", counter.Load())
	fmt.Println("With Atomic:", time.Since(start).Seconds())
}

func StoreLoadSwap() {
	var counter atomic.Int64
	fmt.Println("First load:", counter.Load())
	counter.Store(5) // Публикация нового значения
	fmt.Println("Secondary load after store:",counter.Load()) 
	fmt.Println("Swap:", counter.Swap(10)) // Установка значения и возврат предыдущего
	fmt.Println("Finaly load after swap:", counter.Load())
}

func compareAndSwap() {
	var (
		count 	atomic.Int64
		wg 		sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if !count.CompareAndSwap(0, 1) {
				return
			}

			fmt.Println("Swapped in goroutine id:", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Finaly:", count.Load())

}

// type AtomicCounter struct {
// 	count int64
// }

// func (ac *AtomicCounter) increment() {
// 	atomic.AddInt64(&ac.count, 1)
// }

// func (ac *AtomicCounter) getCount() int64 {
// 	return atomic.LoadInt64(&ac.count)
// }

// func main(){
// 	var wg sync.WaitGroup
// 	counter := &AtomicCounter{}
// 	numGoroutines := 10

// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Result count:", counter.getCount())
// }