package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			fmt.Println("Stopping ticker.")
			return
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		}
	}
}

// ======== SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		default:
// 			continue
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	i := 0
// 	for tick := range ticker.C {
// 		i++
// 		fmt.Println("i:", i)
// 		fmt.Println("Tick at:", tick)
// 	}

// }
