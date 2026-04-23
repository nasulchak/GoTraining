package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, result chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulate
		time.Sleep(time.Second)
		result <- task * 2
	}
	fmt.Printf("----->End worker %d\n", id)
}

func main() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send values to the tasks channel
	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	// Collect the result
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}

	fmt.Println("End of the program")
}
