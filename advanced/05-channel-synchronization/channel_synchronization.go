package main

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done
// 	fmt.Println("Finished.")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value received
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch
// 	fmt.Println(value)

// }

// ================ SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // SENDING SIGNAL OF COMPLETION
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
// 	}

// 	fmt.Println("All goroutines are complete")
// }

// ================ SYNCHRONIZING DATA EXCHANGE
func main() {

	data := make(chan string)

	go func() {
		defer close(data)
		for i := range 5 {
			data <- "hello " + string(rune('0'+i))
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}
}
