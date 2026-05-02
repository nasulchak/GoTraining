package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WroderID %d started %s\n", w.ID, w.task)
	time.Sleep(time.Second)
	fmt.Printf("WroderID %d finished %s\n", w.ID, w.task)
}

func main() {
	var wg sync.WaitGroup

	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i, task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("Construction finished!")
}

// EXAMPLE WITH CHANNELS
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting.\n", id)
// 	for task := range tasks {
// 		results <- task * 2
// 		time.Sleep(time.Second)
// 	}
// 	fmt.Printf("Worker %d finished.\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)
// 	fmt.Println("Closed channel tasks")

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}

// 	fmt.Println("End of the program")
// }

// =========== BASIC EXAMPLE WITHOUT USING CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting.\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished.\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers finished!")
// }
